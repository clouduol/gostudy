// Sha256 shows array characteristics
package main

import "fmt"
import "crypto/sha256"

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	a := [2]int{1, 2}
	fmt.Printf("%d\n%[1]x\n", a)

	b := [2]byte{15, 16}
	fmt.Printf("%d\n%[1]x\n", b)

	c := [...]int{4: 3}
	for i := range c {
		fmt.Printf("%d", c[i])
		if i != len(c)-1 {
			fmt.Printf("\t")
		}
	}
	fmt.Printf("\n")

	d := &c
	for i := range d {
		fmt.Printf("%d", d[i])
		if i != len(d)-1 {
			fmt.Printf("\t")
		}
	}
	fmt.Printf("\n")
	//fmt.Printf("%d\n%d\n%d\n", c[0], d[0], *d[0])
	fmt.Printf("%d\n%d\n%d\n%d\n", c[0], d[0], *d, (*d)[0])
	fmt.Printf("%T\t%T\n", c, d)

	fmt.Printf("stub\n")
	fmt.Printf("%x\n%x\n", &c[0], d)
	fmt.Printf("%x\n%x\n", &d[0], d)

	SetFirstToOne(&c)
	fmt.Printf("%d\n", c)
	TestParam(c)
	fmt.Printf("%d\n", c)

	fmt.Println("test empty array")
	var x [1]int
	//fmt.Println("%t\t%d\n", x == nil, len(x))
	fmt.Printf("%d\t%d\n", len(x), x[0])
}

func SetFirstToOne(ptr *[5]int) {
	ptr[0] = 1
}

func TestParam(arr [5]int) {
	arr[1] = 1
	fmt.Printf("%d\n", arr)
}
