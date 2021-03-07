/*
-- @Time : 2021/3/7 17:17
-- @Author : raoxiaoya
-- @Desc :
*/
package ColorOutput

import (
	"fmt"
	"runtime"
	"testing"
)

func TestPrint(t *testing.T) {
	Colorful.WithFrontColor("green").Println("hhhh")
}

func TestOs(t *testing.T) {
	fmt.Println(runtime.GOOS)
}