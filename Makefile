run-serial:
	mkdir -p ./repos
	go run .

run-parallels:
	mkdir -p ./repos
	go run . -parallels

cleanup:
	rm -rf ./repos/*