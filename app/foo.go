package app

import (
	"fmt"

	"github.com/CrowdSurge/banner"
)

// Foo is a bogus function
func Foo() {
	fmt.Println("This is foo")
	banner.Print("VENDOR")
}
