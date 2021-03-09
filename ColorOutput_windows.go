// +build windows

/*
-- @Time : 2021/3/9 8:53
-- @Author : raoxiaoya
-- @Desc :
*/
package ColorOutput

import (
	"fmt"
	"syscall"
)

var (
	kernel32 *syscall.LazyDLL
)

func init() {
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	CmdPrint = SetCmdPrint
}

// 通过调用windows操作系统API设置终端文本属性，包括：前景色，背景色，高亮。可同时设置多个属性，使用竖线 | 隔开。
// DOC: https://docs.microsoft.com/zh-cn/windows/console/setconsoletextattribute
// Usage: https://docs.microsoft.com/zh-cn/windows/console/using-the-high-level-input-and-output-functions
// 属性值: https://docs.microsoft.com/zh-cn/windows/console/console-screen-buffers#character-attributes
// SetConsoleTextAttribute函数用于设置显示后续写入文本的颜色。在退出之前，程序会还原原始控制台输入模式和颜色属性。但是微软官方
// 建议使用“虚拟终端”来实现终端控制，而且是跨平台的。
//
// 建议使用基于windows提供的“虚拟终端序列”来实现兼容多平台的终端控制，比如：https://github.com/gookit/color
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences
// https://docs.microsoft.com/zh-cn/windows/console/console-virtual-terminal-sequences#samples

func SetCmdPrint(s interface{}, i int) {
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Println(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}
