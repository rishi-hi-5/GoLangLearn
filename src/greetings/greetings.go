package greetings

import (
	"fmt"
)

func HelloSir() string {
	packagePrivateMethod()
	return "Hello Sir"
}

func packagePrivateMethod() {
	fmt.Println("you cannot call packagePrivate method")
}
