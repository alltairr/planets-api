# Planets API

Projeto realiza o cadastro de planetas.

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
