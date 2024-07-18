# krakend-playground

Projeto de demonstração da utilidade de um API Gateway

## Content

### API gateway

- KrakenD

```bash
http://localhost:8080
```

### Upstream Services

- Pets
- Owners
- Doctors

### Authorization Server

- Keycloak
```bash
http://localhost:8090
```

Passos a serem feitos:

1. Acesse o Keycloak
2. Faça login no Keycloak usando o usuário as credenciais [ **admin** - **admin** ].
3. Crie um **REALM** com o nome **veterinary_clinic**.
4. Crie um **client scope** com o name **pets.read.all** com a propriedade **Include in token scope** habilitada.
5. Crie um client do tipo **OpenID Connect** com qualquer Client Id, habilitando as propriedades **Client authentication** e **Authorization**.
6. Atribua o client scope **pets.read.all** ao client criado.
7. Anote as credenciais do client para requisitar o access token via client credential flow (veja coleção postman).

### Postman Collection
Você pode encontrar a coleção do Postman no diretório **/postman**.

### Extraindo Métricas de Traces no Grafana (TraceQL Metrics)

https://github.com/grafana/tempo/blob/main/docs/design-proposals/2023-11%20TraceQL%20Metrics.md

Exemplo de query:
```
{ status = ok } | rate() by (span.http.route)
```