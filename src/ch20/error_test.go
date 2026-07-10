package error_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestError(t *testing.T) {
	_, err := os.Open("no/such/file")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
			return
		}
		errors.Is(err, fs.ErrNotExist)
		if os.IsPermission(err) {
			fmt.Println("权限不足")
			return
		}
		fmt.Println("其它错误:", err)
		//更现代的写法，可以沿着错误链寻找
		// if errors.Is(err, fs.ErrNotExist) {
		// 	fmt.Println("文件不存在")
		// 	return
		// }

		//errors.As 高级类型断言，可取出具体错误类型
		//pe, ok := err.(*os.PathError)只能判断当前的动态类型
		//现代写法
		/*var pe *os.PathError
			if errors.As(err, &pe) {
		    fmt.Println(pe.Op)
		    fmt.Println(pe.Path)
		    fmt.Println(pe.Err)
		}*/


		//errors.As配合%w可以沿着错误链寻找,不只看最外层错误，%v会丢失层级
		//err = fmt.Errorf("读取配置失败: %w", err)
		//_, ok := err.(*os.PathError)就会失败

	}

}
