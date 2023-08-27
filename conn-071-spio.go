package pigpiod

// SPIO https://abyz.me.uk/rpi/pigpio/pigs.html#SPIO
func (c *Conn) SPIO(spiChannel byte, baud uint32, spiFlags uint32) (uint32, error) {
	data := makeData(spiFlags)
	cmd := Cmd{
		cmd:  71,
		p1:   uint32(spiChannel),
		p2:   baud,
		p3:   uint32(len(data)),
		data: data,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p3, err
	}
	return res.p3, nil
}
