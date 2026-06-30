package struct_test
//结构体学习
import (
	"fmt"
	"testing"
)

// 如何初始化结构体
// e:=employee{"0","bob","20"}
// e1:=employee{name:"mike",age:"30"}//省略的自动赋类型零值
// e2:=new(employee)这里返回引用，等价于e：=&employee{}
// e3:=struck{age int}{}声明+赋值
type employee struct {
	Name string
	Age  uint8
}

func (e employee) employeeprint() {
	fmt.Printf("%p\n",&e)
}
func (e *employee) employeeprint_pointer(){
	fmt.Printf("%p\n",e)
}
func TestStruct(t *testing.T) {
	e:=employee{"yanghcenhzhi",18}
	e.employeeprint()
	e.employeeprint_pointer()
}
