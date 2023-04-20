package main

type Device struct {
	DeviceType     string `json:"device_type" binding:"required"`
	DeviceFirmware string `json:"device_firmware" binding:"required"`
}

type Preset struct {
	OnPeriod  int    `json:"on_period" binding:"required"`
	OffPeriod int    `json:"off_period" binding:"required"`
	OnTime    int    `json:"on_time" binding:"required"`
	OnDate    string `json:"on_date" binding:"required"`
	Repeats   int    `json:"repeats" binding:"required"`
	Power     int    `json:"power" binding:"required"`
}

type Response struct {
	Status  bool   `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type IpPreset struct {
	Ip string `json:"ip" binding:"required"`
}

type RouterConfig struct {
	Ssid     string `json:"ssid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type APConfig struct {
	Ssid     string `json:"ssid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RTC struct {
	RtcTime string `json:"rtc_Time" binding:"required"`
}

type Ping struct {
	Response Response `json:"Response" binding:"required"`
	Device   Device   `json:"Device" binding:"required"`
	Preset   Preset   `json:"Preset" binding:"required"`
}

type APResponse struct {
	Response Response `json:"response" binding:"required"`
	IpPreset IpPreset `json:"ipPreset" binding:"required"`
}
