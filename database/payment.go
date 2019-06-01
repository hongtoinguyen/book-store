package database

import "gopkg.in/mgo.v2/bson"

//Payment ...
type Payment struct {
	ID     bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" form:"id,omitempty"`
	Method string        `json:"method,omitempty"`
}

//CreatePayment ...
func CreatePayment(method string) (Payment, error) {
	payment := Payment{}
	var err error
	return payment, err
}

//Update ...
func (payment *Payment) Update(newPayment Payment) error {
	return nil
}

func (payment *Payment) Delete() error {
	return nil
}
