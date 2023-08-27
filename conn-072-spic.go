package pigpiod

// SPIC SPI close handle
func (c *Conn) SPIC(handle uint32) (uint32, error) {
	cmd := Cmd{
		cmd: 72,
		p1:  handle,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p3, err
	}
	return res.p3, nil
}
