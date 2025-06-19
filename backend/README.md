## 🤝 Api-EchoCRUD

Bem-vindo ao projeto da EchoCRUD! Este repositório contém o código-fonte de uma API RESTful, projetada para gerenciar Estabelecimentos e Lojas.

> [!IMPORTANT]
> ⚠️ Foi utilizado o GORM para gerar automaticamente os ID's de Estabelecimentos e Lojas, mas foi perceptível que não é uma boa prática. Os ID's sequenciais são um problema em termos de riscos de segurança e privacidade, criando brechas para IDOR.
>
> 🔄 Testes Unitários
>
> 🔄 SOLID
>
> 🔄 Docker (Frontend + Backend)

## 🎯 Endpoints

### Establishments

|        Endpoint         |               Router               |             Description             |
| :---------------------: | :--------------------------------: | :---------------------------------: |
|    **[`get`](#get)**    |         `/establishments`          |  Buscar todos os estabelecimentos   |
|   **[`post`](#post)**   |          `/establishment`          |    Criar um novo estabelecimento    |
|    **[`get`](#get)**    | `/establishment/{establishmentId}` |    Buscar estabelecimento por ID    |
|    **[`put`](#put)**    | `/establishment/{establishmentId}` | Atualizar estabelecimento existente |
| **[`delete`](#delete)** | `/establishment/{establishmentId}` |   Deletar estabelecimento por ID    |

### Stores

|        Endpoint         |                        Router                        |                 Description                 |
| :---------------------: | :--------------------------------------------------: | :-----------------------------------------: |
|    **[`get`](#get)**    |      `/establishments/{establishmentId}/stores`      | Buscar todas as lojas de um estabelecimento |
|   **[`post`](#post)**   |      `/establishments/{establishmentId}/stores`      |    Criar nova loja em um estabelecimento    |
|    **[`get`](#get)**    | `/establishments/{establishmentId}/stores/{storeId}` |             Buscar loja por ID              |
|    **[`put`](#put)**    | `/establishments/{establishmentId}/stores/{storeId}` |          Atualizar loja existente           |
| **[`delete`](#delete)** | `/establishments/{establishmentId}/stores/{storeId}` |             Deletar loja por ID             |

## 📫 Testing with Postman

To make API testing easier, a Postman collection is included in this repository. You can import it with a single click using the button below.

1. Download the file EchoCRUD_API.collection.json.
2. Open Postman and click File > Import....
3. Select the file you downloaded.
4. The "EchoCRUD API" collection will appear in your list of Postman collections.

## Technologies

### Backend

- [Echo(GO)](https://echo.labstack.com/)
- [gORM](https://gorm.io/)
- [PostgreSQL](https://gorm.io/)

## 🚀 Getting started

- [**Go**](https://go.dev/doc/install) (version 1.24)
- [**Docker**](https://docs.docker.com/engine/install/) **Docker Compose**

1. Clone the project and access the folder

   ```zsh
   $ git clone https://github.com/abraaodev/echocrud.git && cd echocrud
   ```

2. Create .env

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=docker
   DB_PASSWORD=docker
   DB_NAME=echocrud
   ```

3. Install the dependencies

   ```zsh
   go mod tidy
   ```

4. **Execute aplication**

   Unic command for build image API and Database

   ```zsh
   docker-compose up --build -d
   ```

- **A) To populate the database in Execution (Seed):**

  ```zsh
  docker compose exec api /server --seed
  ```

- **B) Stop Aplication:**

  ```zsh
  docker compose down
  ```

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

---

<p align="center">Made with ❤️ by Abraão DEV</p>
