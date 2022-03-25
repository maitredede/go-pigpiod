package pigpiod

func (c *Conn) PRS(gpio int, value uint32) (uint32, error) {
	cmd := Cmd{
		cmd: 6,
		p1:  uint32(gpio),
		p2:  value,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p2, err
	}
	return res.p2, nil
}
