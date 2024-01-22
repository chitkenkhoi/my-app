package model

import (
	"time"
)

type Payload struct {
	Id  string `json:"id"`
	Iat int64  `json:"iat"`
}
type User struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"  bson:"name,omitempty"`
	Price int    `json:"price" bson:"price,omitempty"`
	Unit  string `json:"Unit"  bson:"unit,omitempty"`
	Size  string `json:"Size"  bson:"size,omitempty"`
}
type Order struct {
	ID       string        `json:"id,omitempty"`
	Buyer    string        `json:"buyer"    bson:"buyer,omitempty"`
	Date     time.Time     `json:"date"     bson:"date,omitempty"`
	Discount int           `json:"discount" bson:"discount,omitempty"`
	Price    int           `json:"price"    bson:"price,omitempty"`
	Status1  bool          `json:"status1"  bson:"status1"`
	Status2  bool          `json:"status2"  bson:"status2"`
	ItemList []ItemInOrder `json:"itemList" bson:"itemList,omitempty"`
}
type ItemInOrder struct {
	Item_id string `json:"item_id" bson:"item_id,omitempty"`
	Amount  int    `json:"amount"  bson:"amount,omitempty"`
}
