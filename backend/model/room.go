package model

import "errors"

type Room struct {
	Name string `json:"roomName"`
}

func (room *Room) Validate() error {
	if len(room.Name) < 6 {
		return errors.New("room name cannot be less than 6 characters")
	}
	return nil
}
