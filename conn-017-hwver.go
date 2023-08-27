package pigpiod

// HWVER Get hardware version
// https://abyz.me.uk/rpi/pigpio/pigs.html#HWVER
func (c *Conn) HWVER() (uint32, error) {
	cmd := Cmd{
		cmd:  17,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p3, err
	}
	return res.p3, nil
}
