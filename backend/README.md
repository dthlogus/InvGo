# InvGo API

O projeto InvGo é um software que tem como objetivo coletar informações financeiras e realizar cálculos para fornecer informações relevantes, como PV (Valor Presente), PVP (Preço sobre Valor Patrimonial), DY (Dividend Yield) e outros. Ele foi projetado para ajudar os usuários a entender melhor suas finanças e tomar decisões informadas.

## Status do projeto

:construction: Projeto em fase de contrução e testes :construction

## Funcionalidades do projeto escopo V-1.0

- `Usuario`: CRUD de usuário para que tenha controle sobre os perfies de cada um e trazer dados personalizados

- `Perfil`: Perfil do usuário tem o poder de falar o que ele quer ver como uma oportunidade ou não, e isso será refletido no Front-End

- `Stocks`: Informações que podem ser ou não relevantes para as empresas e fazer o calculo sobre os mesmos.

 ## 🛠️ Como utilizar o projeto

 Antes de rodar o projeto tenha certeza que tenha instalado o [GIT](https://git-scm.com/), [GO](https://go.dev/) e o [Docker](https://docs.docker.com/get-docker/) em sua máquina caso não tinha você pode clicar nos links e será redirecionado para a página de instalação.

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