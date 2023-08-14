package main

import "fmt"

func testGoto() {
	// start:
	// label start defined and not used

	for i := 0; i < 3; i++ {
		if i > 1 {
			goto exit
		}
		println(i)
	}
exit:
	testJump()
}

func testJump() {
	fmt.Println("goto exec jump to here")
}

/*
break: 终止 switch/for/select
continue：仅仅用于for，进入下一层循环
*/
func testContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 立即进入下轮循环。(goto next)
		}

		if i > 5 {
			break // 立即终止整个循环。
		}

		println(i)
	}
}

// todo  outer/inner 是用来干嘛的
func testOuter() {
outer:
	for x := 0; x < 10; x++ {

	inner:
		for y := 0; y < 10; y++ {
			if x%2 == 0 {
				continue outer
			}

			if y > 3 {
				println()
				break inner
			}

			print(x, ":", y, " ")
		}
	}
}
func main() {

	//testGoto()
	//testContinue()
	testOuter()
}
