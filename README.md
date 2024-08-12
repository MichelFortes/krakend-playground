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

### Postman Collection
Você pode encontrar a coleção do Postman no diretório **/postman**.

### Observabilidade

- [Grafana](./apps/grafana/README.md)
- [Alloy](./apps/alloy/README.md)
- [Tempo](./apps/tempo/README.md)
- [Mimir](./apps/mimir/README.md)
- [Loki](./apps/loki/README.md)
- [MinIO](./apps/minio/README.md)

## Antes de começar

### Keycloak

1. Acesse o Keycloak <a href="http://localhost:8090" targert="_blanc">aqui</a>
2. Faça login usando as credenciais [ **admin** - **admin** ].
3. Crie um **REALM** com o nome **veterinary_clinic**.
4. Crie um **client scope** com o nome **pets.read.all** e com a propriedade **Include in token scope** habilitada.
5. Crie um client do tipo **OpenID Connect** com qualquer Client Id, habilitando as propriedades **Client authentication** e **Authorization**.
6. Atribua o client scope **pets.read.all** ao client criado.
7. Anote as credenciais do client para requisitar o access token via client credential flow (veja coleção postman).

### MinIO

1. Acesse o MinIO <a href="http://localhost:9001" targert="_blanc">aqui</a>
2. Faça login usando as credenciais [ **miniorootuser** - **miniorootpass** ].
3. Crie os seguintes **Buckets**:
   - tempo
   - loki
   - mimirblocks
   - mimiralert
   - mimirruler
4. Crie uma **Access Key** para cada um dos backends, inserindo as credenciais nos respectivos serviços no **Docker Compose File**:
   - tempo
   - loki
   - mimir
