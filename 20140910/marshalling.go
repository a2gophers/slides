package main

import (
	"encoding/json"
	"fmt"
)

type Msg struct {
	MsgID int
	Name  string
	Body  string `json:"data"`
}

func main() {
	m := Msg{
		MsgID: 1,
		Name:  "Msg-1",
		Body:  "such a cool message",
	}

	msgBytes, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(msgBytes))
}
