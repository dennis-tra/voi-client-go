package voi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Headers struct {
	Authority       string
	DeviceName      string
	UserAgent       string
	Timezone        string
	OperatingSystem string
	ClientTime      func() string
	DeviceId        func() string
	OsVersion       string
	CarrierName     string
	AppVersion      string
	Locale          string
	BatteryLevel    string
	RequestID       func() string
}

func NewDefaultHeaders() *Headers {
	return &Headers{
		Authority:       "api.voiapp.io",
		DeviceName:      "voi-go-client",
		UserAgent:       "voi-app/34 CFNetwork/1107.1 Darwin/19.0.0",
		Timezone:        "Europe/Berlin",
		OperatingSystem: "iOS",
		ClientTime:      func() string { return fmt.Sprintf("%.4f", float32(time.Now().UnixNano())/float32(time.Second)) },
		DeviceId:        uuid.New().String,
		OsVersion:       "13.1.2",
		CarrierName:     "Vodafone.de",
		AppVersion:      "5.5.2",
		Locale:          "en",
		BatteryLevel:    "0.5",
		RequestID:       uuid.New().String,
	}
}

func (hdr *Headers) fill(req *http.Request) {
	req.Header.Add("x-device-name", hdr.DeviceName)
	req.Header.Add("user-agent", hdr.UserAgent)
	req.Header.Add("x-timezone", hdr.Timezone)
	req.Header.Add("x-os", hdr.OperatingSystem)
	req.Header.Add("x-client-time", hdr.ClientTime())
	req.Header.Add("x-device-id", hdr.DeviceId())
	req.Header.Add("x-os-version", hdr.OsVersion)
	req.Header.Add("x-carrier-name", hdr.CarrierName)
	req.Header.Add("x-app-version", hdr.AppVersion)
	req.Header.Add("x-locale", hdr.Locale)
	req.Header.Add("x-battery-level", hdr.BatteryLevel)
	req.Header.Add("x-request-id", hdr.RequestID())
}
