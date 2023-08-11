package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-demo/model"
)

func init() {
	baseDBModel := model.BaseDBModel{}
	db := baseDBModel.Connect()
	if db == nil {
		fmt.Println("连接失败")
	}
}

func main() {
	//go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
