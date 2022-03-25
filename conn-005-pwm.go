package pigpiod

import (
	"fmt"
)

func (c *Conn) PWMRaw(gpio int, duttyCycle uint32) error {
	cmd := Cmd{
		cmd: 5,
		p1:  uint32(gpio),
		p2:  duttyCycle,
	}
	_ /*res*/, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) PWM(gpio int, duttyCycle int) error {
	if duttyCycle < 0 || duttyCycle > 100 {
		return fmt.Errorf("duttyCycle=%v out of range [0-100]", duttyCycle)
	}

	rMax, found := c.duttyCycleRanges[gpio]
	if !found {
		res, err := c.PRRG(gpio)
		if err != nil {
			return err
		}
		rMax = res
	}

	r := uint32(float32(rMax) * float32(duttyCycle) / float32(100))
	return c.PWMRaw(gpio, r)
}
