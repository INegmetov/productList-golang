# Product List App

Тестовое задание - Список продуктов

Скопировать репозbторий коммандой `git clone https://github.com/INegmetov/productList-golang.git`

Запустить Docker 

Запустить docker контейнер с базов postgres `docker run --name=productsList-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d -rm postgres`

Выполнить миграции `migrate -path ./schema --database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up`

Запустить файл main.go `go run cmd/main/go`

Перейт по пути в SwaggerInfo `http://localhost:8080/swagger/index.html#/`