/*
-- @Time : 2021/3/8 15:31
-- @Author : raoxiaoya
-- @Desc :
*/
package ColorOutput

import (
	"testing"
)

func TestColorOutput(t *testing.T) {
	Colorful.WithFrontColor("green").WithBackColor("red").Println("ColorOutput test...")
	Colorful.WithFrontColor("blue").WithBackColor("red").Println("ColorOutput test...")
}
