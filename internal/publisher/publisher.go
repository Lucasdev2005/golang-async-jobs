package main

import (
	"fmt"

	"github.com/Lucasdev2005/golang-async-jobs/internal/core/database"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/entity"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/enums"
	"github.com/Lucasdev2005/golang-async-jobs/internal/core/rabbitMq"
	"github.com/Lucasdev2005/golang-async-jobs/internal/publisher/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	rabbitMq.ConnectionRabbitMq()
	rabbitMq.InitTransfers()

	database.Connect()
	defer database.Close()

	r := gin.Default()

	s := 1

	fmt.Println(s)
	r.POST("api/usuario/:id/transfer", func(ctx *gin.Context) {
		processRequest(ctx, controller.PublishTransfer)
	})
	r.Run()
}

func makeRequest(ctx *gin.Context) entity.Request {
	return entity.Request{
		Body:          ctx.BindJSON,
		GetParam:      ctx.Param,
		GetQueryParam: ctx.Query,
		GetHeader:     ctx.Request.Header.Get,
	}
}

func processRequest(
	context *gin.Context,
	fn func(request entity.Request) (interface{}, *entity.Error),
) {
	request := makeRequest(context)

	result, err := fn(request)

	if err.ErrorCode == 0 {
		context.JSON(enums.HttpStatusOK, result)
	} else {
		context.JSON(err.ErrorCode, err.Message)
	}
}
