package main

import (
	"github.com/micro/go-log"
	"net/http"
	_ "MyShop/Router/models"
	"github.com/micro/go-web"
	"github.com/julienschmidt/httprouter"
	"MyShop/Router/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.Router"),
		web.Version("latest"),
		web.Address(":8080"),

	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// 使用路由中间件来映射页面
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("html"))

	router.POST("/admin/login",handler.Login)
	service.Handle("/", router)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
