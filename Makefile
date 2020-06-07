backend:
	@go run ./pkg/cmd/main.go

frontend:
	@/bin/bash -c "cd $GOPATH/src/github.com/c-sinclair/restic/pkg/http/web/app && yarn serve" 