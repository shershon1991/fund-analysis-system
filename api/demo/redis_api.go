/**
 * @Description redis验证使用
 **/
package demo

import (
	"context"
	"github.com/gin-gonic/gin"
	"shershon1991/fund-analye-system/global"
	"shershon1991/fund-analye-system/model/response"
	"time"
)

// 验证redis
func RdTest(ctx *gin.Context) {
	method, _ := ctx.GetQuery("type")
	var result string
	var err error
	switch method {
	case "get":
		result, err = global.GvaRedis.Get(context.TODO(), "test").Result()
	case "set":
		result, err = global.GvaRedis.Set(context.TODO(), "test", "hello word!", time.Hour).Result()
	}
	if err != nil {
		response.Error(ctx, err.Error())
		return
	}
	response.OkWithData(ctx, result)
}
