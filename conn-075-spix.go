package pigpiod

// SPIX SPI transfer bytes to handle
func (c *Conn) SPIX(handle uint32, data []byte) ([]byte, error) {
	cmd := Cmd{
		cmd:  75,
		p1:   handle,
		p3:   uint32(len(data)),
		data: data,
	}
	res, err := cmd.ExecuteResData(c.tcp)
	if err != nil {
		return nil, err
	}
	return res.data[:res.p3], nil
}
