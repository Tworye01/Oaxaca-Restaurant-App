package server

import (
	"encoding/json"
	"testing"
)

func TestCreatePacket(t *testing.T) {
	// Tests with valid command
	command := "command"
	expectedBody := "command"
	expectedJSON, _ := json.Marshal(Packet{Body: expectedBody})

	packet := CreatePacket(command)

	if string(packet) != string(expectedJSON) {
		t.Errorf("Expected(%s) but got %s", expectedJSON, packet)
	}

	// Test with empty command
	emptyCommand := ""
	expectedEmptyBody := ""
	expectedEmptyJSON, _ := json.Marshal(Packet{Body: expectedEmptyBody})

	emptyPacket := CreatePacket(emptyCommand)

	if string(emptyPacket) != string(expectedEmptyJSON) {
		t.Errorf("Expected(%s) but got %s", expectedEmptyJSON, emptyPacket)
	}
}
