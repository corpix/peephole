test::
	go run ./$(name)/$(name).go --profile
	[ -e heap.prof ]
	[ -e cpu.prof ]
	go run ./$(name)/$(name).go --trace
	[ -e trace.prof ]

.PHONY: statsd
statsd:
	docker run --rm -p 8125:8125/udp dasch/statsd-debug
