# learn-golang
learn golang repository, Simple REST API using Gin Framework

**Documentation**

1. masuk go.dev, download dulu file installation go
2. cek “go version”
3. download vscode, install extension golang, go doc
4. run “go env -w GO111MODULE=auto” biar debug dr vscode berjalan lancar
5. buat folder “learn-golang”
6. init repo “go mod init example/learn-golang”
7. install gin “go get github.com/gin-gonic/gin” (FRAMEWORK)
8. install gorm "go get -u gorm.io/gorm" (ORM)
9. ikuti struktur folder spt di repo ini, buat mulai dari database, model, controller, baru kemudian main.go
10. ganti config DB di file "database/db.go"
11. jalankan server dengan "go run main.go"
12. buka postman atau terminal (curl) utk mengetest endpoint

routes:
```
// get all books
get “books” -> curl localhost:8081/books
// get book by ID
get “books/:id” -> curl localhost:8081/books/2
// create new book
post “books” -> curl localhost:8081/books --include --header "Content-Type: application/json" --request "POST" --data-raw '{
    "title": "New Book",
    "author": "New Author",
    "quantity": 5
}'
// update book
put "books/:id" -> curl --location --request PUT 'localhost:8081/books/4' --header 'Content-Type: application/json'
--data-raw '{
    "title": "Updated Book",
    "author": "Updated Author"
}'
// checkout book, book - 1
post “checkout/:id” -> curl localhost:8081/checkout/2 —request “PATCH”
// return book, book + 1
post “return/:id” -> curl localhost:8081/return/2 —request “PATCH”
```
