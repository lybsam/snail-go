package util

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"math/rand"
	"net/http"
	"snail/app"
	"snail/pkg/e"
	"strconv"
	"time"
)

//uuid
func SQLUUID() string {
	return generateId()[8:24]
}

func generateId() string {
	chars := randStr(16)
	str := strconv.FormatInt(time.Now().Unix(), 10) + chars
	return EncodeMD5(str)
}
func randStr(n int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = chars[rand.Int63()%int64(len(chars))]
	}
	return string(b)
}

func GetAnalysis(ctx *gin.Context, mt int, args ...string) (app.Gin, []string) {
	at := app.Gin{C: ctx}
	valid := validation.Validation{}
	result := make([]string, len(args))
	switch mt {
	case 1:
		result[0] = ctx.Param(args[0])
		valid.Required(result[0], args[0]).Message(result[0] + "不能为空！")
		if len(args) > 1 {
			for i := 1; i < len(args); i++ {
				result[i] = ctx.Query(args[i])
				valid.Required(result[i], args[i]).Message(result[i] + "不能为空！")
			}
		}
		break
	case 2:
		for i := 0; i < len(args); i++ {
			result[i] = ctx.Query(args[i])
			valid.Required(result[i], args[i]).Message(result[i] + "不能为空！")
		}
		break
	case 3:
		buf := make([]byte, 4096)
		n, _ := ctx.Request.Body.Read(buf)
		for i := 0; i < len(args); i++ {
			result[i] = gjson.GetBytes(buf[:n], args[i]).String()
			valid.Required(result[i], args[i]).Message(result[i] + "不能为空！")
		}
		break
	}
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		at.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return at, nil
	}
	return at, result

}
