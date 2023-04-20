package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPing(c *gin.Context) {
	var pinged = Ping{
		Response: Response{Status: true, Message: "success"}, Device: Device{DeviceType: "XLR8", DeviceFirmware: "1.0.0"},
		Preset: Preset{OnPeriod: 2, OffPeriod: 4, OnTime: 6, OnDate: "Date", Repeats: 1, Power: 100},
	}
	c.IndentedJSON(http.StatusOK, pinged)
}

func setTime(c *gin.Context) {
	var time RTC
	if err := c.BindJSON(&time); err != nil {
		fmt.Println(err)
		var response = Response{Status: false, Message: "Error"}
		c.IndentedJSON(http.StatusMethodNotAllowed, response)
		return
	}
	fmt.Println(time.RtcTime)
	var response = Response{Status: true, Message: "Success"}
	c.IndentedJSON(http.StatusOK, response)

}

func setRouterConfig(c *gin.Context) {
	var router RouterConfig
	if err := c.BindJSON(&router); err != nil {
		fmt.Println(err)
		var response = APResponse{IpPreset: IpPreset{Ip: "192.168.4.x"}, Response: Response{Status: false, Message: "Error"}}
		c.IndentedJSON(http.StatusMethodNotAllowed, response)
		return
	}
	fmt.Println(router.Ssid)
	var response = APResponse{IpPreset: IpPreset{Ip: "192.168.4.x"}, Response: Response{Status: true, Message: "Success"}}
	c.IndentedJSON(http.StatusOK, response)
}

func setAPConfig(c *gin.Context) {
	var router APConfig
	if err := c.BindJSON(&router); err != nil {
		fmt.Println(err)
		var response = APResponse{IpPreset: IpPreset{Ip: "192.168.4.x"}, Response: Response{Status: false, Message: "Error"}}
		c.IndentedJSON(http.StatusMethodNotAllowed, response)
		return
	}
	fmt.Println(router.Ssid)
	var response = APResponse{IpPreset: IpPreset{Ip: "192.168.4.x"}, Response: Response{Status: true, Message: "Success"}}
	c.IndentedJSON(http.StatusOK, response)

}
func setPreset(c *gin.Context) {
	var preset Preset
	if err := c.BindJSON(&preset); err != nil {
		fmt.Println(err)
		var response = Response{Status: false, Message: "Error"}
		c.IndentedJSON(http.StatusMethodNotAllowed, response)
		return
	}
	fmt.Println(preset.Power)
	var response = Response{Status: true, Message: "Success"}
	c.IndentedJSON(http.StatusOK, response)
}
