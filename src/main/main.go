package main

import "fmt"

/**
if
for
switch
goto

*/

func main() {

	//1.
	//age := 30
	//if age > 20 {
	//	fmt.Println("flag=true")
	//}

	//2.
	//if age := 20; age > 18 {
	//	fmt.Println("person")
	//}

	//score := 90
	//if score > 90 {
	//	fmt.Println("A")
	//} else if score > 75 {
	//	fmt.Println("B")
	//} else {
	//	fmt.Println("C")
	//}

	//for expire
	//1 - 10
	//i := 1
	//for ; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println(i)

	//i := 1
	//for i < 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//for i := 0; i <= 50; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i)
	//	}
	//}

	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//
	//count := 0
	//sum1 := 0
	//for i := 1; i <= 100; i++ {
	//	if i%9 == 0 {
	//		count++
	//		sum1 += i
	//	}
	//}
	//fmt.Println(count)
	//fmt.Println(sum1)
	//
	//sum2 := 1
	//for i := 1; i <= 5; i++ {
	//	sum2 *= i
	//}
	//fmt.Println(sum2)

	//var row = 4
	//var column = 4
	//for i := 0; i < row; i++ {
	//	for j := 0; j < column; j++ {
	//		fmt.Print("*\t")
	//	}
	//	fmt.Println("")
	//}

	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		//fmt.Print(i, j, i*j)
	//		fmt.Printf("%v*%v=%v\t", j, i, i*j)
	//	}
	//	fmt.Println("")
	//}

	//var str = "hello,golang"
	//for k, v := range str {
	//	//fmt.Println(k, v)
	//	fmt.Printf("%v--%c\n", k, v)
	//}

	//var arr = []string{"java", "php", "golang", "nodejs", " ", "mysql"}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	//var arr = []string{"java", "php", "golang", "nodejs", " ", "mysql"}
	//for key, val := range arr {
	//	fmt.Println(key, val)
	//}

	//var ext = ".js"
	//scope diff
	//switch ext := ".js"; ext {
	//case ".html":
	//	fmt.Println("text/html")
	//	break
	//case ".css":
	//	fmt.Println("text/css")
	//	break
	//case ".js":
	//	fmt.Println("text/javascript")
	//	break
	//default:
	//	fmt.Println("not found resource")
	//}

	/**
	switch case -> base use
	*/

	//var age = 18
	//switch {
	//case age < 24:
	//	fmt.Println("study")
	//	fallthrough
	//case age >= 24 && age <= 60:
	//	fmt.Println("work")
	//	fallthrough
	//case age > 60:
	//	fmt.Println("pleasure")
	//	fallthrough
	//default:
	//	fmt.Println("error")
	//}

	//switch fallthrough
	//fallthrough once again 1
	//for i := 0; i < 10; i++ {
	//	if i == 2 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	//label:
	//	for i := 0; i < 10; i++ {
	//		for j := 0; j < 10; j++ {
	//			if j == 3 {
	//				break label
	//			}
	//			fmt.Println(i, j)
	//		}
	//	}

	//label:
	//	for i := 0; i < 4; i++ {
	//		for j := 0; j < 10; j++ {
	//			if j == 3 {
	//				continue label
	//			}
	//			fmt.Println(i, j)
	//		}
	//	}

	//	var n = 30
	//	if n > 18 {
	//		fmt.Println("person")
	//		goto label
	//	}
	//	fmt.Println("a")
	//	fmt.Println("b")
	//label:
	//	fmt.Println("c")
	//var arr1 [3]int
	//var arr2 [4]int
	//var strArr [3]string
	//
	//fmt.Printf("%T\n", arr1)
	//fmt.Printf("%T\n", arr2)
	//fmt.Printf("%T\n", strArr)

	//var arr1 = [3]int{1, 18, 20}
	//fmt.Println(arr1)
	//
	//var strArr1 = [3]string{"php", "java", "golang"}
	//fmt.Println(strArr1)
	//
	//arr2 := [3]int{1, 18, 20}
	//fmt.Println(arr2)
	//
	//strArr2 := [3]string{"php", "java", "golang"}
	//fmt.Println(strArr2)
	//
	//var arr [3]int
	//arr[0] = 1
	//arr[1] = 18
	//arr[2] = 20
	//
	//fmt.Println(arr)
	//
	//var strArr [3]string
	//strArr[0] = "php"
	//strArr[1] = "java"
	//strArr[2] = "golang"
	//
	//fmt.Println(strArr)
	//
	//var arr3 = [...]int{1, 241, 552, 66, 42, 4, 7, 5, 7, 2}
	//arr3[1] = 1 //change val
	//fmt.Println(arr3)
	//fmt.Println(len(arr3))
	//
	//arr4 := [...]int{0: 1, 1: 10, 2: 20, 5: 50}
	//fmt.Println(arr4)
	//fmt.Println(len(arr4))//6

	//arr := [...]string{"php", "java", "golang"}
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	arr := [...]string{"php", "java", "golang"}
	for k, v := range arr {
		fmt.Println(k, v)
	}
}
