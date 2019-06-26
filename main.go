package main

import (
	"codes/echo"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(echo.Echo(strings.Join(os.Args[1:], " ")))
}
