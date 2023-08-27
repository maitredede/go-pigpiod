package pigpiod

// GpioMode mode (RW540123)
type GpioMode uint32

const (
	ModeInput  GpioMode = 0
	ModeOutput GpioMode = 1
	ModeAlt0   GpioMode = 4
	ModeAlt1   GpioMode = 5
	ModeAlt2   GpioMode = 6
	ModeAlt3   GpioMode = 7
	ModeAlt4   GpioMode = 3
	ModeAlt5   GpioMode = 2
)
