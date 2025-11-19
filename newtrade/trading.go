package newtrade

import (
	"container/heap"
	"errors"
	"fmt"
	"time"
)

// OrderType 订单类型
type OrderType string

const (
	OrderTypeBuy  OrderType = "BUY"  // 买入
	OrderTypeSell OrderType = "SELL" // 卖出
)

// OrderStatus 订单状态
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"   // 待处理
	OrderStatusFilled    OrderStatus = "FILLED"    // 已成交
	OrderStatusCancelled OrderStatus = "CANCELLED" // 已取消
	OrderStatusRejected  OrderStatus = "REJECTED"  // 已拒绝
)

// Stock 股票信息
type Stock struct {
	Symbol    string    `json:"symbol"`    // 股票代码
	Name      string    `json:"name"`      // 股票名称
	Price     float64   `json:"price"`     // 当前价格
	Change    float64   `json:"change"`    // 涨跌幅
	Volume    int64     `json:"volume"`    // 成交量
	MarketCap float64   `json:"marketCap"` // 市值
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

// Order 订单
type Order struct {
	ID             string      `json:"id"`             // 订单ID
	Symbol         string      `json:"symbol"`         // 股票代码
	Type           OrderType   `json:"type"`           // 订单类型
	Status         OrderStatus `json:"status"`         // 订单状态
	Quantity       int64       `json:"quantity"`       // 下单数量
	FilledQuantity int64       `json:"filledQuantity"` // 已成交数量
	Price          float64     `json:"price"`          // 价格
	Total          float64     `json:"total"`          // 总金额(Quantity * Price)
	UserID         string      `json:"userId"`         // 用户ID
	CreatedAt      time.Time   `json:"createdAt"`      // 创建时间
	UpdatedAt      time.Time   `json:"updatedAt"`      // 更新时间
	seq            int64       `json:"-"`              // 时间优先辅助序号
}

// Trade 交易记录
type Trade struct {
	ID        string    `json:"id"`        // 交易ID
	OrderID   string    `json:"orderId"`   // 订单ID
	Symbol    string    `json:"symbol"`    // 股票代码
	Type      OrderType `json:"type"`      // 交易类型
	Quantity  int64     `json:"quantity"`  // 数量
	Price     float64   `json:"price"`     // 成交价格
	Total     float64   `json:"total"`     // 总金额
	UserID    string    `json:"userId"`    // 用户ID
	Timestamp time.Time `json:"timestamp"` // 成交时间
}

// Portfolio 投资组合
type Portfolio struct {
	UserID     string             `json:"userId"`     // 用户ID
	Holdings   map[string]Holding `json:"holdings"`   // 持仓信息
	Cash       float64            `json:"cash"`       // 现金余额
	TotalValue float64            `json:"totalValue"` // 总资产价值
	UpdatedAt  time.Time          `json:"updatedAt"`  // 更新时间
}

// Holding 持仓信息
type Holding struct {
	Symbol       string  `json:"symbol"`       // 股票代码
	Quantity     int64   `json:"quantity"`     // 持有数量
	AvgPrice     float64 `json:"avgPrice"`     // 平均成本
	TotalCost    float64 `json:"totalCost"`    // 总成本
	MarketValue  float64 `json:"marketValue"`  // 市值
	UnrealizedPL float64 `json:"unrealizedPL"` // 未实现盈亏
}

// 内部：订单堆（买单为最大堆，卖单为最小堆）
type orderHeap struct {
	items []*Order
	isBuy bool
}

func (h *orderHeap) Len() int { return len(h.items) }
func (h *orderHeap) Less(i, j int) bool {
	a, b := h.items[i], h.items[j]
	if h.isBuy {
		if a.Price != b.Price {
			return a.Price > b.Price // 买单价格高优先
		}
	} else {
		if a.Price != b.Price {
			return a.Price < b.Price // 卖单价格低优先
		}
	}
	// 价格相同，按时间（seq）优先
	return a.seq < b.seq
}
func (h *orderHeap) Swap(i, j int)      { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *orderHeap) Push(x interface{}) { h.items = append(h.items, x.(*Order)) }
func (h *orderHeap) Pop() interface{} {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

type orderBook struct {
	buys  *orderHeap // 价格优先：高->低
	sells *orderHeap // 价格优先：低->高
}

// TradingEngine 交易引擎
type TradingEngine struct {
	stocks     map[string]*Stock     // 股票信息
	orders     map[string]*Order     // 订单信息
	trades     map[string]*Trade     // 交易记录
	portfolios map[string]*Portfolio // 用户投资组合

	orderBooks map[string]*orderBook // 每个股票的订单簿
	orderSeq   int64                 // 全局序号（时间优先）
}

// NewTradingEngine 创建新的交易引擎
func NewTradingEngine() *TradingEngine {
	return &TradingEngine{
		stocks:     make(map[string]*Stock),
		orders:     make(map[string]*Order),
		trades:     make(map[string]*Trade),
		portfolios: make(map[string]*Portfolio),
		orderBooks: make(map[string]*orderBook),
	}
}

// EnsurePortfolio 确保用户投资组合存在，并在不存在时以指定现金创建
func (te *TradingEngine) EnsurePortfolio(userID string, initialCash float64) *Portfolio {
	if userID == "" {
		userID = "default_user"
	}
	p, exists := te.portfolios[userID]
	if !exists {
		p = &Portfolio{
			UserID:     userID,
			Holdings:   make(map[string]Holding),
			Cash:       initialCash,
			TotalValue: initialCash,
			UpdatedAt:  time.Now(),
		}
		te.portfolios[userID] = p
	}
	return p
}

func (te *TradingEngine) getBook(symbol string) *orderBook {
	b, ok := te.orderBooks[symbol]
	if !ok {
		b = &orderBook{buys: &orderHeap{isBuy: true}, sells: &orderHeap{isBuy: false}}
		heap.Init(b.buys)
		heap.Init(b.sells)
		te.orderBooks[symbol] = b
	}
	return b
}

// AddStock 添加股票
func (te *TradingEngine) AddStock(stock *Stock) {
	te.stocks[stock.Symbol] = stock
}

// GetStock 获取股票信息
func (te *TradingEngine) GetStock(symbol string) (*Stock, error) {
	if stock, exists := te.stocks[symbol]; exists {
		return stock, nil
	}
	return nil, errors.New("stock not found")
}

// UpdateStockPrice 更新股票价格
func (te *TradingEngine) UpdateStockPrice(symbol string, newPrice float64) error {
	stock, exists := te.stocks[symbol]
	if !exists {
		return errors.New("stock not found")
	}

	oldPrice := stock.Price
	stock.Price = newPrice
	stock.Change = ((newPrice - oldPrice) / oldPrice) * 100
	stock.UpdatedAt = time.Now()

	// 重新计算所有投资组合的价值
	te.RecalculatePortfolioValues()

	return nil
}

// PlaceOrder 下单（自动撮合）
func (te *TradingEngine) PlaceOrder(order *Order) error {
	// 验证股票是否存在
	if _, exists := te.stocks[order.Symbol]; !exists {
		return errors.New("stock not found")
	}

	// 验证订单参数
	if order.Quantity <= 0 || order.Price <= 0 {
		return errors.New("invalid order parameters")
	}

	// 基础设置
	order.Total = float64(order.Quantity) * order.Price
	order.Status = OrderStatusPending
	order.CreatedAt = time.Now()
	order.UpdatedAt = order.CreatedAt
	te.orderSeq++
	order.seq = te.orderSeq
	order.FilledQuantity = 0
	order.ID = fmt.Sprintf("ORD_%d", time.Now().UnixNano())

	// 基础风控：校验下单方是否具备全额资金或持仓（简化）
	if err := te.validateOrderExecution(order); err != nil {
		return err
	}

	// 保存订单
	te.orders[order.ID] = order

	// 自动撮合
	te.matchWithOrderBook(order)

	// 若未完全成交，入本方订单簿
	if order.Status == OrderStatusPending && order.FilledQuantity < order.Quantity {
		book := te.getBook(order.Symbol)
		if order.Type == OrderTypeBuy {
			heap.Push(book.buys, order)
		} else {
			heap.Push(book.sells, order)
		}
	}

	return nil
}

// CancelOrder 取消订单
func (te *TradingEngine) CancelOrder(orderID string) error {
	order, exists := te.orders[orderID]
	if !exists {
		return errors.New("order not found")
	}
	if order.Status != OrderStatusPending {
		return errors.New("order cannot be cancelled")
	}
	order.Status = OrderStatusCancelled
	order.UpdatedAt = time.Now()
	return nil
}

// ExecuteOrder 执行订单（手动强制全额成交，不参与撮合）
func (te *TradingEngine) ExecuteOrder(orderID string) error {
	order, exists := te.orders[orderID]
	if !exists {
		return errors.New("order not found")
	}
	if order.Status != OrderStatusPending {
		return errors.New("order cannot be executed")
	}
	// 预检查
	if err := te.validateOrderExecution(order); err != nil {
		return err
	}
	trade := &Trade{
		ID:        fmt.Sprintf("TRD_%d", time.Now().UnixNano()),
		OrderID:   order.ID,
		Symbol:    order.Symbol,
		Type:      order.Type,
		Quantity:  order.Quantity,
		Price:     order.Price,
		Total:     order.Total,
		UserID:    order.UserID,
		Timestamp: time.Now(),
	}
	te.trades[trade.ID] = trade
	order.Status = OrderStatusFilled
	order.FilledQuantity = order.Quantity
	order.UpdatedAt = time.Now()
	if err := te.updatePortfolio(trade); err != nil {
		return err
	}
	return nil
}

// validateOrderExecution 验证订单执行
func (te *TradingEngine) validateOrderExecution(order *Order) error {
	portfolio, exists := te.portfolios[order.UserID]
	if !exists {
		if order.Type == OrderTypeBuy {
			if order.Total > 100000.0 { // 初始资金
				return errors.New("insufficient funds")
			}
		}
		return nil
	}
	if order.Type == OrderTypeBuy {
		if portfolio.Cash < order.Total {
			return errors.New("insufficient funds")
		}
	} else if order.Type == OrderTypeSell {
		holding, exists := portfolio.Holdings[order.Symbol]
		if !exists || holding.Quantity < order.Quantity {
			return errors.New("insufficient shares")
		}
	}
	return nil
}

// matchWithOrderBook 使用订单簿撮合多个对手方订单
func (te *TradingEngine) matchWithOrderBook(incoming *Order) {
	book := te.getBook(incoming.Symbol)
	for incoming.FilledQuantity < incoming.Quantity {
		var opp *Order
		if incoming.Type == OrderTypeBuy {
			// 清理堆顶非挂单
			for book.sells.Len() > 0 {
				peek := book.sells.items[0]
				if peek.Status != OrderStatusPending || peek.FilledQuantity >= peek.Quantity {
					heap.Pop(book.sells)
					continue
				}
				break
			}
			if book.sells.Len() == 0 {
				break
			}
			opp = book.sells.items[0]
			// 价格是否可成交
			if incoming.Price < opp.Price {
				break
			}
		} else {
			for book.buys.Len() > 0 {
				peek := book.buys.items[0]
				if peek.Status != OrderStatusPending || peek.FilledQuantity >= peek.Quantity {
					heap.Pop(book.buys)
					continue
				}
				break
			}
			if book.buys.Len() == 0 {
				break
			}
			opp = book.buys.items[0]
			if incoming.Price > opp.Price {
				break
			}
		}
		// 现在 opp 为可成交的对手方挂单
		remainingIncoming := incoming.Quantity - incoming.FilledQuantity
		remainingOpp := opp.Quantity - opp.FilledQuantity
		if remainingIncoming <= 0 || remainingOpp <= 0 {
			break
		}
		qty := remainingIncoming
		if remainingOpp < qty {
			qty = remainingOpp
		}
		price := opp.Price // 以被动方价格成交
		total := float64(qty) * price
		timestamp := time.Now()

		// 生成双方成交
		var buyerID, sellerID string
		if incoming.Type == OrderTypeBuy {
			buyerID = incoming.UserID
			sellerID = opp.UserID
		} else {
			buyerID = opp.UserID
			sellerID = incoming.UserID
		}
		buyTrade := &Trade{
			ID: fmt.Sprintf("TRD_%d_B", time.Now().UnixNano()),
			OrderID: func() string {
				if incoming.Type == OrderTypeBuy {
					return incoming.ID
				} else {
					return opp.ID
				}
			}(),
			Symbol:    incoming.Symbol,
			Type:      OrderTypeBuy,
			Quantity:  qty,
			Price:     price,
			Total:     total,
			UserID:    buyerID,
			Timestamp: timestamp,
		}
		sellTrade := &Trade{
			ID: fmt.Sprintf("TRD_%d_S", time.Now().UnixNano()),
			OrderID: func() string {
				if incoming.Type == OrderTypeSell {
					return incoming.ID
				} else {
					return opp.ID
				}
			}(),
			Symbol:    incoming.Symbol,
			Type:      OrderTypeSell,
			Quantity:  qty,
			Price:     price,
			Total:     total,
			UserID:    sellerID,
			Timestamp: timestamp,
		}

		te.trades[buyTrade.ID] = buyTrade
		te.trades[sellTrade.ID] = sellTrade

		incoming.FilledQuantity += qty
		opp.FilledQuantity += qty
		incoming.UpdatedAt = timestamp
		opp.UpdatedAt = timestamp
		if incoming.FilledQuantity == incoming.Quantity {
			incoming.Status = OrderStatusFilled
		}
		if opp.FilledQuantity == opp.Quantity {
			opp.Status = OrderStatusFilled
			// 从堆中弹出已完全成交的对手方
			if incoming.Type == OrderTypeBuy {
				heap.Pop(book.sells)
			} else {
				heap.Pop(book.buys)
			}
		}

		// 更新投资组合（双边）
		_ = te.updatePortfolio(buyTrade)
		_ = te.updatePortfolio(sellTrade)
	}
}

// updatePortfolio 更新投资组合
func (te *TradingEngine) updatePortfolio(trade *Trade) error {
	portfolio, exists := te.portfolios[trade.UserID]
	if !exists {
		portfolio = &Portfolio{
			UserID:     trade.UserID,
			Holdings:   make(map[string]Holding),
			Cash:       100000.0, // 初始资金
			TotalValue: 100000.0,
			UpdatedAt:  time.Now(),
		}
		te.portfolios[trade.UserID] = portfolio
	}

	holding, exists := portfolio.Holdings[trade.Symbol]
	if !exists {
		holding = Holding{Symbol: trade.Symbol}
	}

	if trade.Type == OrderTypeBuy {
		if portfolio.Cash < trade.Total {
			return errors.New("insufficient funds")
		}
		portfolio.Cash -= trade.Total
		totalCost := (holding.AvgPrice * float64(holding.Quantity)) + trade.Total
		totalQuantity := holding.Quantity + trade.Quantity
		holding.AvgPrice = totalCost / float64(totalQuantity)
		holding.Quantity = totalQuantity
		holding.TotalCost = totalCost
	} else if trade.Type == OrderTypeSell {
		if holding.Quantity < trade.Quantity {
			return errors.New("insufficient shares")
		}
		portfolio.Cash += trade.Total
		holding.Quantity -= trade.Quantity
		if holding.Quantity == 0 {
			holding.AvgPrice = 0
			holding.TotalCost = 0
		} else {
			holding.TotalCost = holding.AvgPrice * float64(holding.Quantity)
		}
	}

	portfolio.Holdings[trade.Symbol] = holding
	portfolio.UpdatedAt = time.Now()

	portfolio.TotalValue = portfolio.Cash
	for symbol, h := range portfolio.Holdings {
		if stock, exists := te.stocks[symbol]; exists {
			h.MarketValue = float64(h.Quantity) * stock.Price
			h.UnrealizedPL = h.MarketValue - h.TotalCost
			portfolio.TotalValue += h.MarketValue
			portfolio.Holdings[symbol] = h
		}
	}
	return nil
}

// GetPortfolio 获取用户投资组合
func (te *TradingEngine) GetPortfolio(userID string) (*Portfolio, error) {
	portfolio, exists := te.portfolios[userID]
	if !exists {
		return nil, errors.New("portfolio not found")
	}
	return portfolio, nil
}

// GetOrders 获取用户订单
func (te *TradingEngine) GetOrders(userID string) []*Order {
	var userOrders []*Order
	for _, order := range te.orders {
		if order.UserID == userID {
			userOrders = append(userOrders, order)
		}
	}
	return userOrders
}

// GetTrades 获取用户交易记录
func (te *TradingEngine) GetTrades(userID string) []*Trade {
	var userTrades []*Trade
	for _, trade := range te.trades {
		if trade.UserID == userID {
			userTrades = append(userTrades, trade)
		}
	}
	return userTrades
}

// GetMarketData 获取市场数据
func (te *TradingEngine) GetMarketData() []*Stock {
	var stocks []*Stock
	for _, stock := range te.stocks {
		stocks = append(stocks, stock)
	}
	return stocks
}

// RecalculatePortfolioValues 重新计算所有投资组合的价值
func (te *TradingEngine) RecalculatePortfolioValues() {
	for userID, portfolio := range te.portfolios {
		portfolio.TotalValue = portfolio.Cash
		for symbol, holding := range portfolio.Holdings {
			if stock, exists := te.stocks[symbol]; exists {
				holding.MarketValue = float64(holding.Quantity) * stock.Price
				holding.UnrealizedPL = holding.MarketValue - holding.TotalCost
				portfolio.TotalValue += holding.MarketValue
				portfolio.Holdings[symbol] = holding
			}
		}
		portfolio.UpdatedAt = time.Now()
		te.portfolios[userID] = portfolio
	}
}

// GetOpenOrders 获取所有挂单（状态为PENDING）
func (te *TradingEngine) GetOpenOrders() []*Order {
	var result []*Order
	for _, order := range te.orders {
		if order.Status == OrderStatusPending {
			result = append(result, order)
		}
	}
	return result
}

// GetOpenOrdersByUser 获取指定用户的挂单
func (te *TradingEngine) GetOpenOrdersByUser(userID string) []*Order {
	var result []*Order
	for _, order := range te.orders {
		if order.UserID == userID && order.Status == OrderStatusPending {
			result = append(result, order)
		}
	}
	return result
}

// GetOpenOrdersBySymbol 获取指定股票代码的挂单
func (te *TradingEngine) GetOpenOrdersBySymbol(symbol string) []*Order {
	var result []*Order
	for _, order := range te.orders {
		if order.Symbol == symbol && order.Status == OrderStatusPending {
			result = append(result, order)
		}
	}
	return result
}
