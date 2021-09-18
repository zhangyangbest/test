package main

import (
	"fmt"
	"gotest/errors"
)

func main() {
	var aa errors.ErrCode
	aa = 6
	fmt.Println(aa)
	fmt.Println(errors.ERR_OK.String())
	fmt.Println(aa.String())
}
