package pigpiod

func (c *Conn) I2CC(handle uint32) (uint32, error) {
	cmd := Cmd{
		cmd: 55,
		p1:  handle,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p3, err
	}
	return res.p3, nil
}
