YEAR ?= 2021
DAY ?= 01
PART ?= 1

run:
	go run ./${YEAR}/${DAY}/part${PART}/main.go ./${YEAR}/${DAY}/input

test:
	go run ./${YEAR}/${DAY}/part${PART}/main.go ./${YEAR}/${DAY}/input_test
