test::
	go run ./$(name)/$(name).go --profile
	[ -e heap.prof ]
	[ -e cpu.prof ]
	go run ./$(name)/$(name).go --trace
	[ -e trace.prof ]
