package jin

import (
	"fmt"
)

type Error struct {
	Code     int
	Response any
}

func (e Error) Error() string {
	return fmt.Sprintf("code = %v, msg = %v", e.Code, e.Response)
}
