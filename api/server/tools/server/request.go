package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Result struct {
	Code int
	Message string
	Data map[string]interface{}
}

type MonitorData struct {
	Id string
	Url string
	ExpireTime int64
}

type MonitorMessage struct {
	Msg string
	Code string
	Data MonitorData
}

type MonitorAccessMessage struct {
	Msg string
	Code string
	Data MonitorAccessToken
}

type MonitorAccessToken struct {
	AccessToken string
	ExpireTime int64
}

func GetResponseBodyMonitor(response *http.Response) MonitorMessage {
	var res MonitorMessage
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &res)
	return res
}

type NewCollarRealtimeData struct {
	Area 	string 		`json:"area"`
	Iccid	string  	`json:"iccid"`      
	Police 	int			`json:"police"`     
	AllStep	int			`json:"all_step"`
	LastTime	string	`json:"last_time"`
    Temperature		float32		`json:"temperature"`
    Station 	string	`json:"station"`
    IsOnline	int 		`json:"isOnline"`
    SignalStrength	int 	`json:"signal_strength"`
    Type 	string		`json:"type"`
    Voltage 	float32		`json:"voltage"`
}

type NewCollarRealtimeMessage struct {
	Code int	`json:"code"`
	Msg string	`json:"msg"`
	Data []NewCollarRealtimeData	`json:"data"`
}

func GetResponseBodyNewCollarRealtime(response *http.Response) NewCollarRealtimeMessage {
	var res NewCollarRealtimeMessage
	errj := json.NewDecoder(response.Body).Decode(&res)
	if errj != nil {
		panic(errj.Error())
	}
	return res
}

func GetResponseAccessMonitor(response *http.Response) MonitorAccessToken {
	var res MonitorAccessMessage
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &res)
	return res.Data
}

type NewCollarAccessMessage struct {
	Msg string
	Code string
	Data NewCollarAccessToken
}

type NewCollarAccessToken struct {
	Token string
}

func GetResponseNewCollarAccessToken(response *http.Response) string {
	var res NewCollarAccessMessage
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &res)
	return res.Data.Token
}

func GetResponseBody(response *http.Response) map[string]interface{} {
	var res Result
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &res)
	return res.Data
}