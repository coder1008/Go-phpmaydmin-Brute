/*
	DES:Banner头部横幅信息
	Author:冰封<574578944@qq.com>
	Date:2021-01-13
*/

package main

import (
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func init() {
	isEnabled := true
	isColorEnabled := true
	templ := `{{ .AnsiColor.BrightGreen }}{{ .Title "Brute Force" "" 4}}
	{{ .AnsiColor.BrightCyan }}phpmyadmin brute force cracking V 1.0 By 冰封
	OS: 	{{ .GOOS }}
  	CPU: 	{{ .NumCPU }}
	Now: 	{{ .Now "Monday, 2 Jan 2006" }}{{ .AnsiColor.Default }}
	{{ .AnsiColor.BrightRed }}USAGE:  -h [Help]{{ .AnsiColor.Default }}
`
	banner.InitString(colorable.NewColorableStdout(), isEnabled, isColorEnabled, templ)
}

/*
 * 自定义设置某个输出信息的颜色
 * colors= red, green,yellow, blue ;为空默认为白色
 */
func setFontColor(colors, data string) {
	switch {
	case colors == "red": //红色
		templ := `{{ .AnsiColor.BrightRed }}` + data + `{{ .AnsiColor.Default }}`
		banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	case colors == "green": //绿色
		templ := `{{ .AnsiColor.Green }}` + data + `{{ .AnsiColor.Default }}`
		banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	case colors == "yellow": //黄色
		templ := `{{ .AnsiColor.Yellow }}` + data + `{{ .AnsiColor.Default }}`
		banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	case colors == "blue": //蓝色
		templ := `{{ .AnsiColor.Blue }}` + data + `{{ .AnsiColor.Default }}`
		banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	default: //默认白色
		templ := `{{ .AnsiColor.White }}` + data + `{{ .AnsiColor.Default }}`
		banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	}
}
