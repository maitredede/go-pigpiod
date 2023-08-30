package pigpiod

// I2CWD i2c Write device
// https://abyz.me.uk/rpi/pigpio/pigs.html#I2CWD
func (c *Conn) I2CWD(handle uint32, bvs []byte) error {
	cmd := Cmd{
		cmd:  57,
		p1:   handle,
		p3:   uint32(len(bvs)),
		data: bvs,
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Conn) I2cWriteDevice(handle uint32, data []byte) error {
	return c.I2CWD(handle, data)
}
