export   GOOS=linux
export   GOARCH=amd64
export  CGO_ENABLED=0
go build  -o qpgate main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-qpgate-linux main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o go-qpgate-win.exe main.go