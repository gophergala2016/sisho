go-bindata -o bindata.go -pkg sisho templates/...
go build -o sisho cmd/main.go

zip -9 -r ./sisho.zip ./sisho
# ghr -u kogai $(git describe) sisho.zip
