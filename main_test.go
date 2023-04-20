package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPing(t *testing.T) {
	mockResponse := `{
    "Response": {
        "status": true,
        "message": "success"
    },
    "Device": {
        "device_type": "XLR8",
        "device_firmware": "1.0.0"
    },
    "Preset": {
        "on_period": 2,
        "off_period": 4,
        "on_time": 6,
        "on_date": "Date",
        "repeats": 1,
        "power": 100
    }
}`
	r := SetUpRouter()
	r.GET("/ping", getPing)
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSetRTC(t *testing.T) {
	mockResponse := `{
    "status": true,
    "message": "Success"
}`
	r := SetUpRouter()
	r.POST("/setTime", setTime)
	var body = RTC{RtcTime: "time"}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/setTime", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSetRouterConfig(t *testing.T) {
	mockResponse := `{
    "response": {
        "status": true,
        "message": "Success"
    },
    "ipPreset": {
        "ip": "192.168.4.x"
    }
}`
	r := SetUpRouter()
	r.POST("/routerConfig", setRouterConfig)
	var body = RouterConfig{Ssid: "time", Password: "hamada"}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/routerConfig", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSetAPConfig(t *testing.T) {
	mockResponse := `{
    "response": {
        "status": true,
        "message": "Success"
    },
    "ipPreset": {
        "ip": "192.168.4.x"
    }
}`
	r := SetUpRouter()
	r.POST("/apConfig", setRouterConfig)
	var body = APConfig{Ssid: "time", Password: "hamada"}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/apConfig", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSetPreset(t *testing.T) {
	mockResponse := `{
    "status": true,
    "message": "Success"
}`
	r := SetUpRouter()
	r.POST("/setPreset", setPreset)
	var body = Preset{
		OnPeriod:  1,
		OffPeriod: 1,
		OnTime:    1,
		OnDate:    "sada",
		Repeats:   1,
		Power:     1,
	}
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/setPreset", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}
