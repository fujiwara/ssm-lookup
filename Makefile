.PHONY: test install

test:
	go test ./...

install:
	go install github.com/fujiwara/ssm-lookup/cmd/ssm-lookup
