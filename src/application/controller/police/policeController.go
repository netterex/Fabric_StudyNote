package police

import (
	"encoding/json"
	"fabric-houserental/application/sdk/police"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Query(ctx *gin.Context) {

	id_card := ctx.Query("id_card")

	chaincode_name := "policeCc"
	fcn := "query"
	args := [][]byte{[]byte(id_card)}

	resp, err := police.ChannelQuery(chaincode_name, fcn, args)

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

	id_card := ctx.PostForm("id_card")
	name := ctx.PostForm("name")
	age := ctx.PostForm("age")
	addr := ctx.PostForm("addr")

	chaincode_name := "policeCc"
	fcn := "set"
	args := [][]byte{[]byte(id_card), []byte(name), []byte(age), []byte(addr)}

	_, err := police.ChannelExecute(chaincode_name, fcn, args)
	if error.Error(err) == "Client Status Code: (5) TIMEOUT. Description: request timed out or been cancelled" || err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "set success",
		})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "set failed",
		})
	}
}
