package pigpiod

func (c *Conn) ModeSet(gpio int, mode GpioMode) error {
	cmd := Cmd{
		cmd: 0,
		p1:  uint32(gpio),
		p2:  uint32(mode),
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
