package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 7, "node own ID")
	flag.Parse()

	node := func() *node {
		if len(flag.Args()) == 0 {
			return newRoot(id)
		}

		rootID, errConv := strconv.Atoi(flag.Arg(0))
		if errConv != nil {
			fmt.Printf("Conversion of root ID: %s", errConv.Error())
			os.Exit(1)
		}

		return newNode(id, rootID)
	}
}
