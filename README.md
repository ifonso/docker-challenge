# Docker challenge

A Aplicação precisa de duas variáveis de ambiente:
- MONGO_URI: URI de conexão com o MongoDB
- PORT: Porta em que a aplicação irá rodar

## Docker build
É necessário criar um docker file com duas fases:
1. Na primeira fase, instale as dependências e faça o build da aplicação. (Use qualquer imagem go 1.22)
2. Na segunda fase, copie o binário compilado da primeira fase e rode a aplicação. (Use alpine:3.14)

- Para compilar rodar o comando:
```bash
  $ make build
```
Isso vai criar um binário chamado "app" na raiz do projeto.

Não esqueça de configurar as variáveis de ambiente. Tanto da aplicação quanto do MongoDB.
(É necessário rodar um container do MongoDB)

## Check

Se tudo estiver funcionando, a aplicação vai iniciar sem problemas e vai logar "✅ Connected to mongoDB." e você terá acesso a rota que retorna um json com a mensagem "Hello World!".