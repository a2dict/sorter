# Go Sorter

### Usage
```$xslt
    data := []Any{
		Man{Name: "zhang3", Age: 24, Gender: 1},
		Man{Name: "li4", Age: 21, Gender: 0},
		Man{Name: "wang5", Age: 26, Gender: 0},
		Man{Name: "zhao6", Age: 24, Gender: 1},
	}

	res := NewSorter().ComparingBy(func(a Any) Any { //先按Gender升序
		return a.(Man).Gender
	}).ReversedComparingBy(func(a Any) Any { //再按Age倒序
		return a.(Man).Age
	}).ComparingBy(func(a Any) Any { //先按Name升序
		return a.(Man).Name
	}).Sort(data)
```