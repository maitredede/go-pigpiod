package pigpiod

func (c *Conn) ModeGet(gpio int) (GpioMode, error) {
	cmd := Cmd{
		cmd: 1,
		p1:  uint32(gpio),
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return ModeInput, err
	}
	return GpioMode(res.p3), nil
}
