package rgbled

import (
	"fmt"
	"image/color"

	"github.com/maitredede/go-pigpiod"
)

func NewSimpleRGBLed(c *pigpiod.Conn, gpioRed, gpioGreen, gpioBlue int) (*RGBLed, error) {

	if err := c.ModeSet(gpioRed, pigpiod.ModeOutput); err != nil {
		return nil, fmt.Errorf("failed to configure gpio %v as output: %w", gpioRed, err)
	}
	if err := c.ModeSet(gpioGreen, pigpiod.ModeOutput); err != nil {
		return nil, fmt.Errorf("failed to configure gpio %v as output: %w", gpioGreen, err)
	}
	if err := c.ModeSet(gpioBlue, pigpiod.ModeOutput); err != nil {
		return nil, fmt.Errorf("failed to configure gpio %v as output: %w", gpioBlue, err)
	}

	led := RGBLed{
		c:         c,
		gpioRed:   gpioRed,
		gpioGreen: gpioGreen,
		gpioBlue:  gpioBlue,
	}
	return &led, nil
}

type RGBLed struct {
	c         *pigpiod.Conn
	gpioRed   int
	gpioGreen int
	gpioBlue  int
}

func (l *RGBLed) Set(r byte, g byte, b byte) error {
	rP := int(r) * 100 / 255
	gP := int(g) * 100 / 255
	bP := int(b) * 100 / 255
	if err := l.c.PWM(l.gpioRed, rP); err != nil {
		return fmt.Errorf("failed to set red: %w", err)
	}
	if err := l.c.PWM(l.gpioGreen, gP); err != nil {
		return fmt.Errorf("failed to set green: %w", err)
	}
	if err := l.c.PWM(l.gpioBlue, bP); err != nil {
		return fmt.Errorf("failed to set blue: %w", err)
	}
	return nil
}

func (l *RGBLed) SetRGBA(c color.RGBA) error {
	return l.Set(c.R, c.G, c.B)
}
