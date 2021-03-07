package ColorOutput

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/gogf/gf/container/garray"
)

// Linux -------------------------

// 前景 背景 颜色
// -------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

// Windows -------------------------
// cmd下查看颜色编号： color /?
// 0 = Black       8 = Gray
// 1 = Blue        9 = Light Blue
// 2 = Green       A = Light Green
// 3 = Aqua        B = Light Aqua
// 4 = Red         C = Light Red
// 5 = Purple      D = Light Purple
// 6 = Yellow      E = Light Yellow
// 7 = White       F = Bright White

const (
	FrontBlack = iota + 30
	FrontRed
	FrontGreen
	FrontYellow
	FrontBlue
	FrontPurple
	FrontCyan
	FrontWhite
)

const (
	BackBlack = iota + 40
	BackRed
	BackGreen
	BackYellow
	BackBlue
	BackPurple
	BackCyan
	BackWhite
)

const (
	ModeDefault   = 0
	ModeHighLight = 1
	ModeLine      = 4
	ModeFlash     = 5
	ModeReWhite   = 6
	ModeHidden    = 7
)

const (
	CmdBlack = 0
	CmdRed = 4
	CmdGreen = 2
	CmdYellow = 6
	CmdBlue = 1
	CmdPurple = 5
	CmdCyan = 3
	CmdWhite = 7
)

var modeArr = []int{0, 1, 4, 5, 6, 7}

type ColorOutput struct {
	frontColor int
	backColor  int
	mode       int
}

var Colorful ColorOutput
var frontMap map[string]int
var backMap map[string]int

func init() {
	if runtime.GOOS == "windows" {
		Colorful = ColorOutput{frontColor: CmdGreen, backColor: CmdBlack, mode: ModeDefault}

		frontMap = make(map[string]int)
		frontMap["black"] = CmdBlack
		frontMap["red"] = CmdRed
		frontMap["green"] = CmdGreen
		frontMap["yellow"] = CmdYellow
		frontMap["blue"] = CmdBlue
		frontMap["purple"] = CmdPurple
		frontMap["cyan"] = CmdCyan
		frontMap["white"] = CmdWhite

		backMap = make(map[string]int)
		backMap["black"] = CmdBlack
		backMap["red"] = CmdRed
		backMap["green"] = CmdGreen
		backMap["yellow"] = CmdYellow
		backMap["blue"] = CmdBlue
		backMap["purple"] = CmdPurple
		backMap["cyan"] = CmdCyan
		backMap["white"] = CmdWhite
	} else {
		Colorful = ColorOutput{frontColor: FrontGreen, backColor: BackBlack, mode: ModeDefault}

		frontMap = make(map[string]int)
		frontMap["black"] = FrontBlack
		frontMap["red"] = FrontRed
		frontMap["green"] = FrontGreen
		frontMap["yellow"] = FrontYellow
		frontMap["blue"] = FrontBlue
		frontMap["purple"] = FrontPurple
		frontMap["cyan"] = FrontCyan
		frontMap["white"] = FrontWhite

		backMap = make(map[string]int)
		backMap["black"] = BackBlack
		backMap["red"] = BackRed
		backMap["green"] = BackGreen
		backMap["yellow"] = BackYellow
		backMap["blue"] = BackBlue
		backMap["purple"] = BackPurple
		backMap["cyan"] = BackCyan
		backMap["white"] = BackWhite
	}

}

// 其中0x1B是标记，[开始定义颜色，依次为：模式，背景色，前景色，0代表恢复默认颜色。
func (c ColorOutput) Println(str interface{}) {
	if runtime.GOOS == "windows" {
		//CmdPrint(str, c.frontColor)
		fmt.Println(str)
	} else {
		fmt.Println(fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, c.mode, c.backColor, c.frontColor, str, 0x1B))
	}
}

//func CmdPrint(s interface{}, i int) { //设置终端字体颜色
//	kernel32 := syscall.NewLazyDLL("kernel32.dll")
//	proc := kernel32.NewProc("SetConsoleTextAttribute")
//	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
//	fmt.Print(s)
//	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
//	CloseHandle := kernel32.NewProc("CloseHandle")
//	CloseHandle.Call(handle)
//}

func (c ColorOutput) WithFrontColor(color string) ColorOutput {
	color = strings.ToLower(color)
	co, ok := frontMap[color]
	if ok {
		c.frontColor = co
	}
	return c
}

func (c ColorOutput) WithBackColor(color string) ColorOutput {
	color = strings.ToLower(color)
	co, ok := backMap[color]
	if ok {
		c.backColor = co
	}

	return c
}

func (c ColorOutput) WithMode(mode int) ColorOutput {
	a := garray.NewIntArrayFrom(modeArr, true)
	bo := a.Contains(mode)
	if bo {
		c.mode = mode
	}

	return c
}
