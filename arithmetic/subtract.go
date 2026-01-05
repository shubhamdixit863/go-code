package arithmetic

import "fmt"

// one directory one package
// make sure that directory name is same as package name

const pI = 3.14 // not exported (means we cannot use it in other packages )

const PLANCK_CONSTANT = 3.71 //  exported (means we can use it in other packages )

func add() {
	fmt.Println(2 + 2)
}
