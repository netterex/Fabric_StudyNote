package housemng

import (
	"encoding/json"
	"fabric-houserental/application/sdk/housemng"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Query(ctx *gin.Context) {

	id_house := ctx.Query("id_house")

	chaincode_name := "housemngCc"
	fcn := "query"
	args := [][]byte{[]byte(id_house)}

	resp, err := housemng.ChannelQuery(chaincode_name, fcn, args)

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

	id_house := ctx.PostForm("id_house")
	owner := ctx.PostForm("owner")
	addr_house := ctx.PostForm("addr_house")
	type_house := ctx.PostForm("type_house")

	chaincode_name := "housemngCc"
	fcn := "set"
	args := [][]byte{[]byte(id_house), []byte(owner), []byte(addr_house), []byte(type_house)}

	_, err := housemng.ChannelExecute(chaincode_name, fcn, args)
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
