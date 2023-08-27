package pigpiod

// ModeSet Set GPIO mode
func (c *Conn) ModeSet(gpio int, mode GpioMode) error {
	cmd := Cmd{
		cmd: 0,
		p1:  uint32(gpio),
		p2:  uint32(mode),
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}

// ModeGet Get GPIO mode
func (c *Conn) ModeGet(gpio int) (GpioMode, error) {
	cmd := Cmd{
		cmd: 1,
		p1:  uint32(gpio),
	}
	res, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return ModeInput, err
	}
	return GpioMode(res.p3), nil
}

// PUD Set GPIO pull up/down
func (c *Conn) PUD(gpio int, pull Level) error {
	cmd := Cmd{
		cmd: 2,
		p1:  uint32(gpio),
		p2:  uint32(pull),
	}
	/*res*/ _, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}

// Read Read GPIO level
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

// Write Write GPIO level
func (c *Conn) Write(gpio int, level Level) error {
	cmd := Cmd{
		cmd: 4,
		p1:  uint32(gpio),
		p2:  uint32(level),
	}
	_ /*res*/, err := cmd.ExecuteRes(c.tcp)
	if err != nil {
		return err
	}
	return nil
}
