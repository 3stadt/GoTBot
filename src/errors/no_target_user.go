package fail

import "fmt"

type NoTargetUser struct {
	Name string
}

func (e *NoTargetUser) Error() string {
	return fmt.Sprintf("No target user named '%s' found in database!", e.Name)
}
