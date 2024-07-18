package main

import (
	"fmt"
	"github.com/Kay7271/demo/mypackage"
	"math"
	"slices"
)

var b = 3

const (
	pi = 3.1415926
	e  = 2.71828
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	_ = iota * iota
	a
	c
	d
)

func main() {
	mapp()
}

func mapp() {
	//m := make(map[string]int, 1)
	//m["a"] = 1
	//m["b"] = 2
	//len := len(m)
	//fmt.Println(len)

	//var user map[string]string
	//user = map[string]string{
	//	"name": "kay",
	//	"age":  "12",
	//}
	////user["name"] = "kay"
	////user["age"] = "12"
	//fmt.Println(user)
	//s, ok := user["gender"]
	//fmt.Println(s, ok)
	//
	//for k, v := range user {
	//	fmt.Printf("%s: %s\n", k, v)
	//}
	//
	//delete(user, "a")
	//fmt.Println(user)

	var m = make([]map[string]string, 5)
	m[0] = map[string]string{
		"张三": "12",
	}
	m[1] = map[string]string{
		"李四": "13",
	}
	fmt.Println(m)

}

func slice() {
	//var a = [4]int{1, 2, 3, 4}
	//s := a[1:2]
	////a[0] = "hello"
	//fmt.Printf("Type: %T, Value: %v\n", a, a)
	//s = append(s, 5)
	//fmt.Printf("Type: %T, Value: %v\n", s, s)

	//a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
	//s := a[3:5] // s := a[low:high]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	//fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3:3]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	//a := make([]int, 3, 5)
	//a[1] = 10
	//a = append(a, 1, 2, 3)
	//fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))
	//for i, aa := range a {
	//	fmt.Println(i, aa)
	//}

	//var arr = [5]int{1, 2, 3, 4, 5}
	//arr1 := arr
	//arr1[0] = 100
	//fmt.Printf("arr:%v arr1:%v\n", arr, arr1)

	// 切片的复制
	//arr := []int{1, 2, 3, 4, 5}
	//arr1 := make([]int, 5, 5)
	//copy(arr1, arr)
	//arr1[0] = 100
	//fmt.Printf("arr:%v arr1:%v\n", arr, arr1)

	// 切片的删除
	//arr := []int{1, 2, 3, 4, 5}
	//arr = append(arr[:2], arr[3:]...)
	//fmt.Printf("arr:%v\n", arr)

	//var a = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	a = append(a, fmt.Sprintf("%v", i))
	//}
	//fmt.Printf("a:%v len(a):%v cap(a):%v\n", a, len(a), cap(a))

	a := [...]int{3, 7, 8, 9, 1}
	slices.Sort(a[:])
	fmt.Printf("Type: %T, Value: %v\n", a, a)

}

func array() {
	//var arr [5]int
	//arr[0] = 1
	//fmt.Printf("Type: %T, Value: %v\n", arr, arr)
	//fmt.Println(arr[1])

	//var arr1 = [...]int{1: 1, 3: 6}
	//fmt.Printf("Type: %T, Value: %v\n", arr1, arr1)

	//for _, i := range arr1 {
	//	fmt.Println(i)
	//}

	a := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	add1(a)

	fmt.Printf("Type: %T, Value: %v\n", a, a)

}

func add1(arr [2][3]int) {
	arr[0][0] = 100
}

func ope(a, b int) (sum, sub int) {
	return a + b, a - b
}

func variable() {
	name := "Kay"

	var age int
	age = 12

	var gender string = "female"

	var id = 1

	mypackage.New()
	fmt.Printf("ID: %d,", id)
	fmt.Printf("Name: %s,", name)
	fmt.Printf("Age: %d,", age)
	fmt.Printf("Gender: %s", gender)

	a := 1
	_, sub := ope(a, b)
	fmt.Printf("Sub: %d", sub)

	fmt.Println(d)
}

func datatype() {
	//	v := 0x2p-3
	//	fmt.Printf("%f\n", v)
	//	v2 := 12_345
	//	fmt.Printf("%d\n", v2)
	//	p := math.Pi
	//	fmt.Printf("%.100f\n", p)
	//	var c1 complex64
	//	fmt.Printf("%v\n", c1)
	//	var flag bool
	//	fmt.Printf("%v\n", flag)
	//	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	//
	//	var s1 = `
	//Hello
	//		World
	//				!
	//`
	//	fmt.Println(s1)
	//
	//	l := len(s1)
	//	fmt.Println(l)
	//
	//	str1 := "hello" + "go"
	//	fmt.Println(str1)
	//	s2 := fmt.Sprint("hello", "go")
	//	fmt.Println(s2)
	//
	//	split := strings.Split("a,b,c", ",")
	//	fmt.Printf("%v\n", split)
	//
	//	contains := strings.Contains("seafood", "foo")
	//	fmt.Println(contains)
	//
	//	join := strings.Join([]string{"a", "b", "c"}, ",")
	//	fmt.Println(join)
	//
	//	hasPrefix := strings.HasPrefix("goodbye", "good")
	//	fmt.Printf("%v\n", hasPrefix)
	//
	//	index := strings.LastIndexAny("goodbye", "o")
	//	fmt.Printf("index: %d\n", index)
	//var a uint8 = 'A'
	//fmt.Println(a)
	//var b = '中'
	//fmt.Println(b)

	//s1 := "big"
	//bytes := []uint8(s1) // uint8 == byte
	//bytes[0] = 'p'
	//fmt.Printf("%s\n", string(bytes))

	//s2 := "白萝卜"
	//runes := []int32(s2) // int32 == rune
	//runes[0] = '红'
	//fmt.Printf("%s\n", string(runes))

	a1, a2 := 3, 4
	sqrt := int(math.Sqrt(float64(a1*a1 + a2*a2)))
	// 输出sqrt的类型和值
	fmt.Printf("Type: %T, Value: %v\n", sqrt, sqrt)
}

func operator() {
	//a1 := 0x0101
	//a2 := 0x1111
	//i := a1 & a2
	//fmt.Printf("Type: %T, Value: %b\n", i, i)
	fmt.Printf("%d", 1^2^3^2^1)
}

func controlFlow() {
	//var max int = getMax(1, 3)
	//fmt.Printf("Max: %d", max)
	//for _, i2 := range "abcd" {
	//	fmt.Printf("%c\n", i2)
	//}
	//
	//n := 7
	//switch {
	//case n > 0:
	//	fmt.Println("n > 0")
	//case n > 3:
	//	fmt.Println("n > 3")
	//case n > 5:
	//	fmt.Println("n > 5")
	//default:
	//	fmt.Println("default")
	//}

	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
			//if i*j > 18 {
			//	fmt.Println()
			//	goto breakHere
			//}
		}
		fmt.Println()
	}
	//breakHere:
	//	fmt.Println("breakHere")
}

func getMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
