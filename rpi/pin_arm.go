// build
package rpi

import (
	"time"

	"github.com/mrmorphic/hwio"
)

// NewPin is a pin constructor
func NewPin(pinName string) (*Pin, error) {
	pin, err := hwio.GetPin(pinName)
	if err != nil {
		return nil, err
	}

	return &Pin{
		pin: pin,
	}, nil
}

// Pin represents a pin on an rpi
type Pin struct {
	pin     hwio.Pin
	pressed bool
}

// Press presses the button for a short time
func (p *Pin) Press() {
	if !p.pressed {
		p.pressed = true
		defer func() {
			p.pressed = false
		}()

		hwio.DigitalWrite(p.pin, hwio.HIGH)
		time.Sleep(500 * time.Millisecond)
		hwio.DigitalWrite(p.pin, hwio.LOW)
		time.Sleep(500 * time.Millisecond)
	}
}
