package trading_engine

import (
	"github.com/shopspring/decimal"
)

type Order struct {
	orderId    string
	price      decimal.Decimal
	quantity   decimal.Decimal
	createTime int64
	index      int
}

func (o *Order) GetIndex() int {
	return o.index
}

func (o *Order) SetIndex(index int) {
	o.index = index
}

func (o *Order) SetQuantity(qnt decimal.Decimal) {
	o.quantity = qnt
}

func (o *Order) GetUniqueId() string {
	return o.orderId
}

func (o *Order) GetPrice() decimal.Decimal {
	return o.price
}

func (o *Order) GetQuantity() decimal.Decimal {
	return o.quantity
}

func (o *Order) GetCreateTime() int64 {
	return o.createTime
}

// 这个方法留在具体的 ask/bid 队列中实现
// func (o *Order) Less() {}

type AskItem struct {
	Order
}

func (a *AskItem) Less(other QueueItem) bool {
	//价格优先，时间优先原则
	//价格低的在最上面
	return (a.price.Cmp(other.(*AskItem).price) == -1) || (a.price.Cmp(other.(*AskItem).price) == 0 && a.createTime < other.(*AskItem).createTime)
}

func (a *AskItem) GetAskOrBid() string {
	return "ask"
}

type BidItem struct {
	Order
}

func (a *BidItem) Less(other QueueItem) bool {
	//价格优先，时间优先原则
	//价格高的在最上面
	return (a.price.Cmp(other.(*BidItem).price) == 1) || (a.price.Cmp(other.(*BidItem).price) == 0 && a.createTime < other.(*BidItem).createTime)
}

func (a *BidItem) GetAskOrBid() string {
	return "bid"
}

func NewAskItem(uniqId string, price, quantity decimal.Decimal, createTime int64) *AskItem {
	return &AskItem{Order{
		orderId:    uniqId,
		price:      price,
		quantity:   quantity,
		createTime: createTime,
	}}
}

func NewBidItem(uniqId string, price, quantity decimal.Decimal, createTime int64) *BidItem {
	return &BidItem{Order{
		orderId:    uniqId,
		price:      price,
		quantity:   quantity,
		createTime: createTime,
	}}
}
