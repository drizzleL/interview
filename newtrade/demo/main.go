package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"newtrade"
)

func main() {
	fmt.Println("=== 股票交易系统演示 ===")
	fmt.Println()

	// 运行交易模拟
	runTradingSimulation()
}

// runTradingSimulation 运行交易模拟
func runTradingSimulation() {
	fmt.Println("开始股票交易模拟...")
	fmt.Println(strings.Repeat("=", 50))

	// 创建交易引擎
	engine := newtrade.NewTradingEngine()

	// 添加一些股票
	engine.AddStock(&newtrade.Stock{
		Symbol:    "AAPL",
		Name:      "Apple Inc.",
		Price:     150.0,
		Change:    0.0,
		Volume:    1000000,
		MarketCap: 2500000000000,
		UpdatedAt: time.Now(),
	})

	engine.AddStock(&newtrade.Stock{
		Symbol:    "GOOGL",
		Name:      "Alphabet Inc.",
		Price:     2800.0,
		Change:    0.0,
		Volume:    500000,
		MarketCap: 1800000000000,
		UpdatedAt: time.Now(),
	})

	engine.AddStock(&newtrade.Stock{
		Symbol:    "TSLA",
		Name:      "Tesla Inc.",
		Price:     800.0,
		Change:    0.0,
		Volume:    800000,
		MarketCap: 800000000000,
		UpdatedAt: time.Now(),
	})

	// 用户ID
	userID := "user123"

	// 创建买入订单（AAPL）
	buyOrder := &newtrade.Order{
		Symbol:   "AAPL",
		Type:     newtrade.OrderTypeBuy,
		Quantity: 100,
		Price:    150.0,
		UserID:   userID,
	}

	// 下单
	if err := engine.PlaceOrder(buyOrder); err != nil {
		log.Printf("下单失败: %v", err)
		return
	}
	fmt.Printf("买入订单已创建: %s\n", buyOrder.ID)

	// 展示当前挂单（执行前）
	printOpenOrders(engine, "下单后、执行前")

	// 执行订单
	if err := engine.ExecuteOrder(buyOrder.ID); err != nil {
		log.Printf("执行订单失败: %v", err)
		return
	}
	fmt.Println("买入订单已执行")

	// 展示当前挂单（执行后）
	printOpenOrders(engine, "买单执行后")

	// 创建卖出订单（AAPL）
	sellOrder := &newtrade.Order{
		Symbol:   "AAPL",
		Type:     newtrade.OrderTypeSell,
		Quantity: 50,
		Price:    155.0,
		UserID:   userID,
	}
	if err := engine.PlaceOrder(sellOrder); err != nil {
		log.Printf("下单失败: %v", err)
		return
	}
	fmt.Printf("卖出订单已创建: %s\n", sellOrder.ID)

	// 再创建一个未执行的挂单（GOOGL），用于演示过滤
	pendingOrder := &newtrade.Order{
		Symbol:   "GOOGL",
		Type:     newtrade.OrderTypeBuy,
		Quantity: 10,
		Price:    2800.0,
		UserID:   userID,
	}
	if err := engine.PlaceOrder(pendingOrder); err != nil {
		log.Printf("创建演示挂单失败: %v", err)
	} else {
		fmt.Printf("额外挂单已创建(不执行，用于演示): %s\n", pendingOrder.ID)
	}

	// 展示当前挂单（卖单未执行前）
	printOpenOrders(engine, "卖单未执行前")

	// 执行卖出订单
	if err := engine.ExecuteOrder(sellOrder.ID); err != nil {
		log.Printf("执行订单失败: %v", err)
		return
	}
	fmt.Println("卖出订单已执行")

	// 展示当前挂单（卖单执行后，保留GOOGL挂单）
	printOpenOrders(engine, "卖单执行后（剩余GOOGL挂单）")

	// 更新股票价格
	engine.UpdateStockPrice("AAPL", 160.0)
	engine.UpdateStockPrice("GOOGL", 2850.0)
	engine.UpdateStockPrice("TSLA", 850.0)

	// 获取用户投资组合
	portfolio, err := engine.GetPortfolio(userID)
	if err != nil {
		log.Printf("获取投资组合失败: %v", err)
		return
	}

	fmt.Printf("\n=== 用户投资组合 ===\n")
	fmt.Printf("用户ID: %s\n", portfolio.UserID)
	fmt.Printf("现金余额: $%.2f\n", portfolio.Cash)
	fmt.Printf("总资产价值: $%.2f\n", portfolio.TotalValue)
	fmt.Printf("更新时间: %s\n", portfolio.UpdatedAt.Format("2006-01-02 15:04:05"))

	fmt.Printf("\n=== 持仓详情 ===\n")
	for symbol, holding := range portfolio.Holdings {
		fmt.Printf("股票: %s\n", symbol)
		fmt.Printf("  持有数量: %d\n", holding.Quantity)
		fmt.Printf("  平均成本: $%.2f\n", holding.AvgPrice)
		fmt.Printf("  总成本: $%.2f\n", holding.TotalCost)
		fmt.Printf("  市值: $%.2f\n", holding.MarketValue)
		fmt.Printf("  未实现盈亏: $%.2f\n", holding.UnrealizedPL)
		fmt.Println()
	}

	// 获取用户订单
	orders := engine.GetOrders(userID)
	fmt.Printf("=== 用户订单 ===\n")
	for _, order := range orders {
		fmt.Printf("订单ID: %s\n", order.ID)
		fmt.Printf("  股票: %s\n", order.Symbol)
		fmt.Printf("  类型: %s\n", order.Type)
		fmt.Printf("  状态: %s\n", order.Status)
		fmt.Printf("  数量: %d\n", order.Quantity)
		fmt.Printf("  价格: $%.2f\n", order.Price)
		fmt.Printf("  总金额: $%.2f\n", order.Total)
		fmt.Printf("  创建时间: %s\n", order.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}

	// 获取用户交易记录
	trades := engine.GetTrades(userID)
	fmt.Printf("=== 交易记录 ===\n")
	for _, trade := range trades {
		fmt.Printf("交易ID: %s\n", trade.ID)
		fmt.Printf("  订单ID: %s\n", trade.OrderID)
		fmt.Printf("  股票: %s\n", trade.Symbol)
		fmt.Printf("  类型: %s\n", trade.Type)
		fmt.Printf("  数量: %d\n", trade.Quantity)
		fmt.Printf("  价格: $%.2f\n", trade.Price)
		fmt.Printf("  总金额: $%.2f\n", trade.Total)
		fmt.Printf("  成交时间: %s\n", trade.Timestamp.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}

	// 获取市场数据
	marketData := engine.GetMarketData()
	fmt.Printf("=== 市场数据 ===\n")
	for _, stock := range marketData {
		fmt.Printf("股票: %s (%s)\n", stock.Symbol, stock.Name)
		fmt.Printf("  当前价格: $%.2f\n", stock.Price)
		fmt.Printf("  涨跌幅: %.2f%%\n", stock.Change)
		fmt.Printf("  成交量: %d\n", stock.Volume)
		fmt.Printf("  市值: $%.0f\n", stock.MarketCap)
		fmt.Printf("  更新时间: %s\n", stock.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println()
	}

	// 新增：双用户自动撮合演示
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("自动撮合演示（价格-时间优先）...")
	buyer := &newtrade.Order{UserID: "alice", Symbol: "AAPL", Type: newtrade.OrderTypeBuy, Quantity: 120, Price: 161.0}
	sellerEarly := &newtrade.Order{UserID: "bob", Symbol: "AAPL", Type: newtrade.OrderTypeSell, Quantity: 50, Price: 160.0}
	sellerLate := &newtrade.Order{UserID: "carl", Symbol: "AAPL", Type: newtrade.OrderTypeSell, Quantity: 100, Price: 161.0}

	// 先下一个卖单(160)，再下买单(161)，再下另一个卖单(161)以观察时间优先
	_ = engine.PlaceOrder(sellerEarly)
	_ = engine.PlaceOrder(buyer)
	_ = engine.PlaceOrder(sellerLate)

	// 打印撮合结果（可能部分成交）
	printOrderMatchResult("买方(alice)", buyer)
	printOrderMatchResult("卖方(bob)", sellerEarly)
	printOrderMatchResult("卖方(carl)", sellerLate)

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("股票交易模拟完成!")
}

func printOpenOrders(engine *newtrade.TradingEngine, title string) {
	fmt.Printf("\n=== 挂单列表（%s） ===\n", title)
	all := engine.GetOpenOrders()
	fmt.Printf("全部挂单: %d\n", len(all))
	for _, o := range all {
		fmt.Printf("  [%s] %s %d @ $%.2f (用户: %s, 订单ID: %s)\n", o.Symbol, o.Type, o.Quantity, o.Price, o.UserID, o.ID)
	}
	fmt.Printf("按用户筛选(user123): %d\n", len(engine.GetOpenOrdersByUser("user123")))
	fmt.Printf("按股票筛选(AAPL): %d\n\n", len(engine.GetOpenOrdersBySymbol("AAPL")))
}

func printOrderMatchResult(title string, o *newtrade.Order) {
	fmt.Printf("%s: 订单ID=%s, 类型=%s, 价格=%.2f, 数量=%d, 已成交=%d, 状态=%s\n",
		title, o.ID, o.Type, o.Price, o.Quantity, o.FilledQuantity, o.Status)
}
