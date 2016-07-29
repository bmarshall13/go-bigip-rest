package f5api

import (
	"fmt"
)

// To use this library, create a Client object, and use it to access the API
func Example() {
	user := "admin"
	pass := "admin"
	tlsNoVerify := true
	host := "127.0.0.1"

	f5 := NewClient(host, user, pass, tlsNoVerify)
	err := f5.DoLogin()
	if err != nil {
		panic(fmt.Sprintf("Error logging in: %v", err))
	}

	// Example: list all the virtuals (compare to "show /ltm virtual")
	virtuals, _, err := f5.Ltm.GetVirtualList()
	for _, virtual := range virtuals.Items {
		fmt.Printf("Virtual %v Destination %v\n", virtual.Name, virtual.Destination)
	}
	// Will print something like:
	// Virtual my_ltm_virtual Destination /Common/1.2.3.5:443
}
