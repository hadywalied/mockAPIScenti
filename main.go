package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/ping", getPing)
	router.POST("/setTime", setTime)
	router.POST("/routerConfig", setRouterConfig)
	router.POST("/apConfig", setAPConfig)
	router.POST("/setPreset", setPreset)

	err := router.Run("localhost:5001")
	if err != nil {
		return
	}
}
func Fooer(input int) string {
	isfoo := (input % 3) == 0
	if isfoo {
		return "Foo"
	}
	return strconv.Itoa(input)
}
