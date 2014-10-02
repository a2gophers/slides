package main

import (
	"encoding/json"
	"github.com/ttacon/pretty"
)

type Msg struct {
	MsgID int
	Name  string
	Body  string `json:"data"`
}

func main() {
	mBytes := []byte(`{
            "MsgID": 1,
            "Name":  "Msg-1",
            "data":  "such a cool message"
          }`)
	var m Msg
	_ = json.Unmarshal(mBytes, &m)
	pretty.Println(m)
}
