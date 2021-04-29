# Запуск проекта

    go build ./cmd/main.go

# env
    PORT=8080
    DATABASE=host\=localhost user\=postgres password\=qwerty123 dbname\=petitions port\=5432
    EMAIL_VALIDATOR_TOKEN=d6ba3449152602cbd63d215ed1ee612e6b3134eb
    EMAIL_VALIDATOR_SECRET=fbce4c61182373033f9184812ef91cb7180a05d6

# Инициация базы без миграции
    CREATE TABLE petitions
    (
        id serial PRIMARY KEY UNIQUE,
        name character varying,
        author_email character varying,
        body character varying
    );

# Миграция базы
 Env:
    
    TRY_MIGRATE = TRUE

# Кодогенерация

    .\openAPI\go-swagger.exe generate server -C .\templates\codegen-config.yaml --template-dir .\templates -f .\openAPI\swagger.yaml -P models.Principal
