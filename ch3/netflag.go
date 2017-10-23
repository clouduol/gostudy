// Netflag test iota const generator
package main

import (
	"fmt"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	a = 2
	b
	c
	d = true
	e
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const (
	Smode  = '世'
	Umode  = '\u4e16'
	Xmode  = '\x12'
	Xmodes = "\xe4\xb8\x96"
	Smodes = "世"
	Umodes = "\u4e16"
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
	fmt.Printf("%v %v %v\n", a, b, c)
	fmt.Printf("%v %v\n", d, e)
	fmt.Printf("%v\n", TiB)
	//fmt.Printf("%v\n", ZiB)
	fmt.Printf("%v\n", PiB/TiB)
	fmt.Printf("%T %[1]x\n", Xmode)
	fmt.Printf("%T %[1]x\n", Smode)
	fmt.Printf("%T %[1]x\n", Umode)
	fmt.Printf("%T %[1]x\n", Xmodes)
	fmt.Printf("%T %[1]x\n", Smodes)
	fmt.Printf("%T %[1]x\n", Umodes)
}
