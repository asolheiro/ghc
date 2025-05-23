![](https://app.gita.cloud/assets/logo-gita-dark-0a5eed80.png)

# GHC

GHC é uma interface de linha de comando para geração de relatórios de saúde de clusters registrados na plataforma [GITA](https://gita.cloud).

- [GHC](#ghc)
  - [Visão geral](#visão-geral)
  - [Uso da ferramenta](#uso-da-ferramenta)
    - [Clone e criação do binário](#clone-e-criação-do-binário)
  - [Melhorias futuras](#melhorias-futuras)

## 1. Visão geral

GHC utiliza as credenciais do usuário para acessar a API da plataforma, coletar dados e, então, gerar o relatório mostrado no template abaixo:

![](./print-screens/ghc-report.png)

As informações coletadas são para geração de relatórios são:

- Uso de recursos
- Versionamento do Kubernetes
- Número de nodes
- Número de problemas
- Número de alertas de segurança
- Dados sobre Incidentes
- Dados sobre Alertas

## 2. Uso da ferramenta

**Pré-requisitos:**

- Go instalado na versão 1.23.5 ou superior
- Uma conta na plataforma Gita


### 2.1. Clone e criação do binário

Para utilizar o binário, faça o clone do repositório:

```bash
git clone https://gitlab.com/jackexperts/healthcheck-jack-experts/gita-healthcheck-generator
```

Então, construa o binário:

```bash
$ cd gita-healthcheck-generator

$ go build -o ./bin/ghc main.go

$ chmod +x ./bin/ghc
```

Então, mova o arquivo para a pasta desejada e execute a ferramenta.

### 2.2. Autenticação

Independente da operação executada, a ferramenta requisita credenciais para autenticação na plataforma.

A primeira é utilizando variáveis de ambiente ou um arquivo `.env` no mesmo nível do binário. Em toda execução, a ferramenta buscará as variáveis `EMAIL` e `PASSWORD` no ambiente ou, somente senão encontrar, requisitará ao usuário conforme mostrado abaixo:

![auth](./print-screens/ghc-auth.png)

### 2.2 Geração de relatórios

GHC tem duas operações definidas para o usuário:

-   `gen-md`      Gera um arquivo markdown com um relatório simples da plataforma Gita.
-   `gen-pdf`     Gera um PDF a partir dos markdowns criados.
-   `gql-gen`     Gera uma requisição GraphQL POST para criar uma página diretamente na Wiki.Js.
-   `pg-gen`      Gera uma interface para visualização do relatório no terminal a partir do relatório markdown.
-   `prin-hc`     Imprime no terminal um relatório simples da plataforma GITA.
-   `help`        Ajuda para os comandos

## 3. Modelo arquivo .Env

```.env
EMAIL=<GITA_EMAIL>
PASSWORD=<GITA_PASSWORD>
WIKIJS_URL=<GRAPH_QL_API_URL>         # Opcional, apenas para a função GraphQL
WIKIJS_API_TOKEN=<GRAPH_QL_BASE_URL>  # Opcional, apenas para a função GraphQL
```
## 4. Melhorias futuras

- Funcionalidade para gerar relatórios de clusters específicos
- Melhorar as saídas no terminal
- Adicionar no Gita
- Autenticação a partir de token