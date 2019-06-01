package database

import "gopkg.in/mgo.v2/bson"

//ShippingInfo ...
type ShippingInfo struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Province string        `json:"province,omitempty"`
	City     string        `json:"city,omitempty"`
	District string        `json:"district,omitempty"`
	Ward     string        `json:"ward,omitempty"`
	Address  string        `json:"address,omitempty"`
}

// //CreateShippingInfo ...
// func CreateShippingInfo(bookID bson.ObjectId) (Cart, error) {
// 	cart := Cart{}
// 	var err error
// 	return cart, err
// }

//Update ...
func (shippingInfo *ShippingInfo) Update(newShippingInfo ShippingInfo) error {
	return nil
}

func (shippingInfo *ShippingInfo) Delete() error {
	return nil
}
