package sorter

import (
	"fmt"
	"testing"
)

type Man struct {
	Name   string
	Age    int
	Gender int8 //0 female, 1 male
}

func TestSorter(t *testing.T) {
	data := []Man{
		Man{Name: "zhang3", Age: 24, Gender: 1},
		Man{Name: "li4", Age: 21, Gender: 0},
		Man{Name: "wang5", Age: 26, Gender: 0},
		Man{Name: "zhao6", Age: 24, Gender: 1},
	}

	NewSorter().MoveForward(func(a interface{}) bool { //move female forward
		return a.(Man).Gender == 0
	}).ReversedComparingBy(func(a interface{}) interface{} { //order by Age desc
		return a.(Man).Age
	}).ComparingBy(func(a interface{}) interface{} { //order by Name asc
		return a.(Man).Name
	}).Sort(&data)

	fmt.Printf("%+v", data)
	// output:
	//[{Name:wang5 Age:26 Gender:0} {Name:li4 Age:21 Gender:0} {Name:zhang3 Age:24 Gender:1} {Name:zhao6 Age:24 Gender:1}]--- PASS: TestSorter (0.00s)

}

func TestSorter_MoveForward(t *testing.T) {
	data := []Man{
		Man{Name: "zhang3", Age: 24, Gender: 1},
		Man{Name: "li4", Age: 21, Gender: 0},
		Man{Name: "wang5", Age: 26, Gender: 0},
		Man{Name: "zhao6", Age: 24, Gender: 1},
	}

	NewSorter().MoveForward(func(a interface{}) bool {
		return a.(Man).Gender == 0
	}).Sort(&data)
	fmt.Printf("%+v", data)

	// output:
	//[{Name:li4 Age:21 Gender:0} {Name:wang5 Age:26 Gender:0} {Name:zhang3 Age:24 Gender:1} {Name:zhao6 Age:24 Gender:1}]--- PASS: TestSorter_MoveForward

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
	fmt.Printf("%+v", data)

	// output:
	//[{Name:zhang3 Age:24 Gender:1} {Name:zhao6 Age:24 Gender:1} {Name:li4 Age:21 Gender:0} {Name:wang5 Age:26 Gender:0}]--- PASS: TestSorter_MoveBackward
}
