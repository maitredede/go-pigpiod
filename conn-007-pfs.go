package pigpiod

// PFS Set GPIO PWM frequency
func (c *Conn) PFS(gpio int, value uint32) error {
	cmd := Cmd{
		cmd: 7,
		p1:  uint32(gpio),
		p2:  value,
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
