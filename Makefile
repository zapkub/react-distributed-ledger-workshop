prepare:
	@GO111MODULE=on go mod download
	cd views && yarn install

generate:
	@GO111MODULE=on go run cmd/generate/main.go
genesis:
	@GO111MODULE=on go run cmd/genesis/main.go


.PHONY: views

views:
	cd views&&yarn next -p 3000
