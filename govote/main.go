package govote

import (
	"fmt"
	"os"
)

func main() {
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println(home)
	fmt.Println(user)
	fmt.Println(gopath)
}
