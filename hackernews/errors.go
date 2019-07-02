package hackernews

import "fmt"

var (
	ConvertErr        = fmt.Errorf("Failed to convert item to story.")
	OutOfRangeErr     = fmt.Errorf("Value is out of range")
	InvalidTimeOutErr = fmt.Errorf("timeout must be a positive number greater than 0")
	EmptyStringErr    = fmt.Errorf("Empty string not allowed!")
	MaxStringErr      = fmt.Errorf("Max string length must be more than 0")
)

type ClientErr struct {
	msg string
}

type MinValErr struct {
	min    int
	actual int
}

type InvalidItemTypeErr struct {
	expectedtype string
	actualType   string
}

func (e *ClientErr) Error() string {
	return fmt.Sprintf("Failed to create Client. \t %s", e.msg)
}

func (e *MinValErr) Error() string {
	return fmt.Sprintf("The value was smaller than the minimum. Must be more than or equal to %d, but was %d", e.min, e.actual)
}

func (e *InvalidItemTypeErr) Error() string {
	return fmt.Sprintf("Type was not as expected. Expected %s got %s ", e.expectedtype, e.actualType)
}
