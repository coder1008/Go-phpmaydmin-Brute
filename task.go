/*
	DES:文件读写操作
	Author:冰封<574578944@qq.com>
	Date:2021-01-13
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
 * 读文件数据
 */
func ReadFiles(files string) []string {
	file, err := os.Open(files) //打开文件
	if err != nil {
		log.Fatalln("Messages: ", err)
	}

	defer file.Close()              //及时关闭 file 句柄，否则会有内存泄漏
	reader := bufio.NewReader(file) //创建一个 *Reader ， 是带缓冲的
	var fileList []string           //创建一个未设定长度的数组来存储读取到的多条文件内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}

		//fmt.Print(str)
		str = strings.Replace(str, "\n", "", -1) // 去除换行符
		str = strings.Replace(str, "\r", "", -1) // 去除回车符
		fileList = append(fileList, str)
	}
	return fileList
}

/*
 * 写文件
 */
func WriteFiles(files, datas string) {
	file, err := os.OpenFile(files, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Messages: ", err)
		return
	}

	defer file.Close() //及时关闭

	str := datas                    //写入内容
	writer := bufio.NewWriter(file) //写入时，使用带缓存的 *Writer
	_, errs := writer.WriteString(str + "\n")
	if errs != nil {
		fmt.Println("Messages: ", errs)
		return
	}
	writer.Flush() //调用flush方法,将缓存的数据真正写入到文件中。
}
