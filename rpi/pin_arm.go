package rpi

import (
	"log"
	"time"

	"github.com/mrmorphic/hwio"
)

// NewPin is a pin constructor
func NewPin(pinName string) (*Pin, error) {
	pin, err := hwio.GetPin(pinName)
	if err != nil {
		return nil, err
	}
	err = hwio.PinMode(pin, hwio.OUTPUT)
	if err != nil {
		return nil, err
	}

	return &Pin{
		name: pinName,
		pin:  pin,
	}, nil
}

// Pin represents a pin on an rpi
type Pin struct {
	pin     hwio.Pin
	name    string
	pressed bool
}

// Press presses the button for a short time
func (p *Pin) Press() {
	go func() {
		if !p.pressed {
			log.Printf("pin %s is now being pressed", p.name)
			p.pressed = true
			defer func() {
				p.pressed = false
			}()
			err := hwio.DigitalWrite(p.pin, hwio.HIGH)
			if err != nil {
				log.Printf("Error while pressing button %s: %s", p.name, err)
				return
			}
			time.Sleep(500 * time.Millisecond)
			err = hwio.DigitalWrite(p.pin, hwio.LOW)
			if err != nil {
				log.Printf("Error while pressing button %s: %s", p.name, err)
				return
			}
			time.Sleep(500 * time.Millisecond)
		} else {
			log.Printf("pin %s was already pressed", p.name)
		}
	}()
}
