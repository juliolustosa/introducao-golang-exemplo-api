## nome do executavel
BINARY=api-server

default: build

## compila para a plataforma atual
build:
	go mod tidy
	go build -o ../bin/${BINARY}

## compila para as outras plataformas
build-all:
	GOOS=windows GOARCH=amd64 go build -o ../bin/windows-amd64/${BINARY}.exe
	GOOS=linux GOARCH=amd64 go build -o ../bin/linux-amd64/${BINARY}
	GOOS=darwin GOARCH=amd64 go build -o ../bin/macOS-amd64/${BINARY}
	GOOS=darwin GOARCH=arm64 go build -o ../bin/macOS-arm64/${BINARY}