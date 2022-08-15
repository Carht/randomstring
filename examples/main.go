package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"github.com/carht/randomstring"
)

func version() {
	version := "0.1.1"
	fmt.Println(version)
}

func usage() {
	usage := `randomstring [Argument]
                  
                  randomstring:
                    Without arguments return a string of 32 characters between uppercases, lowercases and numbers.
                  
                  randomstring version | randomstring -v:
                    Return the version of this example.

                  randomstring help | randomstring -h:
                    Return this help.`

	fmt.Println(usage)
}


func main() {
	abc := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if len(os.Args) > 1 {
		arg := os.Args[1]

		switch arg {
		case "version":
			version()
		case "-v":
			version()
		case "help":
			usage()
		case "-h":
			usage()
		default:
			inputLenStr, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatal(err)
			}

			randStr := randomstring.GenRandStr(abc, inputLenStr)
			fmt.Println(randStr)
		}
	} else {
		randStr := randomstring.GenRandStr(abc, 32)
		fmt.Println(randStr)
	}
}
