package common

import "fmt"

func val(d [3]byte) [3]byte {
	fmt.Printf("val: %p\n", &d)

	d[0] += 100
	return d
}

func ptr(p *[3]byte) *[3]byte {
	fmt.Printf("ptr: %p\n", p)

	p[0] += 200
	return p
}

func ArrayOpt() {
	d := [...]byte{1, 2, 3}
	d2 := d
	d3 := *(&d)

	fmt.Printf(" d: %p\n", &d)
	fmt.Printf("d2: %p\n", &d2)
	fmt.Printf("d3: %p\n", &d3)

	// ---------------------

	d4 := val(d)
	fmt.Printf("val.ret: %p\n", &d4)

	p := ptr(&d)
	fmt.Printf("val.ret: %p\n", p)

	// ---------------------

	fmt.Printf("d: %v\n", d)
}
