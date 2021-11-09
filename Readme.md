# Planets API

O Projeto criado para controle cadastros de planetas.

## Objetivo

Realizar o cadastros de planetas e consultar a quantidade de vezes que o mesmo aparece no filme Star Wars.

## Execução

#### Variáveis de ambiente(.env):

```
PORT=5000

SWAPI_BASE_URL=https://swapi.dev/api

MONGO_URI=mongodb://localhost:27017
MONGO_DATABASE=planetsDB
```
#### Para iniciar:
```bash
go run .\main.go
```
## Exemplos API

Exemplos de Requisições para a API

## Listar todos Planetas

### Request:

`GET /planets/`

    curl -i -H 'Accept: application/json' http://localhost:5000/planets/

### Response:

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    [
        {
            "id": "6189b9011f3c71cea77710e8",
            "name": "Tatooine",
            "climate": "Arid",
            "terrain": "Desert",
            "viewed_quantity": 0
        }
    ]

## Busca por nome

### Request:

`GET /planets/?search=Tatooine`

    curl -i -H 'Accept: application/json' \
        http://localhost:5000/planets?search=Tatooine

### Response:

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 09 Nov 2021 13:53:56 GMT
    Content-Length: 21

    [
        {
            "id": "6189b9011f3c71cea77710e8",
            "name": "Tatooine",
            "climate": "Arid",
            "terrain": "Desert",
            "viewed_quantity": 0
        }
    ]

## Criar

### Request:

`POST /planets/`

    curl -i -H 'Content-Type: application/json' \
        -X POST \
        -d '{"name": "Tatooine","Climate": "Arid","Terrain": "Desert"}' \
        http://localhost:5000/planets/

### Response:

    HTTP/1.1 201 Created
    Content-Type: application/json
    Date: Tue, 09 Nov 2021 13:53:56 GMT
    Content-Length: 21

    {
        "id": "618a7c0cfc27011a4ae8a1e9",
        "name": "Tatooine",
        "climate": "Arid",
        "terrain": "Desert",
        "viewed_quantity": 5
    }

## Atualizar 

### Request:

`PUT /planets/`

    curl -i -H 'Content-Type: application/json' \
        -X PUT \
        -d '{"id": "618a7c0cfc27011a4ae8a1e9", "name": "Tatooine","Climate": "Arid","Terrain": "Desert"}' \
        http://localhost:5000/planets/618a7c0cfc27011a4ae8a1e9

### Response:

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 09 Nov 2021 13:53:56 GMT
    Content-Length: 21

## Deletar

### Request:

`DELETE /planets/`

    curl -i -H 'Content-Type: application/json' \
        -X DELETE \
        http://localhost:5000/planets/618a7c0cfc27011a4ae8a1e9

### Response:

    HTTP/1.1 200 OK
    Content-Type: application/json
    Date: Tue, 09 Nov 2021 13:53:56 GMT
    Content-Length: 21

## Pré-requisitos

* Golang versão 1.17 ou superior


## Bibliotecas 

* [Gorilla/MUX](https://github.com/gorilla/mux)
* [Godotenv](github.com/joho/godotenv)
* [Mongo Driver](go.mongodb.org/mongo-driver)

## Pontos de melhoria

* Colocar o consumo da SWAPI como Assincrona
* Incrementar mais Testes(Repositorios, Modelos e etc...)

## Contribuidor

* [Altair Moura](https://github.com/alltairr)
