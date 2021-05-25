package main //一个包定义一组编译过的代码，包的名字类似于命名空间，可以用来简介访问包内声明的标识符

import ( //声明导入项
	"log" //标准库包里导入的只需要写包名，回去GOROOT和GOPATH环境变量引用的位置去寻找
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers" //注意前面有下划线，这是对包进行初始化操作，并且调用包内对应代码的定义的init函数
	"github.com/goinaction/code/chapter2/sample/search" //导入包，也需要执行包里面的init函数
)

// init is called prior to main.
func init() { //init函数会在执行main函数之前调用
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {// 如果main函数不在main包里，构建工具不会生成可执行文件
	// Perform the search for the specified term.
	search.Run("president")//调用search包里面是Run函数
}
