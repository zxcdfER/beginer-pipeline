package modules

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Filestruct struct {
	Filename string
	Filter   string
}
type Fl interface {
	Readfile() (<-chan string, error)
	Filt(<-chan string) <-chan string
}

func (Zxcm Filestruct) Filt(out <-chan string) <-chan string {
	exits := make(chan string)
	go func() {
		defer close(exits)
		for el := range out {
			el = strings.TrimSpace(el)
			filt := strings.TrimSpace(Zxcm.Filter)
			if strings.Contains(strings.ToLower(el), strings.ToLower(filt)) {
				exits <- el
			}
		}
	}()
	return exits
}
func (Filestru Filestruct) Readfile() (<-chan string, error) {
	out := make(chan string)
	file, err := os.Open(Filestru.Filename)
	if err != nil {
		return nil, fmt.Errorf("The error is %w", err)
	}
	reader := bufio.NewReader(file)

	go func() {
		defer file.Close()
		defer close(out)
		for {
			data, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			out <- strings.TrimSpace(data)
		}
	}()
	return out, nil
}
func Woкkwithfile(f Fl) {
	fl, err := f.Readfile()
	if err != nil {
		return
	}
	a := f.Filt(fl)
	for el := range a {
		fmt.Println(el)
	}
}
