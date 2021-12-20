package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Pair interface {
	Left() int
	Right() int
	AddLeft(int)
	AddRight(int)
	IsNumber() bool
	String() string
	Magnitude() int
}

type pair struct {
	l Pair
	r Pair
}

func (p pair) Left() int {
	return p.l.Right()
}

func (p pair) Right() int {
	return p.r.Left()
}

func (p pair) AddLeft(x int) {
	p.l.AddLeft(x)
}

func (p pair) AddRight(x int) {
	p.r.AddRight(x)
}

func (p pair) IsNumber() bool {
	return false
}

func (p pair) Magnitude() int {
	return p.l.Magnitude()*3 + p.r.Magnitude()*2
}

func (p pair) String() string {
	return fmt.Sprintf("[%s,%s]", p.l.String(), p.r.String())
}

type number struct {
	self int
}

func (n *number) AddLeft(x int) {
	n.self += x
}

func (n *number) AddRight(x int) {
	n.self += x
}

func (n number) Left() int {
	return n.self
}

func (n number) Right() int {
	return n.self
}

func (n number) IsNumber() bool {
	return true
}

func (n number) Magnitude() int {
	return n.self
}

func (n number) String() string {
	return strconv.Itoa(n.self)
}

func main() {
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var numbers []Pair

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var stack []Pair

		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			var popped Pair
			switch line[i] {
			case '[':
				stack = append(stack, &pair{})
			default:
				stack = append(stack, &number{parseInt(line[i : i+1])})
			case ']':
				stack, popped = pop(stack)
				last := stack[len(stack)-1]
				last.(*pair).r = popped
			case ',':
				stack, popped = pop(stack)
				last := stack[len(stack)-1]
				last.(*pair).l = popped
			}
		}

		if len(stack) > 1 {
			panic("got more than one left in the stack")
		}

		numbers = append(numbers, stack[0])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for len(numbers) > 1 {
		answer := add(numbers[0], numbers[1])
		more := true
		for more {
			answer, more = reduceAndExplode(answer, nil, nil, 1)

			if more {
				continue
			}

			answer, more = reduceAndSplit(answer, nil, nil, 1)
		}
		numbers = append([]Pair{answer}, numbers[2:]...)
	}

	mag := numbers[0].Magnitude()

	fmt.Printf("%d\n", mag)
}

func add(l, r Pair) Pair {
	return &pair{l, r}
}

func reduceAndExplode(p, left, right Pair, depth int) (Pair, bool) {
	if p.IsNumber() {
		return p, false
	}

	if depth >= 5 {
		return explode(p, left, right), true
	} else {
		var more bool
		p := p.(*pair)
		l := p.l
		r := p.r

		l, more = reduceAndExplode(p.l, left, p.r, depth+1)
		if !more {
			r, more = reduceAndExplode(p.r, p.l, right, depth+1)
		}

		return &pair{l, r}, more
	}
}

func reduceAndSplit(p, left, right Pair, depth int) (Pair, bool) {
	if p.IsNumber() {
		p := p.(*number)
		if p.self >= 10 {
			return split(p), true
		}

		return p, false
	} else {
		var more bool
		p := p.(*pair)
		l := p.l
		r := p.r

		l, more = reduceAndSplit(p.l, left, p.r, depth+1)
		if !more {
			r, more = reduceAndSplit(p.r, p.l, right, depth+1)
		}

		return &pair{l, r}, more
	}
}

func explode(p, left, right Pair) Pair {
	if left != nil {
		left.AddRight(p.Left())
	}

	if right != nil {
		right.AddLeft(p.Right())
	}

	return &number{0}
}

func split(p *number) Pair {
	half := float64(p.self) / 2
	return &pair{
		l: &number{int(math.Floor(half))},
		r: &number{int(math.Ceil(half))},
	}
}

func pop(stack []Pair) ([]Pair, Pair) {
	last := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, last
}

func parseInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return n
}
