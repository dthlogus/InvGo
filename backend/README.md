# InvGo API

O projeto InvGo é um software que tem como objetivo coletar informações financeiras e realizar cálculos para fornecer informações relevantes, como PV (Valor Presente), PVP (Preço sobre Valor Patrimonial), DY (Dividend Yield) e outros. Ele foi projetado para ajudar os usuários a entender melhor suas finanças e tomar decisões informadas.

## Indices

- [InvGo API](#invgo-api)
- [Índices](#indices)
- [Status do projeto](#status-do-projeto)
- [Documentaçao](#documentação)
- [Funcionalidades do projeto escopo V-1.0](#funcionalidades-do-projeto-escopo-v-10)
- [Como utilizar o projeto](#%EF%B8%8F-como-utilizar-o-projeto)
- [Técnicas e tecnologias utilizadas](#%EF%B8%8F-técnicas-e-tecnologias-utilizadas)


## Status do projeto

Atualmente V1.0

## Documentação

[Documentaça](https://documenter.getpostman.com/view/15824835/2s946pXoYv)

## Funcionalidades do projeto escopo V-1.0

- `Usuario`: CRUD de usuário para que tenha controle sobre os perfies de cada um e trazer dados personalizados

- `Perfil`: Perfil do usuário tem o poder de falar o que ele quer ver como uma oportunidade ou não, e isso será refletido no Front-End

- `Stocks`: Informações que podem ser ou não relevantes para as empresas e fazer o calculo sobre os mesmos.

 ## 🛠️ Como utilizar o projeto

 Antes de rodar o projeto tenha certeza que tenha instalado o [GIT](https://git-scm.com/), [GO](https://go.dev/) e o [Docker](https://docs.docker.com/get-docker/) em sua máquina caso não tenha você pode clicar nos links e será redirecionado para a página de instalação.

 Não se também de adicionar um arquivo chamado ".env" na raiz do backend e acionar a seguintes váriaveis de ambiente.

 - MONGODB_URL = `URL para acessar o banco de dados Mongo`

Clone o projeto

```bash
  git clone https://github.com/dthlogus/InvGo.git
```

Entre no diretório do projeto

```bash
  cd InvGo/backend
```

Instale as dependências

```bash
  docker compose up -d
```

Inicie o servidor

```bash
  go run main.go
```

## ✔️ Técnicas e tecnologias utilizadas

- ``GO 1.20.6``
- ``Visual Studio Code``
- ``gin-gonic``
- ``Mongodb``
- ``Docker``
- ``Postman``