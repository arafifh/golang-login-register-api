package main

import (
	authcontroller "github.com/arafifh/golang-login-register-api/controllers"
	"net/http"
	//"database/sql"
	//"github.com/gin-gonic/gin"
	//"errors"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
}