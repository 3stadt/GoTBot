package fail

import "fmt"

type NotEnoughArgs struct {
	Min int
}

func (e *NotEnoughArgs) Error() string {
	return fmt.Sprintf("Not enough arguments, min. %d required!", e.Min)
}
