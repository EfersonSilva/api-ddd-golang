# GO API 

Api para cadastrar conta e fazer transferencia entre elas.

## Arquitetura

* 3 áreas
    * application
    * framework
    * interfaces


main.go

application
    ├── account.go
    │── login.go
    └── transfer.go

framework
    ├── config
    │   └── db.go
    ├── infrastructure
    │   ├── account_repository.go
    │   └── transfer_repository.go
    └── tokenJWT
        └── jwt.go

Interfaces
    ├── handler.go
    └── respond.go
```

## Endpoints
-- atualmente configurado para: http://localhost:8000

### Auth Login

| Description | http | path |
|:--:|:--:|:--|
| loga | POST | /api/v1/login |


### Account

| Description | http | path |
|:--:|:--:|:--|
| cadastra | POST | /api/v1/account |
| lista | GET | /api/v1/account |
| retorna  | GET | /api/v1/account/:id |
| retorna  | GET | /api/v1/accounts/:account_id/balance |
| atualiza | PUT | /api/V1/account |
| deleta | DELETE | /api/v1/account/:id |


### Transfer

| Description | http | path |
|:--:|:--:|:--|
| cadastra | POST | /api/v1/transfer |
| retorna  | GET | /api/v1/transfer |


## Build & Deploy com Docker

use: docker-compose up --build 
```

