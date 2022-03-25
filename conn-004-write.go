package pigpiod

func (c *Conn) Write(gpio int, level Level) error {
	cmd := Cmd{
		cmd: 4,
		p1:  uint32(gpio),
		p2:  uint32(level),
	}
	_ /*res*/, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
