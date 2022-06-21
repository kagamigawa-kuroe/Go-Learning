//这个demo是关于使用go自带的mod包管理工具的介绍

//在正常使用go的情况下
//我们不仅要吧go的目录添加到环境变量中
//而且对应的项目也必须要在go源文件的目录下
//我们安装的第三方包也需要放在这个目录下专门的文件中
//这其实非常不方便

//go1.11以后 新增了一个go的包管理工具
//go mod
//可以帮助我们快速实现go项目的包管理和引入

//用法为 新建一个项目的根目录
//在目录下使用 go mod init 项目名
//就会生成一个go.mod文件
//然后就可以直接开始写你的项目
//当有第三方库无法在本地找到时
//就可以使用命令 go mod tidy
//该工具就会自动帮你下载所有的缺失依赖
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
