package modules

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Filestruct struct {
	Filename string
	Filter   string
}
type Read interface {
	Readfile() (<-chan string, error)
}
type Filt interface {
	Filt(<-chan string) <-chan string
}
type Fl interface {
	Read
	Filt
}
type Jsonstruct struct {
	Filename string
	Filter   string
}

func ras(filenames, filters string) Fl {
	zxc := strings.SplitN(filenames, ".", 2)
	switch zxc[1] {
	case "json":
		return Jsonstruct{Filename: filenames, Filter: filters}
	case "txt", "log":
		return Filestruct{Filename: filenames, Filter: filters}
	default:
		return nil
	}
}
func (js Jsonstruct) Readfile() (<-chan string, error) {
	file, err := os.Open(js.Filename)
	if err != nil {
		return nil, fmt.Errorf("The error:%w", err)
	}
	out := make(chan string)
	reader := bufio.NewReader(file)
	go func() {
		defer close(out)
		defer file.Close()
		for {
			str, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			var data map[string]interface{}
			err = json.Unmarshal([]byte(str), &data)
			var path []string
			for i, el := range data {
				path = append(path, fmt.Sprintf("%s:%v", i, el))
			}
			out <- strings.Join(path, " ")
		}
	}()
	return out, nil
}
func (js Jsonstruct) Filt(out <-chan string) <-chan string {
	exits := make(chan string)
	go func() {
		defer close(exits)
		for el := range out {
			if strings.Contains(strings.ToLower(el), strings.ToLower(js.Filter)) {
				exits <- el
			}
		}
	}()
	return exits
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
func Workwithfile(f Fl) {
	fl, err := f.Readfile()
	if err != nil {
		return
	}
	a := f.Filt(fl)
	for el := range a {
		fmt.Println(el)
	}
}
