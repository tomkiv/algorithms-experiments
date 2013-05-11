package main

import (
  "bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var from uint64
var to uint64

var data map[uint64]bool
var res map[uint64]bool

func hasSum(number uint64) {
	for i := from; i <= to; i++ {
		if i == number*2 {
			continue
		}
		_, ok := data[i-number]
		if ok {
			res[i] = true
		}
	}
}

func readFile() {
	fi, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	k := 0
	defer fi.Close()
	r := bufio.NewReader(fi)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		number, err := strconv.ParseUint(strings.Trim(line, " \n'\""), 10, 0)
		if err != nil {
			break
		}
		hasSum(number)
		data[number] = true
		k = k + 1
	}
}

var fileName = flag.String("file", "", "File name")

func main() {
	from = 2500
	to = 4000
	data = make(map[uint64]bool)
	res = make(map[uint64]bool)
	flag.Parse()
	readFile()
	fmt.Println(len(res))
}
