# 股票交易系统 (Stock Trading System)

这是一个用Go语言实现的股票交易系统，提供了完整的股票交易功能，包括下单、成交、投资组合管理等核心功能。

## 功能特性

### 核心结构体

- **Stock**: 股票信息，包含代码、名称、价格、涨跌幅等
- **Order**: 订单信息，支持买入/卖出订单
- **Trade**: 交易记录，记录所有成交的交易
- **Portfolio**: 投资组合，管理用户的持仓和现金
- **Holding**: 持仓信息，包含股票数量、成本、盈亏等
- **TradingEngine**: 交易引擎，核心业务逻辑处理

### 主要功能

1. **股票管理**
   - 添加股票
   - 更新股票价格
   - 获取股票信息
   - 计算涨跌幅

2. **订单管理**
   - 创建买入/卖出订单
   - 订单状态管理（待处理、已成交、已取消、已拒绝）
   - 订单验证和错误处理

3. **交易执行**
   - 订单成交处理
   - 自动更新投资组合
   - 计算平均成本和盈亏

4. **投资组合管理**
   - 实时计算总资产价值
   - 持仓管理
   - 现金余额跟踪
   - 未实现盈亏计算

5. **数据查询**
   - 用户订单查询
   - 交易记录查询
   - 投资组合查询
   - 市场数据查询

## 使用方法

### 基本使用

```go
package main

import "github.com/your-repo/newtrade"

func main() {
    // 创建交易引擎
    engine := newtrade.NewTradingEngine()
    
    // 添加股票
    engine.AddStock(&newtrade.Stock{
        Symbol: "AAPL",
        Name:   "Apple Inc.",
        Price:  150.0,
    })
    
    // 创建买入订单
    order := &newtrade.Order{
        Symbol:   "AAPL",
        Type:     newtrade.OrderTypeBuy,
        Quantity: 100,
        Price:    150.0,
        UserID:   "user123",
    }
    
    // 下单
    engine.PlaceOrder(order)
    
    // 执行订单
    engine.ExecuteOrder(order.ID)
    
    // 查询投资组合
    portfolio, _ := engine.GetPortfolio("user123")
    fmt.Printf("总资产: $%.2f\n", portfolio.TotalValue)
}
```

### 运行示例

```go
// 运行完整的交易模拟
newtrade.RunTradingSimulation()
```

## 系统架构

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   TradingEngine │    │     Orders      │    │      Trades     │
│                 │    │                 │    │                 │
│  - 股票管理     │    │  - 买入订单     │    │  - 成交记录     │
│  - 订单处理     │    │  - 卖出订单     │    │  - 交易历史     │
│  - 投资组合     │    │  - 订单状态     │    │  - 成交明细     │
│  - 交易执行     │    │  - 订单验证     │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │    Portfolio    │
                    │                 │
                    │  - 现金余额     │
                    │  - 持仓信息     │
                    │  - 资产价值     │
                    │  - 盈亏计算     │
                    └─────────────────┘
```

## 数据结构

### Stock (股票)
```go
type Stock struct {
    Symbol    string    // 股票代码
    Name      string    // 股票名称
    Price     float64   // 当前价格
    Change    float64   // 涨跌幅(%)
    Volume    int64     // 成交量
    MarketCap float64   // 市值
    UpdatedAt time.Time // 更新时间
}
```

### Order (订单)
```go
type Order struct {
    ID        string      // 订单ID
    Symbol    string      // 股票代码
    Type      OrderType   // 订单类型(BUY/SELL)
    Status    OrderStatus // 订单状态
    Quantity  int64       // 数量
    Price     float64     // 价格
    Total     float64     // 总金额
    UserID    string      // 用户ID
    CreatedAt time.Time   // 创建时间
    UpdatedAt time.Time   // 更新时间
}
```

### Portfolio (投资组合)
```go
type Portfolio struct {
    UserID     string             // 用户ID
    Holdings   map[string]Holding // 持仓信息
    Cash       float64            // 现金余额
    TotalValue float64            // 总资产价值
    UpdatedAt  time.Time          // 更新时间
}
```

## 错误处理

系统包含完善的错误处理机制：

- 股票不存在
- 订单参数无效
- 资金不足
- 持仓不足
- 订单状态错误

## 扩展功能

可以基于现有系统扩展以下功能：

1. **实时价格更新**: 集成外部API获取实时股票价格
2. **订单匹配引擎**: 实现买卖订单的自动匹配
3. **风险控制**: 添加止损、止盈等风险控制功能
4. **历史数据分析**: 提供交易历史分析和统计
5. **多用户支持**: 支持多用户并发交易
6. **数据库持久化**: 将数据保存到数据库
7. **Web API**: 提供RESTful API接口
8. **WebSocket**: 实时推送市场数据和订单状态

## 注意事项

1. 当前版本为内存存储，重启后数据会丢失
2. 价格更新需要手动调用，未实现自动更新
3. 订单执行是即时的，未实现限价单等高级订单类型
4. 建议在生产环境中添加数据持久化和并发安全控制

## 许可证

MIT License



