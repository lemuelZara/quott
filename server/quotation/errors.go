package quotation

import "fmt"

// ErrCoinNotExists is used to indicate that the selected coin does not exist.
var ErrCoinNotExists error = fmt.Errorf("coin not exists")
