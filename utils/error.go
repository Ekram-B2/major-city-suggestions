package utils

// import "fmt"

// // Error struct created to simplify managing error cases
// type Error struct {
// 	data interface{}
// 	msg  string
// }

// func makeError(data interface{}, msgAndArgs ...interface{}) error {
// 	msg := ""
// 	switch len(msgAndArgs) {
// 	case 0:
// 		// Ignore
// 	case 1:
// 		// If its an error code
// 		msg = fmt.Sprint(msgAndArgs[0])
// 	default:
// 		// if its a error message
// 		if str, ok := msgAndArgs[0].(string); ok {
// 			msg = fmt.Sprintf(str, msgAndArgs[1:]...)
// 		}
// 		msg = fmt.Sprint(msgAndArgs...)
// 	}
// 	return &Error{data: data, msg: msg}
// }
