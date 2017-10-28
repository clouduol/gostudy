// Digest test digest features
package main

import (
	_ "crypto/sha256"
	_ "crypto/sha512"
	"fmt"
	"github.com/opencontainers/go-digest"
)

/**
* digest.go
* type Digest string
 */
type ChainID digest.Digest
type DiffID digest.Digest

func main() {
	id := digest.FromBytes([]byte("my content"))
	fmt.Printf("%v\t%[1]T\n", id)
	id = digest.FromBytes([]byte("cec7521cdf36a6d4ad8f2e92e41e3ac1b6fa6e05be07fa53cc84a63503bc5700 453fc2d51e11412f191e21f29cd098cc912995076b1bbf0931f228adc33b039d"))
	fmt.Printf("%v\t%[1]T\n", id)
	id = digest.FromBytes([]byte("cec7521cdf36a6d4ad8f2e92e41e3ac1b6fa6e05be07fa53cc84a63503bc5700 sha256:453fc2d51e11412f191e21f29cd098cc912995076b1bbf0931f228adc33b039d"))
	fmt.Printf("%v\t%[1]T\n", id)
	// the second lowest layer, parent and diffid both begin with "sha256"
	id = digest.FromBytes([]byte("sha256:cec7521cdf36a6d4ad8f2e92e41e3ac1b6fa6e05be07fa53cc84a63503bc5700 sha256:453fc2d51e11412f191e21f29cd098cc912995076b1bbf0931f228adc33b039d"))
	fmt.Printf("%v\t%[1]T\n", id)
	// other layers, parent has no prefix "sha256", diffid has
	id = digest.FromBytes([]byte(string(id) + " " + "sha256:a1a53f8d99b57834ca1d6370a3988d4bbd4a5235d5ff3741d0d6ecdd099872d7"))
	fmt.Printf("%v\t%[1]T\n", id)
}
