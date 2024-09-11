default: build test run

build:
	tailwindcss -i input.css -o assets/style.css

test:
	go test

run: 
	go run main.go

watch:
	air --build.cmd "just"
