package errors

import (
	"fmt"
	"testing"
)

func TestErrCode_String(t *testing.T) {
	var code ErrCode = 6
	fmt.Println(code.String())
	fmt.Println(ERR_OK.String())
	fmt.Println(ERR_PARAMS.String())
}
