prepare:
	@GO111MODULE=on go mod download

generate:
	go run cmd/generate/main.go
genesis:
	go run cmd/genesis/main.go


.PHONY: views

views:
	cd views&&yarn next -p 3000
