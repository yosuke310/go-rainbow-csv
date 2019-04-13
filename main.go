package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	black int = 30 + iota
	red
	green
	yellow
	blue
	magenda
	cyan
	white
)

var rainbow = []int{red, green, yellow, blue, magenda, cyan, white}

func main() {
	filepath := flag.String("f", "", "filepath")
	delimiter := flag.String("d", ",", "delimiter")
	flag.Parse()

	var fp *os.File
	if *filepath == "" {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(*filepath)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := csv.NewReader(fp)
	reader.Comma = []rune(*delimiter)[0]
	reader.LazyQuotes = true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		line := ""
		for i, field := range record {
			line += fmt.Sprintf("\x1b[%vm%s", rainbow[i%7], field)
			if i < len(record)-1 {
				line += *delimiter
			}
			line += "\x1b["
		}
		fmt.Println(line)
	}
}
