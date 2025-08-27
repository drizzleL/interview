package main

type LfuCache struct {
	dict     map[int]*ValNode
	capacity int
	freqHead *FreqNode
}

func NewLfuCache(c int) *LfuCache {
	return &LfuCache{
		dict:     make(map[int]*ValNode),
		capacity: c,
		freqHead: newFreqNode(0),
	}
}

func (c *LfuCache) evict() {
	if len(c.dict) == 0 {
		return
	}
	freq := c.freqHead.Next
	item := freq.ValHead.Next // first item
	delete(c.dict, item.Key)
	item.delete()
	freq.clear()
}

func (c *LfuCache) Get(key int) (int, bool) {
	item, ok := c.dict[key]
	if !ok {
		return 0, false
	}
	item.upgrade()
	return item.Val, true
}

func (c *LfuCache) Put(key, val int) {
	item, ok := c.dict[key]
	if !ok {
		if len(c.dict) == c.capacity {
			c.evict()
		}
		freq := c.freqHead.Next
		if freq.Freq != 1 {
			freq = newFreqNode(1)
			freq.Prev = c.freqHead
			freq.Next = c.freqHead.Next
			if c.freqHead.Next != nil {
				c.freqHead.Next.Prev = freq
			}
			c.freqHead.Next = freq
		}
		item = &ValNode{
			Key:      key,
			Val:      val,
			FreqNode: freq,
		}
		freq.addVal(item)
		c.dict[key] = item
		return
	}
	item.Val = val
	item.upgrade()
}

func (item *ValNode) upgrade() {
	freq := item.FreqNode
	nextFreq := freq.Next
	if nextFreq == nil || nextFreq.Freq != freq.Freq+1 {
		nextFreq = newFreqNode(freq.Freq + 1)
		nextFreq.Prev = freq
		nextFreq.Next = freq.Next
		if freq.Next != nil {
			freq.Next.Prev = nextFreq
		}
		freq.Next = nextFreq
	}
	item.delete() // delete from old relation
	freq.clear()
	nextFreq.addVal(item)
	item.FreqNode = nextFreq

}

type ValNode struct {
	Prev, Next *ValNode
	Key        int
	Val        int
	FreqNode   *FreqNode
}

func (n *ValNode) delete() {
	prev, next := n.Prev, n.Next
	prev.Next = next
	next.Prev = prev
}

type FreqNode struct {
	Prev, Next       *FreqNode
	Freq             int
	ValHead, ValTail *ValNode
}

func (n *FreqNode) addVal(val *ValNode) {
	tail := n.ValTail
	tailPrev := tail.Prev
	val.Prev = tailPrev
	val.Next = tail
	tailPrev.Next = val
	tail.Prev = val
}

func (n *FreqNode) clear() {
	if n.ValHead.Next != n.ValTail { // not empty
		return
	}
	prev, next := n.Prev, n.Next
	prev.Next = next
	if next != nil {
		next.Prev = prev
	}
}

func newFreqNode(freq int) *FreqNode {
	h, t := &ValNode{}, &ValNode{}
	h.Next = t
	t.Prev = h
	return &FreqNode{
		Freq:    freq,
		ValHead: h,
		ValTail: t,
	}
}
