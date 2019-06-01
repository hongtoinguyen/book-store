package database

import "gopkg.in/mgo.v2/bson"

//Customer ...
type Customer struct {
	ID            bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Email         string          `json:"email,omitempty" form:"email,omitempty"`
	Password      string          `json:"-" form:"-" bson:"password,omitempty"`
	FirstName     string          `json:"first_name,omitempty" form:"first_name,omitempty"`
	LastName      string          `json:"last_name,omitempty" form:"last_name,omitempty"`
	Phone         string          `json:"phone,omitempty" form:"phone,omitempty"`
	Username      string          `json:"username,omitempty" form:"username,omitempty"`
	ShippingInfos []ShippingInfo  `json:"shipping_infos,omitempty"`
	Avatar        string          `json:"avatar,omitempty"`
	Carts         []Cart          `json:"carts,omitempty"`
	Orders        []bson.ObjectId `json:"orders,omitempty" bson:"orders,omitempty"`
	Token         string          `json:"token,omitempty" form:"token,omitempty"`
}

//Update ...
func (customer *Customer) Update() error {
	return nil
}
