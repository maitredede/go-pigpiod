package pigpiod

func (c *Conn) I2CO(i2cBus byte, i2cAddress byte, i2cFlags uint32) (uint32, error) {
	data := makeData(i2cFlags)
	cmd := Cmd{
		cmd:  54,
		p1:   uint32(i2cBus),
		p2:   uint32(i2cAddress),
		p3:   uint32(len(data)),
		data: data,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p3, err
	}
	return res.p3, nil
}
