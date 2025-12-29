package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ApiResponse struct {
	Value int `json:"return_value"`
}

type params struct {
	DeviceID string `json:"deviceID"`
	Token    string `json:"token"`
}

type status struct {
	Status string `json:"status"`
}

func callParticle(deviceID string, token string, function string) *http.Response {
	particleAPI := fmt.Sprintf("https://api.particle.io/v1/devices/%s/%s", deviceID, function)

	params := url.Values{}
	params.Add("arg", "")
	params.Add("access_token", token)

	body := strings.NewReader(params.Encode())

	log.Println(particleAPI)
	log.Println(body)
	req, err := http.NewRequest("POST", particleAPI, body)
	if err != nil {
		fmt.Printf("Error sending request: %s", err)
		var result status
		result.Status = "unknown"
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error sending request: %s", err)
	}

	return resp
}

func getParams(r *http.Request) (string, string) {
	var data params

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		return "", ""
	}

	log.Println("JSON")
	log.Println(data)

	return data.DeviceID, data.Token
}

func particleDoorStatus(deviceID string, token string) status {
	resp := callParticle(deviceID, token, "getState")

	log.Println(resp)
	defer resp.Body.Close()

	var r ApiResponse

	err := json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Println(err.Error())
	}

	var result status

	log.Println(r.Value)

	if r.Value == 1 {
		result.Status = "open"
	} else {
		result.Status = "closed"
	}

	return result
}

func (s *server) DoorStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("DoorStatus")

	deviceID, token := getParams(r)

	log.Println("deviceID: ")
	log.Println(deviceID)
	log.Println("token: ")
	log.Println(token)
	status := particleDoorStatus(deviceID, token)

	log.Println("status")
	log.Println(status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(status.Status))
}

func (s *server) OpenDoor(w http.ResponseWriter, r *http.Request) {
	log.Println("OpenDoor")

	deviceID, token := getParams(r)

	callParticle(deviceID, token, "open")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (s *server) CloseDoor(w http.ResponseWriter, r *http.Request) {
	log.Println("CloseDoor")

	deviceID, token := getParams(r)

	callParticle(deviceID, token, "close")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}

func (s *server) OpenClose(w http.ResponseWriter, r *http.Request) {
	log.Println("OpenClose")

	deviceID, token := getParams(r)

	callParticle(deviceID, token, "openClose")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
