package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var answer int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineHex := strings.Split(scanner.Text(), "")
		var lineBin string
		for _, hex := range lineHex {
			n, err := strconv.ParseInt(hex, 16, 8)
			if err != nil {
				panic(err)
			}

			lineBin += fmt.Sprintf("%04b", n)
		}

		answer, _ = parsePacket(lineBin)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", answer)
}

func parsePacket(str string) (int64, string) {
	version := parseBin(str[0:3])
	t := str[3:6]
	body := str[6:]

	if t == "100" {
		var binStr string
		for {
			c := body[0]
			binStr += body[1:5]
			body = body[5:]

			if c == '0' {
				break
			}
		}

		return version, body
	} else {
		mode := body[0:1]

		var v int64
		switch mode {
		case "0":
			l := parseBin(body[1:16])
			subpackets := body[16 : 16+l]
			body = body[16+l:]

			for len(subpackets) >= 4 {
				v, subpackets = parsePacket(subpackets)
				version += v
			}

			return version, body

		case "1":
			n := parseBin(body[1:12])
			body = body[12:]

			for i := 0; i < int(n); i++ {
				v, body = parsePacket(body)
				version += v
			}

			return version, body
		}
	}

	return -1, ""
}

func parseBin(str string) int64 {
	n, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}

	return n
}
