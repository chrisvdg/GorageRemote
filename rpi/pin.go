package rpi

import (
	"log"
)

// NewPin is a pin constructor
func NewPin(pinName string) (*Pin, error) {
	return &Pin{
		pin: pinName,
	}, nil
}

// Pin represents a pin on an rpi
type Pin struct {
	pin string
}

// Press presses the button for a short time
func (p *Pin) Press() {
	log.Printf("pin %s was pressed", p.pin)
}
