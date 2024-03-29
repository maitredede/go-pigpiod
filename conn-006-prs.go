package pigpiod

// PRS Set GPIO PWM range
func (c *Conn) PRS(gpio int, value uint32) error {
	cmd := Cmd{
		cmd: 6,
		p1:  uint32(gpio),
		p2:  value,
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
