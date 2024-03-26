package credit

import (
	"fabric-houserental/application/sdk/credit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Query(ctx *gin.Context) {

	id_card := ctx.Query("id_card")

	chaincode_name := "creditCc"
	fcn := "query"
	args := [][]byte{[]byte(id_card)}

	resp, err := credit.ChannelQuery(chaincode_name, fcn, args)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "query failed",
			"data": nil,
		})
	} else {
		map_data := map[string]interface{}{
			"creditLevel": string(resp.Payload),
			"id_card":     id_card,
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "query success",
			"data": map_data,
		})
	}

}

func Set(ctx *gin.Context) {

	id_card := ctx.PostForm("id_card")
	creditLevel := ctx.PostForm("creditLevel")

	chaincode_name := "creditCc"
	fcn := "set"
	args := [][]byte{[]byte(id_card), []byte(creditLevel)}

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
