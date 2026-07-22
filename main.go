package main

import (
	"bufio"
	"flag"
	"fmt"
	"ltz/modules"
	"os"
	"strings"
)

var fil = flag.String("f", "", "Write here your filename for pipeline")

func main() {
	flag.Parse()
	if *fil == "" {
		fmt.Println("Please try again write filename")
		os.Exit(1)
	}
	filt, err := getmsg()
	if err != nil {
		fmt.Println("You have error")
		os.Exit(1)
	}
	filestruc := modules.Filestruct{
		Filename: *fil,
		Filter:   filt,
	}
	modules.Workwithfile(filestruc)
}
func getmsg() (string, error) {
	fmt.Print("Write your filter:")
	reader := bufio.NewReader(os.Stdin)
	zx, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("You have error %w", err)
	}
	return strings.TrimSpace(zx), nil
}
