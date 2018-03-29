test::
	go run ./go-boilerplate/go-boilerplate.go --profile
	[ -e heap.prof ]
	[ -e cpu.prof ]
	go run ./go-boilerplate/go-boilerplate.go --trace
	[ -e trace.prof ]
