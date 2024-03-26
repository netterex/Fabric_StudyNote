package contract

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fabric-houserental/application/sdk/credit"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Query(ctx *gin.Context) {

	id_contract := ctx.Query("id_contract")

	chaincode_name := "contractCc"
	fcn := "query"
	args := [][]byte{[]byte(id_contract)}

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

	id_contract := ctx.PostForm("id_contract")
	id_tenant := ctx.PostForm("id_tenant")
	id_house := ctx.PostForm("id_house")

	fileHeader, err := ctx.FormFile("img_contract")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "upload image failed",
		})
	}

	// hash image
	sha256Hash := sha256.New()

	file, err := fileHeader.Open()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "upload image failed",
		})
	}

	_, err = io.Copy(sha256Hash, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "upload image failed",
		})
	}

	sumByte := sha256Hash.Sum(nil)
	hash_contractImg := hex.EncodeToString(sumByte)

	// todo donot allow the same hash to be chained

	chaincode_name := "contractCc"
	fcn := "set"
	args := [][]byte{[]byte(id_contract), []byte(id_tenant), []byte(id_house), []byte(hash_contractImg)}

	_, err = credit.ChannelExecute(chaincode_name, fcn, args)
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
