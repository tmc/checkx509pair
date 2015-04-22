package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var (
	key  = flag.String("key", "", "path to private key")
	cert = flag.String("cert", "", "path to cert")
)

func main() {
	flag.Parse()
	if *key == "" || *cert == "" {
		flag.Usage()
		os.Exit(1)
	}

	c, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	spew.Dump(c)
}
