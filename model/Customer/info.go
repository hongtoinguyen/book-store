package Customer

import (
	"book-store/database"

	"gopkg.in/mgo.v2/bson"
)

//GetCustomerByID
func GetCustomerByID(customerID string) (database.Customer, error) {
	customer := database.Customer{}
	var err error
	customer.ID = bson.ObjectIdHex(customerID)
	err = database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{}).One(&customer)
	if err != nil {
		return customer, err
	}
	return customer, err
}
