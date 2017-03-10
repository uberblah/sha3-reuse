package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/uberblah/sha3-reuse"
	"golang.org/x/crypto/sha3"
)

var hash sha3.ShakeHash

func write(data string) {
	fmt.Printf("writing '%s'...\n", data)
	n, e := hash.Write([]byte(data))
	fmt.Printf("wrote %d bytes, err=%#v\n", n, e)
}

func read(toread int) {
	fmt.Printf("reading %d bytes...\n", toread)
	data := make([]byte, toread)
	n, e := hash.Read(data)
	if n == toread && e == nil {
		fmt.Printf("%s\n", hex.EncodeToString(data))
	}
	fmt.Printf("got %d bytes, err=%#v\n", n, e)
}

func main() {
	hash = sha3r.NewSHA3rm()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("choose action (r/w): ")
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()

		switch line {
		case "w":
			fmt.Printf("enter data: ")
			if !scanner.Scan() {
				return
			}
			write(scanner.Text())
		case "r":
			for {
				fmt.Printf("enter nbytes to read: ")
				if !scanner.Scan() {
					return
				}
				n, err := strconv.Atoi(scanner.Text())
				if err == nil {
					read(n)
					break
				}
				fmt.Println("invalid integer. try again.")
				continue
			}
		default:
			fmt.Println("invalid command. try again.")
			continue
		}
	}
}
