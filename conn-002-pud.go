package pigpiod

// PUD Set GPIO pull up/down
func (c *Conn) PUD(gpio int, pull Level) error {
	cmd := Cmd{
		cmd: 2,
		p1:  uint32(gpio),
		p2:  uint32(pull),
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
