package newtrade

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// ExampleUsage 展示交易系统的使用示例
func ExampleUsage() {
	// 创建交易引擎
	engine := NewTradingEngine()

	// 添加一些股票
	engine.AddStock(&Stock{
		Symbol:    "AAPL",
		Name:      "Apple Inc.",
		Price:     150.0,
		Change:    0.0,
		Volume:    1000000,
		MarketCap: 2500000000000,
		UpdatedAt: time.Now(),
	})

	engine.AddStock(&Stock{
		Symbol:    "GOOGL",
		Name:      "Alphabet Inc.",
		Price:     2800.0,
		Change:    0.0,
		Volume:    500000,
		MarketCap: 1800000000000,
		UpdatedAt: time.Now(),
	})

	engine.AddStock(&Stock{
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

	// 创建买入订单
	buyOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeBuy,
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

	// 执行订单
	if err := engine.ExecuteOrder(buyOrder.ID); err != nil {
		log.Printf("执行订单失败: %v", err)
		return
	}

	fmt.Println("买入订单已执行")

	// 创建卖出订单
	sellOrder := &Order{
		Symbol:   "AAPL",
		Type:     OrderTypeSell,
		Quantity: 50,
		Price:    155.0,
		UserID:   userID,
	}

	// 下单
	if err := engine.PlaceOrder(sellOrder); err != nil {
		log.Printf("下单失败: %v", err)
		return
	}

	fmt.Printf("卖出订单已创建: %s\n", sellOrder.ID)

	// 执行订单
	if err := engine.ExecuteOrder(sellOrder.ID); err != nil {
		log.Printf("执行订单失败: %v", err)
		return
	}

	fmt.Println("卖出订单已执行")

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
}

// RunTradingSimulation 运行交易模拟
func RunTradingSimulation() {
	fmt.Println("开始股票交易模拟...")
	fmt.Println(strings.Repeat("=", 50))

	ExampleUsage()

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("股票交易模拟完成!")
}
