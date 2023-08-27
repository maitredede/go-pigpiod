package pigpiod

// TICK Get current tick
// https://abyz.me.uk/rpi/pigpio/pigs.html#T/TICK
func (c *Conn) TICK() (uint32, error) {
	cmd := Cmd{
		cmd: 16,
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return res.p3, err
	}
	return res.p3, nil
}
