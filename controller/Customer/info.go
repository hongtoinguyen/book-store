package Customer

import (
	"book-store/controller/SetUp"
	"book-store/model/Customer"
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetCustomerByID ...
func GetCustomerByID(c *gin.Context) {
	customerID := c.Param("customer_id")
	customer, err := Customer.GetCustomerByID(customerID)
	if err != nil {
		c.JSON(SetUp.ErrorMap[fmt.Sprintf("%s", err)], gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	c.JSON(200, customer)
}

//UpdateCustomerInfo
func UpdateCustomerInfo(c *gin.Context) {

}
