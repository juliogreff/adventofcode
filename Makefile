YEAR ?= 2022
DAY ?= 01
PART ?= 1

run:
	go run ./${YEAR}/${DAY}/part${PART}/main.go ./${YEAR}/${DAY}/input

test:
	go run ./${YEAR}/${DAY}/part${PART}/main.go ./${YEAR}/${DAY}/input_test

day:
	mkdir -p ${YEAR}/${DAY}/part{1,2}
	cp template.go ${YEAR}/${DAY}/part1/main.go
	cp template.go ${YEAR}/${DAY}/part2/main.go
	touch ${YEAR}/${DAY}/{input,input_test,README.markdown}

