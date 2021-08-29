binary="server"
mkdir -p dist

cd "api"
go get -t -d -v ./...
go build -a -o ${binary}

mv ${binary} "../dist/"
