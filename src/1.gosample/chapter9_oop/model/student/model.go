package student

type Student struct {
	Name  string
	Age   int
	Score float64
}

type studentsss struct {
	Name   string
	Age    int
	Score  float64
	height float64
}

// 因为student结构体首字母是小写，因此是只能在model使用//我们通过工厂模式来解决
func CreateStudent(name string, age int, score float64, h float64) *studentsss {
	return &studentsss{
		Name:   name,
		Age:    age,
		Score:  score,
		height: h,
	}
}

//如果score字段首字母小写，则，在其它包不可以直接方法，我们可以提供一个方法

func (s *studentsss) GetHeight() float64 {
	return s.height
}
