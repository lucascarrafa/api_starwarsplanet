# DesafioB2W Api Rest

## Descrição
     
 Criar um jogo com algumas informações da franquia. Para possibilitar  a equipe de front criar essa aplicação, queremos desenvolver uma API que contenha os dados dos planetas.


## Tecnologias utilizadas

- Golang
- MongoDB Atlas
- k6
- Postman
- Visual Studio Code


## Informações para utilizar a API
- Para utilizar a API é necessário configurar o servidor do MongoDB Atlas


## Funcionalidades desenvolvidas
- Adicionar um planeta (com nome, clima e terreno)
(ao adicinar um planeta a API se comunica com API do Star Wars (https://swapi.dev/), obtem a quantidade de aparições e finaliza o processo de salvar)
- Listar planetas
- Buscar por nome
- Buscar por ID
- Remover planeta

## Teste de carga

- Para rodar o teste de carga:

```
$ docker run -i loadimpact/k6 run - <teste_k6.js
```

## API Rest

- Adicionar um planeta: [POST] http://localhost:3333/add

*Body:*

*{ "nome": "Endor",* 
*"clima": "temperate",* 
*"terreno": "forests, mountains, lakes" }*

- Listar os planetas  : [GET] http://localhost:3333/lista

- Buscar por nome : [GET] http://localhost:3333/busca/{nome}

- Buscar por ID: [GET] http://localhost:3333/buscaID/{id}

- Remover planeta: [DELETE] http://localhost:3333/del/{nome}

