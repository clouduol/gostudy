// TempconvTest test package tempconv
package main

import "fmt"
import "tempconv"

func main() {
	fmt.Printf("%v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("%v\n", tempconv.FreezingC)
	fmt.Printf("%v\n", tempconv.BoilingC)
	fmt.Printf("%v\n", tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Printf("%v\n", tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("%v\n", tempconv.CToF(tempconv.BoilingC))
}
