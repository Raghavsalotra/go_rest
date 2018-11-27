package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Raghavsalotra/go_rest/config"

	"github.com/Raghavsalotra/go_rest/handler"
	appLog "github.com/Raghavsalotra/go_rest/log"
	m "github.com/Raghavsalotra/go_rest/middleware"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	appLog.Init(os.Stdout, os.Stdout)
	apiNS := "/api"

	router.Handler("GET", apiNS+"/user", m.Adapt(handler.GetUser, m.AccessLog(), m.Authorize()))
	router.Handler("POST", apiNS+"/user", m.Adapt(handler.SaveUser, m.AccessLog()))
	router.Handler("GET", apiNS+"/username", m.Adapt(handler.UserNameExists, m.AccessLog()))

	router.Handler("POST", apiNS+"/user/authorize", m.Adapt(handler.SignIn, m.AccessLog()))
	router.Handler("DELETE", apiNS+"/user/authorize", m.Adapt(handler.SignOut, m.AccessLog(), m.Authorize()))

	router.Handler("GET", apiNS+"/todo", m.Adapt(handler.GetTodo, m.AccessLog(), m.Authorize()))
	router.Handler("GET", apiNS+"/todo/user", m.Adapt(handler.GetUserTodos, m.AccessLog(), m.Authorize()))
	router.Handler("POST", apiNS+"/todo", m.Adapt(handler.PostTodo, m.AccessLog(), m.Authorize()))
	router.Handler("PUT", apiNS+"/todo", m.Adapt(handler.PutTodo, m.AccessLog(), m.Authorize()))
	router.Handler("DELETE", apiNS+"/todo", m.Adapt(handler.DeleteTodo, m.AccessLog(), m.Authorize()))
	appPort := 8080
	if p, ok := config.GetInt("app_port"); ok {
		appPort = p
	}
	if err := http.ListenAndServe(":"+strconv.Itoa(appPort), router); err != nil {
		log.Fatal("Unable to start server", err)
	}

}
