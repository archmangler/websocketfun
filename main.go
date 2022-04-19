package main

import (
  "fmt"
   "github.com/archmangler/websocketfun/web"
   "github.com/archmangler/websocketfun/socket"
   "github.com/archmangler/websocketfun/fun"
)

func main() {

	fmt.Println("This is the main package")
        fmt.Println(Web.Web())
        fmt.Println(Socket.Socket())
        fmt.Println(Fun.Fun())

}
