package newtrade

import (
	"testing"
	"time"
)

func TestNewTradingEngine(t *testing.T) {
	engine := NewTradingEngine()
	if engine == nil {
		t.Fatal("Failed to create trading engine")
	}

	if len(engine.stocks) != 0 {
		t.Error("New engine should have no stocks")
	}

	if len(engine.orders) != 0 {
		t.Error("New engine should have no orders")
	}

	if len(engine.trades) != 0 {
		t.Error("New engine should have no trades")
	}

	if len(engine.portfolios) != 0 {
		t.Error("New engine should have no portfolios")
	}
}

func TestAddStock(t *testing.T) {
	engine := NewTradingEngine()

	stock := &Stock{
		Symbol:    "AAPL",
		Name:      "Apple Inc.",
		Price:     150.0,
		Change:    0.0,
		Volume:    1000000,
		MarketCap: 2500000000000,
		UpdatedAt: time.Now(),
	}

	engine.AddStock(stock)

	if len(engine.stocks) != 1 {
		t.Error("Stock should be added")
	}

	if engine.stocks["AAPL"] != stock {
		t.Error("Stock should be stored with correct symbol")
	}
}

func TestGetStock(t *testing.T) {
	engine := NewTradingEngine()

	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}

	engine.AddStock(stock)

	// 测试获取存在的股票
	retrievedStock, err := engine.GetStock("AAPL")
	if err != nil {
		t.Errorf("Should not return error for existing stock: %v", err)
	}

	if retrievedStock != stock {
		t.Error("Should return the same stock instance")
	}

	// 测试获取不存在的股票
	_, err = engine.GetStock("GOOGL")
	if err == nil {
		t.Error("Should return error for non-existing stock")
	}
}

func TestUpdateStockPrice(t *testing.T) {
	engine := NewTradingEngine()

	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  100.0,
	}

	engine.AddStock(stock)

	// 测试更新价格
	err := engine.UpdateStockPrice("AAPL", 110.0)
	if err != nil {
		t.Errorf("Should not return error for valid update: %v", err)
	}

	if stock.Price != 110.0 {
		t.Errorf("Price should be updated to 110.0, got %f", stock.Price)
	}

	if stock.Change != 10.0 {
		t.Errorf("Change should be 10.0%%, got %f", stock.Change)
	}

	// 测试更新不存在的股票
	err = engine.UpdateStockPrice("GOOGL", 100.0)
	if err == nil {
		t.Error("Should return error for non-existing stock")
	}
}

func TestPlaceOrder(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 测试有效订单
	order := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	err := engine.PlaceOrder(order)
	if err != nil {
		t.Errorf("Should not return error for valid order: %v", err)
	}

	if order.ID == "" {
		t.Error("Order ID should be generated")
	}

	if order.Status != OrderStatusPending {
		t.Errorf("Order status should be PENDING, got %s", order.Status)
	}

	if order.Total != 15000.0 {
		t.Errorf("Order total should be 15000.0, got %f", order.Total)
	}

	// 测试无效订单参数
	invalidOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 0, // 无效数量
		Price:    150.0,
		UserID:   "user123",
	}

	err = engine.PlaceOrder(invalidOrder)
	if err == nil {
		t.Error("Should return error for invalid order parameters")
	}

	// 测试不存在的股票
	nonExistingStockOrder := &Order{
		Symbol:   "GOOGL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	err = engine.PlaceOrder(nonExistingStockOrder)
	if err == nil {
		t.Error("Should return error for non-existing stock")
	}
}

func TestExecuteOrder(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 创建订单
	order := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(order)

	// 执行订单
	err := engine.ExecuteOrder(order.ID)
	if err != nil {
		t.Errorf("Should not return error for valid execution: %v", err)
	}

	if order.Status != OrderStatusFilled {
		t.Errorf("Order status should be FILLED, got %s", order.Status)
	}

	// 检查交易记录
	trades := engine.GetTrades("user123")
	if len(trades) != 1 {
		t.Errorf("Should have 1 trade, got %d", len(trades))
	}

	trade := trades[0]
	if trade.OrderID != order.ID {
		t.Errorf("Trade should reference the order ID")
	}

	// 检查投资组合
	portfolio, err := engine.GetPortfolio("user123")
	if err != nil {
		t.Errorf("Should not return error for existing portfolio: %v", err)
	}

	if portfolio.Cash != 85000.0 { // 100000 - 15000
		t.Errorf("Portfolio cash should be 85000.0, got %f", portfolio.Cash)
	}

	holding, exists := portfolio.Holdings["AAPL"]
	if !exists {
		t.Error("Should have AAPL holding")
	}

	if holding.Quantity != 100 {
		t.Errorf("Holding quantity should be 100, got %d", holding.Quantity)
	}

	if holding.AvgPrice != 150.0 {
		t.Errorf("Holding average price should be 150.0, got %f", holding.AvgPrice)
	}
}

