package main

import (
	"fmt"
	"io"
	"os"

	"github.com/archmangler/websocketfun/fun"
	"github.com/archmangler/websocketfun/socket"
	"github.com/archmangler/websocketfun/web"
)

func displayGreetings(w io.Writer) {

	fmt.Fprintln(w, web.Web())
	fmt.Fprintln(w, socket.Socket())
	fmt.Fprintln(w, fun.Fun())

}

func main() {

	fmt.Println("This is the main package.")
	displayGreetings(os.Stdout)

}
