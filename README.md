# TTL-monitoring

## Задание
Реализовать сервис для мониторинга TTL сертификатор заданных URL для возможности выявлять сертификаты, которые близки к истечению и планировать его перевыпуск.

Функционал сервиса:
1. добавление отслеживаемых URL в базу (HTTP запрос методом POST)
2. вывод всех отслеживаемых URL из базы с выводом для каждого TTL сертификата
3. вывод TTL сертификата по идентификатору URL
4. вывод TTL сертификата ресурса без сохранения в базу данных

Сервис должен быть реализован на Go версии не ниже 1.19, базу данных использовать PostgreSQL, так же необходимо написать Dockerfile для сборки в контейнер.


## Запуск
    Из ./deployments


    создание go.mod и go.sum (для Windows с PowerShell)
    ```
    docker run --rm -v "${PWD}:/app" -w /app golang:1.25-alpine sh -c "go mod init app && go mod tidy"
    ```

    ```
    docker compose --env-file ../configs/.env up --build -d
    ```

## Миграции
    инициализация миграций:
    ```
    migrate create -ext sql -dir migrations/ -seq init_mg
    ```

    применение:
    ```
    migrate -path migrations// -database "postgresql://tls_monitoring:tls_monitoring_password@localhost:5433/tls_monitoring?sslmode=disable" -verbose up
    ```
    
    шаблон:
    ```
    migrate -path migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
    ```

## Swagger
    http://localhost:8080/swagger/index.htm