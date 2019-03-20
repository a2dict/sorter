package sorter

import "testing"

type Man struct {
	Name   string
	Age    int
	Gender int8
}

func TestSorter(t *testing.T) {
	data := []Man{
		Man{Name: "zhang3", Age: 24, Gender: 1},
		Man{Name: "li4", Age: 21, Gender: 0},
		Man{Name: "wang5", Age: 26, Gender: 0},
		Man{Name: "zhao6", Age: 24, Gender: 1},
	}

	NewSorter().ComparingBy(func(a interface{}) interface{} { //先按Gender升序
		return a.(Man).Gender
	}).ReversedComparingBy(func(a interface{}) interface{} { //再按Age倒序
		return a.(Man).Age
	}).ComparingBy(func(a interface{}) interface{} { //先按Name升序
		return a.(Man).Name
	}).Sort(&data)
	t.Log(data)
}

func TestSorter_MoveBackward(t *testing.T) {
	data := []Man{
		Man{Name: "zhang3", Age: 24, Gender: 1},
		Man{Name: "li4", Age: 21, Gender: 0},
		Man{Name: "wang5", Age: 26, Gender: 0},
		Man{Name: "zhao6", Age: 24, Gender: 1},
	}

	NewSorter().MoveBackward(func(a interface{}) bool {
		return a.(Man).Gender == 0
	}).Sort(&data)
	t.Log(data)

}
