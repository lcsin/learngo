package json

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	Ip   string `json:"ip"`
}

func SerializeMap() {
	m := make(map[string]interface{})
	m["name"] = "linux"
	m["port"] = 8080
	m["ip"] = "192.168.5.132"

	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("serialize map:", string(json))
}

func SerializeStruct() {
	server := Server{
		Name: "windows",
		Port: 8080,
		Ip:   "10.225.5.132",
	}

	json, err := json.Marshal(server)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("serialize struct:", string(json))
}

func UnSerializeStruct() {
	jsonStr := `{"name":"windows","ip":"192.168.5.132","port":8080}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonStr), server)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("unmarshal:", server)
}

func UnSerializeMap() {
	jsonStr := `{"name":"windows","ip":"192.168.5.132","port":8080}`
	server := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &server)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("unmarshal:", server)
}
