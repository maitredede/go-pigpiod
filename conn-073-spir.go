package pigpiod

// SPIR SPI read bytes from handle
func (c *Conn) SPIR(handle uint32, num uint32) ([]byte, error) {
	cmd := Cmd{
		cmd: 73,
		p1:  handle,
		p2:  num,
	}
	res, err := cmd.ExecuteResData(c.tcp)
	if err != nil {
		return nil, err
	}
	return res.data[:res.p3], nil
}
