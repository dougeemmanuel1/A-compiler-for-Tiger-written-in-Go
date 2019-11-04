all: clean
	tools/goyacc tiger.y
	go build -o tigerc
clean:
	rm -f 	y.output
	rm -rf ./tigerc
	rm -rf ./y.go
