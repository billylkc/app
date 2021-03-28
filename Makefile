app_w: main.go
	GOOS=windows GOARCH=386 go build -o bin/app.exe main.go
