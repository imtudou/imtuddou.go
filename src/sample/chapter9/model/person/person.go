package person

type person struct {
	Name string
	age  int
	sal  float64
}

// 函数：函数是 Golang 程序的基本组成部分之一，是一段独立的代码块，可以被独立地定义和调用。函数的定义以 func 关键字开始，后面跟着函数名、参数列表、返回值类型和函数体
func CreatePerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 方法：方法是与特定类型（结构体类型、函数类型、接口类型等）相关联的函数。是在类型的定义之外定义的，但与该类型紧密关联。方法的定义类似于函数，但在函数名之前会添加一个接收者（receiver），指定方法属于哪个类型。
func (p *person) SetAge(age int) {
	if age > 0 && age <= 200 {
		p.age = age
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal > 3000 && sal <= 30000 {
		p.sal = sal
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
