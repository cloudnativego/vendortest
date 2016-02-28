package integrations_test

import (
	"fmt"
	"testing"

	"github.com/cloudnativego/vendortest/app"
)

func TestIntegration(t *testing.T) {
	fmt.Println("Integration test")

	app.Foo()
}
