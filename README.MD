## 1. 命令行输入以下命令来安装第三方包:
- go get github.com/dimiro1/banner
- go get github.com/mattn/go-colorable
--------------------------------------------  


## 1. GoLang 在Windows下编译生成带图标的exe 
### (1) Windows系统上安装MinGW
- (2)MinGW下载地址：https://sourceforge.net/projects/mingw-w64/files/mingw-w64/
- (3)下载64位x86_64-posix-seh压缩包,直接解压就行(例如我这里解压到 D:\mingw64 )
- (4)如何让go编译器调用gcc呢？需要设置环境变量path，添加 D:\mingw64\bin 到path环境变量值下,最后在命令行下输入 gcc -v 检查是否安装成功

### (2) 添加一个ICO图标
* (1)准备1张logo.ico图片放到项目目录下
* (2)在项目目录下创建一个空白文本文件,文件命名为 main.rc 并写入内容: IDI_ICON1 ICON "logo.ico"
* (3)在项目目录下执行CMD命令 windres -o main.syso main.rc 
* (4)在项目目录下直接go build, 生成的exe就有了一个好看的图标了.

### (3) 编译出带图标的exe 
* 在项目目录下执行 go build -o phpmyadmin暴力破解.exe
--------------------------------------------  


## 3. GoLang 在Windows下编译Mac, Linux平台的64位可执行程序
### (1) 在Windows下编译Mac平台的64位可执行程序

- 在项目目录下执行以下命令编译生成
- SET CGO_ENABLED=0
- SET GOOS=darwin
- SET GOARCH=amd64
- go build -o bf_phpmyadmin
  
### (2) 在Windows下编译Linux平台的64位可执行程序
- 在项目目录下执行以下命令编译生成
- SET CGO_ENABLED=0
- SET GOOS=linux 
- SET GOARCH=amd64 
- go build -o bf_phpmyadmin

#### GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows
#### GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm

### (3) GoLang 在Windows下编译Mac, Linux平台的64位可执行程序报错处理:
打包出现如下错误 cannot find package "golang.org/x/sys/unix" in 
- 解决办法如下:
- git clone https://github.com/golang/sys.git
- 下载到gopath目录下的golang.org\x\sys 下即可
--------------------------------------------  