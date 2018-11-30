package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:user_name/videos")
	router.GET("/user/:user_name/videos/:vid-id")
	router.DELETE("/user/:user_name/videos/:vid-id")
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}
