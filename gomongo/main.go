package main

import (
	"gopkg.in/mgo.v2"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController()
}