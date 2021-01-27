package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const ip = "192.168.8.1"

var netmodeURL = fmt.Sprintf("http://%s/api/net/net-mode", ip)
var sesTokInfoURL = fmt.Sprintf("http://%s/api/webserver/SesTokInfo", ip)

// Session struct
type Session struct {
	SesInfo string
	TokInfo string
}

// NetMode struct
type NetMode struct {
	NetworkMode string
	NetworkBand string
	LTEBand     string
}

func getInfo() *Session {
	response, err := http.Get(sesTokInfoURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	session := Session{}
	xml.Unmarshal(responseData, &session)

	return &session
}

func getNetMode(session *Session) *NetMode {
	client := &http.Client{}
	req, err := http.NewRequest("GET", netmodeURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", session.SesInfo)
	req.Header.Set("__RequestVerificationToken", session.TokInfo)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	netmode := NetMode{}
	xml.Unmarshal(responseData, &netmode)

	return &netmode
}

func setNetMode(session *Session, netmode *NetMode) {
	data, err := xml.Marshal(netmode)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", netmodeURL, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", session.SesInfo)
	req.Header.Set("__RequestVerificationToken", session.TokInfo)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	response.Body.Close()
}

func main() {
	session := *getInfo()
	netmode := *getNetMode(&session)
	if netmode.LTEBand == "40" {
		netmode.LTEBand = "4"
	} else {
		netmode.LTEBand = "40"
	}
	setNetMode(&session, &netmode)
}