func TestCancelOrder(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 创建订单
	order := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(order)

	// 取消订单
	err := engine.CancelOrder(order.ID)
	if err != nil {
		t.Errorf("Should not return error for valid cancellation: %v", err)
	}

	if order.Status != OrderStatusCancelled {
		t.Errorf("Order status should be CANCELLED, got %s", order.Status)
	}

	// 测试取消已成交的订单
	engine.ExecuteOrder(order.ID)

	err = engine.CancelOrder(order.ID)
	if err == nil {
		t.Error("Should return error for filled order cancellation")
	}
}

func TestPortfolioCalculations(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 买入100股
	buyOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(buyOrder)
	engine.ExecuteOrder(buyOrder.ID)

	// 更新股票价格到160
	engine.UpdateStockPrice("AAPL", 160.0)

	// 获取投资组合
	portfolio, _ := engine.GetPortfolio("user123")

	// 检查持仓市值和盈亏
	holding := portfolio.Holdings["AAPL"]
	expectedMarketValue := float64(100) * 160.0
	expectedUnrealizedPL := expectedMarketValue - (150.0 * 100.0)

	if holding.MarketValue != expectedMarketValue {
		t.Errorf("Market value should be %f, got %f", expectedMarketValue, holding.MarketValue)
	}

	if holding.UnrealizedPL != expectedUnrealizedPL {
		t.Errorf("Unrealized P&L should be %f, got %f", expectedUnrealizedPL, holding.UnrealizedPL)
	}

	// 检查总资产价值
	expectedTotalValue := 85000.0 + expectedMarketValue // 现金 + 市值
	if portfolio.TotalValue != expectedTotalValue {
		t.Errorf("Total value should be %f, got %f", expectedTotalValue, portfolio.TotalValue)
	}
}

func TestSellOrder(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 先买入100股
	buyOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(buyOrder)
	engine.ExecuteOrder(buyOrder.ID)

	// 卖出50股
	sellOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeSell,
		Quantity: 50,
		Price:    155.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(sellOrder)
	engine.ExecuteOrder(sellOrder.ID)

	// 检查投资组合
	portfolio, _ := engine.GetPortfolio("user123")

	// 现金应该增加
	expectedCash := 85000.0 + (50 * 155.0) // 原现金 + 卖出收入
	if portfolio.Cash != expectedCash {
		t.Errorf("Cash should be %f, got %f", expectedCash, portfolio.Cash)
	}

	// 持仓应该减少
	holding := portfolio.Holdings["AAPL"]
	if holding.Quantity != 50 {
		t.Errorf("Holding quantity should be 50, got %d", holding.Quantity)
	}

	// 平均成本应该保持不变
	if holding.AvgPrice != 150.0 {
		t.Errorf("Average price should remain 150.0, got %f", holding.AvgPrice)
	}
}

func TestInsufficientFunds(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 尝试买入超过现金的股票
	order := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 1000, // 需要150000，但只有100000
		Price:    150.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(order)

	// 执行订单应该失败
	err := engine.ExecuteOrder(order.ID)
	if err == nil {
		t.Error("Should return error for insufficient funds")
	}

	// 订单状态应该仍然是待处理
	if order.Status != OrderStatusPending {
		t.Errorf("Order status should remain PENDING, got %s", order.Status)
	}
}

func TestInsufficientShares(t *testing.T) {
	engine := NewTradingEngine()

	// 添加股票
	stock := &Stock{
		Symbol: "AAPL",
		Name:   "Apple Inc.",
		Price:  150.0,
	}
	engine.AddStock(stock)

	// 先买入100股
	buyOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(buyOrder)
	engine.ExecuteOrder(buyOrder.ID)

	// 尝试卖出超过持仓的股票
	sellOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeSell,
		Quantity: 150, // 只有100股，但想卖150股
		Price:    155.0,
		UserID:   "user123",
	}

	engine.PlaceOrder(sellOrder)

	// 执行订单应该失败
	err := engine.ExecuteOrder(sellOrder.ID)
	if err == nil {
		t.Error("Should return error for insufficient shares")
	}

	// 订单状态应该仍然是待处理
	if sellOrder.Status != OrderStatusPending {
		t.Errorf("Order status should remain PENDING, got %s", sellOrder.Status)
	}
}



