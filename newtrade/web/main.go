package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"newtrade"
)

type WebServer struct {
	engine *newtrade.TradingEngine
	tmpl   *template.Template
}

func main() {
	// 创建交易引擎
	engine := newtrade.NewTradingEngine()

	// 初始化一些示例股票
	initializeSampleData(engine)

	// 创建web服务器
	server := &WebServer{
		engine: engine,
		tmpl:   template.Must(template.ParseGlob("templates/*.html")),
	}

	// 设置路由
	http.HandleFunc("/", server.handleHome)
	http.HandleFunc("/api/stocks", server.handleGetStocks)
	http.HandleFunc("/api/place-order", server.handlePlaceOrder)
	http.HandleFunc("/api/cancel-order", server.handleCancelOrder)
	http.HandleFunc("/api/orders", server.handleGetOrders)
	http.HandleFunc("/api/portfolio", server.handleGetPortfolio)
	http.HandleFunc("/api/trades", server.handleGetTrades)
	http.HandleFunc("/api/open-orders", server.handleGetOpenOrders)

	// 静态文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("股票交易系统Web界面启动在 http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func initializeSampleData(engine *newtrade.TradingEngine) {
	// 添加示例股票
	stocks := []*newtrade.Stock{
		{
			Symbol:    "AAPL",
			Name:      "Apple Inc.",
			Price:     150.0,
			Change:    0.0,
			Volume:    1000000,
			MarketCap: 2500000000000,
			UpdatedAt: time.Now(),
		},
		{
			Symbol:    "GOOGL",
			Name:      "Alphabet Inc.",
			Price:     2800.0,
			Change:    0.0,
			Volume:    500000,
			MarketCap: 1800000000000,
			UpdatedAt: time.Now(),
		},
		{
			Symbol:    "TSLA",
			Name:      "Tesla Inc.",
			Price:     800.0,
			Change:    0.0,
			Volume:    800000,
			MarketCap: 800000000000,
			UpdatedAt: time.Now(),
		},
		{
			Symbol:    "MSFT",
			Name:      "Microsoft Corp.",
			Price:     350.0,
			Change:    0.0,
			Volume:    800000,
			MarketCap: 2600000000000,
			UpdatedAt: time.Now(),
		},
		{
			Symbol:    "AMZN",
			Name:      "Amazon.com Inc.",
			Price:     180.0,
			Change:    0.0,
			Volume:    1200000,
			MarketCap: 1900000000000,
			UpdatedAt: time.Now(),
		},
	}

	for _, stock := range stocks {
		engine.AddStock(stock)
	}

	// 初始化默认用户资金
	engine.EnsurePortfolio("default_user", 1_000_000)
}

func (s *WebServer) handleHome(w http.ResponseWriter, r *http.Request) {
	s.tmpl.ExecuteTemplate(w, "index.html", nil)
}

func (s *WebServer) handleGetStocks(w http.ResponseWriter, r *http.Request) {
	stocks := s.engine.GetMarketData()
	// 按symbol排序，稳定输出
	for i := 0; i < len(stocks); i++ {
		for j := i + 1; j < len(stocks); j++ {
			if stocks[i].Symbol > stocks[j].Symbol {
				stocks[i], stocks[j] = stocks[j], stocks[i]
			}
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    stocks,
	})
}

func (s *WebServer) handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Symbol   string  `json:"symbol"`
		Type     string  `json:"type"`
		Quantity int64   `json:"quantity"`
		Price    float64 `json:"price"`
		UserID   string  `json:"userId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 创建订单
	var orderType newtrade.OrderType
	if req.Type == "BUY" {
		orderType = newtrade.OrderTypeBuy
	} else {
		orderType = newtrade.OrderTypeSell
	}

	order := &newtrade.Order{
		Symbol:   req.Symbol,
		Type:     orderType,
		Quantity: req.Quantity,
		Price:    req.Price,
		UserID:   req.UserID,
	}

	if err := s.engine.PlaceOrder(order); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	log.Printf("[DEBUG] Engine@%p after PlaceOrder, total open orders: %d", s.engine, len(s.engine.GetOpenOrders()))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    order,
	})
}

func (s *WebServer) handleCancelOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		OrderID string `json:"orderId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := s.engine.CancelOrder(req.OrderID); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

func (s *WebServer) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "default_user"
	}

	orders := s.engine.GetOrders(userID)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    orders,
	})
}

func (s *WebServer) handleGetPortfolio(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "default_user"
	}

	portfolio, err := s.engine.GetPortfolio(userID)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    portfolio,
	})
}

func (s *WebServer) handleGetTrades(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "default_user"
	}

	trades := s.engine.GetTrades(userID)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    trades,
	})
}

func (s *WebServer) handleGetOpenOrders(w http.ResponseWriter, r *http.Request) {
	log.Printf("[DEBUG] Engine@%p before GetOpenOrders, total open orders: %d", s.engine, len(s.engine.GetOpenOrders()))
	openOrders := s.engine.GetOpenOrders()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"data":    openOrders,
	})
}
