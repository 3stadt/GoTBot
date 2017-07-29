package fail

import (
	"fmt"
	"strings"
)

type InvalidStruct struct {
	MissingFields []string
}

func (e *InvalidStruct) Error() string {
	return fmt.Sprintf("The following struct fields have to be provided and must not be empty: %s", strings.Join(e.MissingFields, ", "))
}
