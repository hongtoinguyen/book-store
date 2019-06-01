package database

import "gopkg.in/mgo.v2/bson"

//Order ...
type Order struct {
	ID          bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Note        string          `json:"note,omitempty"`
	OrderDate   uint64          `json:"order_date,omitempty"`
	ShippingFee uint64          `json:"shipping_fee,omitempty"`
	Status      uint64          `json:"status,omitempty"`
	BookIDs     []bson.ObjectId `json:"book_id,omitempty" bson:"_book_id,omitempty"`
	Payment     Payment         `json:"payment,omitempty"`
}

// //CreateCart ...
// func CreateCart(bookID bson.ObjectId) (Cart, error) {
// 	cart := Cart{}
// 	var err error
// 	return cart, err
// }
