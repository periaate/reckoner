package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	search  string
	searchl string
	vl      bool
)

func main() {
	s := flag.String("s", "", "")
	sl := flag.String("sl", "", "")

	flag.Parse()
	if *s != "" && *sl != "" {
		vl = true
		search = *s
		searchl = *sl
	}
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	// I didn't know this was possible, I know that the syntax was possible.
	// But didn't know that the functionality was. Good find! 🤓
	if info.Mode()&os.ModeCharDevice != 0 {
		panic("not pipe")
	}

	reckon(os.Stdin)
}

type datamap map[string]interface{}

func reckon(in *os.File) {
	dec := json.NewDecoder(in)

	for {
		data := datamap{}
		err := dec.Decode(&data)
		if err == io.EOF {
			return
		}

		for _, k := range os.Args {
			v, f := data[k]
			if !f {
				continue
			}

			if !vl {
				fmt.Println(v)
			} else if strings.Contains(strings.ToLower(fmt.Sprint(v)), strings.ToLower(search)) {
				fmt.Println(data[searchl])
				break
			}
		}
	}
}
