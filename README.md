# learn-golang
learn golang repository, Simple REST API using Gin Framework

**Documentation**

1. masuk go.dev, download dulu file installation go
2. cek “go version”
3. download vscode, install extension golang, go doc
4. run “go env -w GO111MODULE=auto” biar debug dr vscode berjalan lancar
5. buat folder “learn-golang”
6. init repo “go mod init example/learn-golang”
7. install gin “go get github.com/gin-gonic/gin”

routes:
```
get “books” -> curl localhost:8081/books

get “books/:id” -> curl localhost:8081/books/2

post “books” -> curl localhost:8081/books --include --header "Content-Type: application/json" -d @body.json --request "POST"

post “checkout” -> curl localhost:8081/checkout?id=2 —request “PATCH”

post “return” -> curl localhost:8081/return?id=2 —request “PATCH”
```
