binary="server"
mkdir -p dist

cd "api"
go build -o ${binary}

mv ${binary} "../dist/"