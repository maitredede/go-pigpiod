package pigpiod

// SPIW SPI write bytes to handle
func (c *Conn) SPIW(handle uint32, data []byte) error {
	cmd := Cmd{
		cmd:  74,
		p1:   handle,
		p3:   uint32(len(data)),
		data: data,
	}
	_, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
