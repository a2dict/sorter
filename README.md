# Go Sorter

### Usage
```go
// see sorter_test.go
func TestSorter(t *testing.T) {
    data := []Man{
        Man{Name: "zhang3", Age: 24, Gender: 1},
        Man{Name: "li4", Age: 21, Gender: 0},
        Man{Name: "wang5", Age: 26, Gender: 0},
        Man{Name: "zhao6", Age: 24, Gender: 1},
    }

    err := NewSorter().ComparingBy(func(a interface{}) interface{} { // order by Gender asc
        return a.(Man).Gender
    }).ReversedComparingBy(func(a interface{}) interface{} { //then order by Age desc
        return a.(Man).Age
    }).ComparingBy(func(a interface{}) interface{} { //then order by Name asc
        return a.(Man).Name
    }).Sort(&data)
    if err != nil {
        t.Log(data)
    }
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

```