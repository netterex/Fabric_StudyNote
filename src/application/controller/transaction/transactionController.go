package transaction

import (
	"encoding/json"
	"fabric-houserental/application/sdk/credit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Query(ctx *gin.Context) {

	id_order := ctx.Query("id_order")
	period := ctx.Query("period")

	chaincode_name := "transactionCc"
	fcn := "query"
	args := [][]byte{[]byte(id_order), []byte(period)}

	resp, err := credit.ChannelQuery(chaincode_name, fcn, args)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "query failed",
			"data": nil,
		})
	} else {
		map_data := make(map[string]interface{})
		json.Unmarshal([]byte(string(resp.Payload)), &map_data)
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "query success",
			"data": map_data,
		})
	}
}

func Set(ctx *gin.Context) {

	id_order := ctx.PostForm("id_order")
	period := ctx.PostForm("period")
	id_landlord := ctx.PostForm("id_landlord")
	id_tenant := ctx.PostForm("id_tenant")
	name_landlord := ctx.PostForm("name_landlord")
	name_tenant := ctx.PostForm("name_tenant")
	id_house := ctx.PostForm("id_house")
	rent := ctx.PostForm("rent")
	time_tx := ctx.PostForm("time_tx")
	id_contract := ctx.PostForm("id_contract")
	desc := ctx.PostForm("desc")

	chaincode_name := "transactionCc"
	fcn := "set"
	args := [][]byte{
		[]byte(id_order),
		[]byte(period),
		[]byte(id_landlord),
		[]byte(id_tenant),
		[]byte(name_landlord),
		[]byte(name_tenant),
		[]byte(id_house),
		[]byte(rent),
		[]byte(time_tx),
		[]byte(id_contract),
		[]byte(desc)}

	_, err := credit.ChannelExecute(chaincode_name, fcn, args)
	if error.Error(err) == "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" || err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "set success",
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "set failed",
		})
	}
}
