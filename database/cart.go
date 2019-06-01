package database

import "gopkg.in/mgo.v2/bson"

//Cart ...
type Cart struct {
	ID     bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	BookID bson.ObjectId `json:"book_id,omitempty" bson:"_book_id,omitempty" form:"book_id,omitempty"`
}

//CreateCart ...
func CreateCart(bookID bson.ObjectId) (Cart, error) {
	cart := Cart{}
	var err error
	return cart, err
}

//Update ...
func (cart *Cart) Update(newCart Cart) error {
	return nil
}

func (cart *Cart) Delete() error {
	return nil
}
