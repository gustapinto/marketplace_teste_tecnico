# marketplace_teste_tecnico

## Instalando o projeto
1. Clone esse repositório
2. Inicie o ambiente docker com `docker-compose up -d --build`
3. Entre no container Go com `docker-compose exec go bash`
4. Inicie o serviço API com `go run src/main.go`
5. Pronto! Agora é só realizar requisições para a API :D

## Requisições

### Obtendo todos os itens (GET)
```bash
curl localhost:80/api/products
```

### Obtendo item pelo código (GET)
```bash
curl localhost:80/api/products/foo
```

### Adicionando novo item (POST)
```bash
curl -X POST localhost:80/api/products \
    -H 'Content-Type: application/json' \
    -d '{"codigo":"foo","nome":"foobar","preco_de":100.00,"preco_por":70.00,"estoque":{"total":100,"corte":80}}'
```

### Atualizando um item (PUT)
```bash
curl -X PUT localhost:80/api/products/foo \
    -H 'Content-Type: application/json' \
    -d '{"codigo":"foo","nome":"foobar","preco_de":100.00,"preco_por":70.00,"estoque":{"total":150,"corte":95}}'
```

### Deletando um item (DELETE)
```bash
curl -X DELETE localhost:80/api/products/foo
```
