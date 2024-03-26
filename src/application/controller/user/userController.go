package user

import (
	"encoding/json"
	"fabric-houserental/application/sdk/credit"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {

	role := ctx.PostForm("role")
	username := ctx.PostForm("username")
	pwd := ctx.PostForm("pwd")

	chaincode_name := "userCc"
	fcn := "query"
	args := [][]byte{
		[]byte(role),
		[]byte(username),
		[]byte(pwd)}

	resp, err := credit.ChannelQuery(chaincode_name, fcn, args)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "login failed",
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

func Register(ctx *gin.Context) {

	role := ctx.PostForm("role")
	username := ctx.PostForm("username")
	pwd := ctx.PostForm("pwd")
	age := ctx.PostForm("age")
	sex := ctx.PostForm("sex")

	chaincode_name := "userCc"
	fcn := "set"
	args := [][]byte{
		[]byte(role),
		[]byte(username),
		[]byte(pwd),
		[]byte(age),
		[]byte(sex)}

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
