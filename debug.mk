test::
	go run ./peephole.go --profile
	[ -e heap.prof ]
	[ -e cpu.prof ]
	go run ./peephole.go --trace
	[ -e trace.prof ]
