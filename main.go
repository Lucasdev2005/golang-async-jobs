package main

import (
	"github.com/Lucasdev2005/golang-async-jobs/internal/controller"
	"github.com/Lucasdev2005/golang-async-jobs/internal/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/enums"
	"github.com/Lucasdev2005/golang-async-jobs/internal/rabbitMq"
	"github.com/Lucasdev2005/golang-async-jobs/internal/types"
	"github.com/gin-gonic/gin"
)

func main() {
	rabbitMq.ConnectionRabbitMq()
	rabbitMq.InitTransfers()
	r := gin.Default()
	database.Connect()

	r.POST("api/usuario/:id/transfer", func(ctx *gin.Context) {
		processRequest(ctx, controller.CreateTransfer)
	})
	r.Run()
}

func makeRequest(ctx *gin.Context) types.Request {
	return types.Request{
		Body:          ctx.BindJSON,
		GetParam:      ctx.Param,
		GetQueryParam: ctx.Query,
		GetHeader:     ctx.Request.Header.Get,
	}
}

func processRequest(
	context *gin.Context,
	fn func(request types.Request) (interface{}, *types.Error),
) {
	request := makeRequest(context)

	result, err := fn(request)

	if err == nil {
		context.JSON(enums.HttpStatusOK, result)
	} else {
		context.JSON(err.ErrorCode, err.Message)
	}
}
