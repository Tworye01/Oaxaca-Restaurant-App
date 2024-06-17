package server

import (
	"encoding/json"
	"log"
)

// functions to put in the worker maybe

type Packet struct {
	Body string `json:"body"`
}

func CreatePacket(command string) []byte {
	ret, err := json.Marshal(Packet{
		Body: command,
	})

	if err != nil {
		log.Println(err)
	}
	return ret
}
