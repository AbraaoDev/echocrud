<h1 align="center">
    <img alt="Capa do Projeto EchoCRUD" title="EchoCRUD" src=".github/capa.png" />
</h1>

## 🤝 App-EchoCRUD

Bem-vindo ao projeto da EchoCRUD! Este repositório contém o código-fonte de uma API RESTful, projetada para gerenciar Estabelecimentos e Lojas.

> [!IMPORTANT]
> ⚠️ Foi utilizado o GORM para gerar automaticamente os ID's de Estabelecimentos e Lojas, mas foi perceptível que não é uma boa prática. Os ID's sequenciais são um problema em termos de riscos de segurança e privacidade, criando brechas para IDOR.
>
> 🔄 Testes Unitários
>
> ✅ SOLID
>
> ✅ Docker (Frontend + Backend)

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

[![Run in Postman](https://run.pstmn.io/button.svg)](https://api.postman.com/collections/38248876-076684d1-ba80-45fd-b45c-37aede86f81b?[chaveDeAcessoAqui])

1. Download the file EchoCRUD_API.collection.json.
2. Open Postman and click File > Import....
3. Select the file you downloaded.
4. The "EchoCRUD API" collection will appear in your list of Postman collections.

## Technologies

### Frontend

- [Nuxt](https://nuxt.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [TanStack Table](https://tanstack.com/table/latest)
- [TanStack Query](https://tanstack.com/query/latest)
- [Zod](https://tanstack.com/query/latest)
- [VueTheMask](https://vuejs-tips.github.io/vue-the-mask/)
- [Shadcn-Vue](https://www.shadcn-vue.com/)
- [Tailwind](https://tailwindcss.com/)

### Backend

- [Echo(GO)](https://echo.labstack.com/)
- [gORM](https://gorm.io/)
- [PostgreSQL](https://gorm.io/)

## 🚀 Getting started

- [**Docker**](https://docs.docker.com/engine/install/) **Docker Compose**

1. Clone the project and access the folder

   ```zsh
   $ git clone https://github.com/abraaodev/echocrud.git && cd echocrud
   ```

2. Create .env

   ```env
   DB_USER=echouser
   DB_PASSWORD=echopass
   DB_NAME=echocrud_db
   DB_HOST=db
   DB_PORT=5432
   API_PORT=8080
   FRONTEND_PORT=3000
   ```

3. **Execute aplication**

   Unic command for build image API and Database

   ```zsh
   docker-compose up --build
   ```

   ```zsh
   docker-compose up
   ```

- **A) Stop Aplication:**

  ```zsh
  docker compose down
  ```

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

---

<p align="center">Made with ❤️ by Abraão DEV</p>
