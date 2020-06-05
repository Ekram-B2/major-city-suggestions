package relevantreader

import (
	"fmt"
	"testing"
)

func Test_isWellWritten(t *testing.T) {
	var res interface{}
	ret := isWellFormatted(res)
	t.Fatalf(fmt.Sprintf("%v", ret))
	if ret == true {
		t.Fatalf("Failure")
	}
}
