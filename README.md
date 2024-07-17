# krakend-playground

This project was created to demonstrate the api gateway architectural pattern.

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
6. Anote as credenciais do client para requisitar o access token via client credential flow (veja coleção postman).

### Postman Collection
You can find a postman collection under the /postman folder