produce:
	@go run cmd/producer/main.go --length=7 --amount=500 --path storage/tokens

consume:
	@go run cmd/consumer/main.go --path storage/tokens
