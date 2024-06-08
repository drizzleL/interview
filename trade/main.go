package main

import (
	"container/heap"
	"log"
	"time"
)

func main() {
	tm := NewTradeManager()
	tm.Buy(&TradeItem{
		Price:  240,
		Amount: 200,
	})
	tm.Buy(&TradeItem{
		Price:  300,
		Amount: 20,
	})
	tm.Buy(&TradeItem{
		Price:  100,
		Amount: 2000,
	})

	tm.Sell(&TradeItem{
		Price:  1000,
		Amount: 20000,
	})
	tm.Sell(&TradeItem{
		Price:  100,
		Amount: 200,
	})
	tm.Sell(&TradeItem{
		Price:  300,
		Amount: 20,
	})

	select {}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TradeManager struct {
	Sells     TradeWay
	Buys      TradeWay
	CurrPrice int
}

func NewTradeManager() *TradeManager {
	tm := &TradeManager{
		Sells: TradeWay{
			PriceCmp: func(price1, price2 int) bool {
				return price1 < price2
			},
		},
		Buys: TradeWay{
			PriceCmp: func(price1, price2 int) bool {
				return price1 > price2
			},
		},
	}
	go func() {
		for range time.NewTicker(time.Millisecond * 200).C {
			tm.TryTrade()
		}
	}()
	return tm
}

func (tm *TradeManager) TryTrade() {
	if tm.Sells.Len() == 0 || tm.Buys.Len() == 0 {
		return
	}
	topSell := tm.Sells.TradeItems[0]
	topBuy := tm.Buys.TradeItems[0]
	if topSell.Price > topBuy.Price {
		return
	}
	dealAmount := min(topSell.Amount, topBuy.Amount)
	heap.Pop(&tm.Sells)
	heap.Pop(&tm.Buys)
	topSell.Amount -= dealAmount
	if topSell.Amount != 0 {
		heap.Push(&tm.Sells, topSell)
	}
	topBuy.Amount -= dealAmount
	if topBuy.Amount != 0 {
		heap.Push(&tm.Buys, topBuy)
	}
	tm.CurrPrice = topSell.Price
	log.Printf("deal made, price %d, amount %d", tm.CurrPrice, dealAmount)
}

func (m *TradeManager) Buy(item *TradeItem) {
	item.CreateTime = time.Now()
	heap.Push(&m.Buys, item)
}

func (m *TradeManager) Sell(item *TradeItem) {
	item.CreateTime = time.Now()
	heap.Push(&m.Sells, item)
}

type TradeHistory struct {
}

type TradeItem struct {
	TraderId   int
	Price      int
	Amount     int
	CreateTime time.Time
}

type TradeWay struct {
	TradeItems
	PriceCmp func(price1, price2 int) bool
}

func (tw *TradeWay) Less(i, j int) bool {
	a, b := tw.TradeItems[i], tw.TradeItems[j]
	if a.Price == b.Price {
		return a.CreateTime.Before(b.CreateTime)
	}
	return tw.PriceCmp(a.Price, b.Price)
}

type TradeItems []*TradeItem

func (pq TradeItems) Len() int { return len(pq) }

func (pq TradeItems) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TradeItems) Push(x any) {
	item := x.(*TradeItem)
	*pq = append(*pq, item)
}

func (pq *TradeItems) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
