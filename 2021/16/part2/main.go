package main

import (
	"bufio"
	"fmt"
	"math"
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
	_ = parseBin(str[0:3]) // version is unused
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

		return parseBin(binStr), body
	} else {
		var (
			o        int64
			operands []int64
		)

		switch body[0:1] {
		case "0":
			l := parseBin(body[1:16])
			subpackets := body[16 : 16+l]
			body = body[16+l:]

			for len(subpackets) >= 4 {
				o, subpackets = parsePacket(subpackets)
				operands = append(operands, o)
			}

		case "1":
			n := parseBin(body[1:12])
			body = body[12:]

			for i := 0; i < int(n); i++ {
				o, body = parsePacket(body)
				operands = append(operands, o)
			}
		}

		var result int64
		switch t {
		case "000":
			result = sum(operands)
		case "001":
			result = product(operands)
		case "010":
			result = min(operands)
		case "011":
			result = max(operands)
		case "101":
			result = gt(operands)
		case "110":
			result = lt(operands)
		case "111":
			result = eq(operands)
		default:
			panic(fmt.Sprintf("operator %d not implented", parseBin(t)))
		}

		return result, body
	}

	return -1, ""
}

func sum(operands []int64) int64 {
	var sum int64
	for _, o := range operands {
		sum += o
	}
	return sum
}

func product(operands []int64) int64 {
	var product int64 = 1
	for _, o := range operands {
		product *= o
	}
	return product
}

func min(operands []int64) int64 {
	var min int64 = math.MaxInt64
	for _, o := range operands {
		if o < min {
			min = o
		}
	}
	return min
}

func max(operands []int64) int64 {
	var max int64 = math.MinInt64
	for _, o := range operands {
		if o > max {
			max = o
		}
	}
	return max
}

func gt(operands []int64) int64 {
	if operands[0] > operands[1] {
		return int64(1)
	} else {
		return int64(0)
	}
}

func lt(operands []int64) int64 {
	if operands[0] < operands[1] {
		return int64(1)
	} else {
		return int64(0)
	}
}

func eq(operands []int64) int64 {
	if operands[0] == operands[1] {
		return int64(1)
	} else {
		return int64(0)
	}
}

func parseBin(str string) int64 {
	n, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}

	return n
}
