# Teste Back-end

## Descrição do Projeto

O Teste Back-end é um projeto de gerenciamento de produtos que permite criar, ler, atualizar e excluir produtos. Ele é composto por duas APIs (Go e Ruby on Rails), um sistema de mensagens (Kafka) e dois bancos de dados (SQLite e MongoDB).

## Pré-requisitos

- Docker
- Git

## Como Iniciar o Projeto

1. Clone o repositório no seu ambiente local:
    ```sh
    git clone <URL_DO_REPOSITORIO>
    cd <NOME_DO_REPOSITORIO>
    ```

2. Construa as imagens do Docker:
    ```sh
    docker-compose build
    ```

3. Inicie os containers:
    ```sh
    docker-compose up
    ```

Isso inicializará os seguintes containers:

- `ms_rails_app`: API Rails
- `ms_go_app`: API Go
- `mongo`: Banco de dados MongoDB
- `mongo-express-1`: Interface para MongoDB
- `kafka`: Sistema de mensagens Kafka
- `kafdrop-1`: Interface para Kafka
- `zookeeper-1`: Sistema de Servidor


![image](https://github.com/lipiw/teste-backend-2024/assets/47393970/d29c7f98-bb1c-4488-97a0-3f16aa9f315a)



## Como Testar

1. Utilize o Postman ou Insomnia.
2. Importe o arquivo `Insomnia_teste_backend.json`.

Serão criados dois projetos (ms-rails e ms-go), ambos com as seguintes funcionalidades:

- `index`: Listar produtos
- `show`: Exibir um produto
- `create`: Criar um produto
- `update`: Atualizar um produto
  

![image](https://github.com/lipiw/teste-backend-2024/assets/47393970/37a453da-9a1b-4569-9cdb-fa316e100af4)
