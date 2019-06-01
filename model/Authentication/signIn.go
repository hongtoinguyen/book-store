package Authentication

import (
	"book-store/database"
	"book-store/utils"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

//SignIn ...
func SignIn(email string, password string) (database.Customer, error) {
	customer := database.Customer{}
	var err error
	if !(checkAuthData(email, password)) {
		err = fmt.Errorf("%s", "Email or password is invalid")
		return customer, err
	}
	database.GetMongoDB().C(database.COL_CUSTOMERS).Find(bson.M{"email": email}).One(&customer)

	check, err := utils.Compare(customer.Password, password)

	if err != nil {
		err = fmt.Errorf("%s", "Password is not right")
		return customer, err
	}

	if check == false {
		err = fmt.Errorf("%s", "Password is not right")
		return customer, err
	}

	return customer, err
}
