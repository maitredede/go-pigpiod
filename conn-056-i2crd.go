package pigpiod

// I2CRD i2c Read device
// https://abyz.me.uk/rpi/pigpio/pigs.html#I2CRD
func (c *Conn) I2CRD(handle uint32, num uint32) ([]byte, error) {
	cmd := Cmd{
		cmd: 56,
		p1:  handle,
	}
	res, err := cmd.ExecuteResData(c.tcp)
	if err != nil {
		return nil, err
	}
	return res.data[:res.p3], nil
}

func (c *Conn) I2cReadDevice(handle uint32, length int) ([]byte, error) {
	return c.I2CRD(handle, uint32(length))
}
