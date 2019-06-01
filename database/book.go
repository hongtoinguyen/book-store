package database

import "gopkg.in/mgo.v2/bson"

//Books ...
type Book struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Title         string        `json:"title,omitempty"`
	Description   string        `json:"description,omitempty"`
	Organization  string        `json:"organization,omitempty"`
	Inventory     uint64        `json:"inventory,omitempty"`
	FromTime      uint64        `json:"from_time,omitempty"`
	ToTime        uint64        `json:"to_time,omitempty"`
	Price         uint64        `json:"price,omitempty"`
	DiscountPrice uint64        `json:"discount_price,omitempty"`
	CategoryName  string        `json:"category_name,omitmpey"`
	Images        []string      `json:"images,omitempty"`
	Comments      []Comment     `json:"comments,omitempty"`
}

// //CreateCart ...
// func CreateCart(bookID bson.ObjectId) (Cart, error) {
// 	cart := Cart{}
// 	var err error
// 	return cart, err
// }

//Update ...
func (book *Book) Update(newBook Book) error {
	return nil
}

func (book *Book) Delete() error {
	return nil
}
