package pigpiod

func (c *Conn) Read(gpio int) (Level, error) {
	cmd := Cmd{
		cmd: 3,
		p1:  uint32(gpio),
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return LevelLow, err
	}
	return Level(res.p3), nil
}
