package fail

import "fmt"

type TooManyArgs struct {
	Max int
}

func (e *TooManyArgs) Error() string {
	return fmt.Sprintf("Too many arguments, max. %d allowed!", e.Max)
}
