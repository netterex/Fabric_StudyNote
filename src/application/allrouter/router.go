package allrouter

import (
	"fabric-houserental/application/controller/contract"
	"fabric-houserental/application/controller/credit"
	"fabric-houserental/application/controller/housemng"
	"fabric-houserental/application/controller/police"
	"fabric-houserental/application/controller/transaction"
	"fabric-houserental/application/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {

	police_grp := router.Group("/police")
	housemng_grp := router.Group("/housemng")
	credit_grp := router.Group("/credit")
	contract_grp := router.Group("/contract")
	transaction_grp := router.Group("/transaction")
	user_grp := router.Group("/user")

	police.Router(police_grp)

	housemng.Router(housemng_grp)

	credit.Router(credit_grp)

	contract.Router(contract_grp)

	transaction.Router(transaction_grp)

	user.Router(user_grp)

}
