package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	example "github.com/micro/examples/template/srv/proto/example"
	"github.com/julienschmidt/httprouter"
	ADMINLOGIN "MyShop/AdminLogin/proto/example"

	"github.com/micro/go-grpc"
	"github.com/astaxie/beego"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	beego.Info("Login")
	w.Header().Set("Content-Type","application/json")

	beego.Info("-----------")
	beego.Info(r.Body)
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 500)
		return
	}
	username := request["username"].(string)
	password := request["password"].(string)
	beego.Info(username,password)
	service := grpc.NewService()
	service.Init()
	// call the backend service
	exampleClient := ADMINLOGIN.NewExampleService("go.micro.srv.AdminLogin", service.Client())
	rsp, err := exampleClient.AdminLogin(context.TODO(), &ADMINLOGIN.Request{
		UserName:username,
		Password:password,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 500)
		return
	}
}