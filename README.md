# 🚀 Go-Boilerplate

Go-Boilerplate é um robusto repositório base para aplicações back-end em Go. Este projeto fornece uma estrutura inicial para o desenvolvimento de aplicações web robustas, focando na autenticação e gerenciamento de usuários.

## 📚 Descrição

O projeto é organizado em vários diretórios, cada um com um propósito específico no ciclo de vida da aplicação:

- 🎛️ `controllers`: Este diretório contém a lógica de negócios para operações de autenticação de usuário (loginController.go) e gerenciamento de usuários (usersController.go). As funções nesses arquivos lidam com a obtenção, criação, atualização e exclusão de usuários.

- 🚀 `initializers`: Este diretório é responsável por inicializar os componentes necessários para o funcionamento da aplicação.

- 🛡️ `middleware`: Este diretório gerencia os middlewares usados ​​na aplicação.

- 📄 `models`: Este diretório contém o modelo de usuário (userModel.go), que define a estrutura de um 'User' no banco de dados com campos para 'Name', 'Email' e 'Password'.

- 🌐 `routers`: Este diretório gerencia as rotas da aplicação.

## 📦 Dependências

O projeto utiliza várias bibliotecas Go de alto nível:

- `github.com/gofiber/fiber/v2`: Para a criação de uma aplicação web eficiente e flexível.

- `github.com/golang-jwt/jwt/v5`: Usada para a geração de tokens JWT para autenticação de usuário.

- `golang.org/x/crypto/bcrypt`: Usada para o hashing de senhas ao criar novos usuários.

- `gorm.io/gorm`: Usada para a interação com o banco de dados.

Além disso, o projeto também inclui arquivos Dockerfile e Makefile, indicando que o projeto pode ser contêinerizado para implantação e utiliza o `make` para automatizar tarefas comuns.

## 🎉 Conclusão

Go-Boilerplate é uma base sólida para qualquer desenvolvedor que deseja iniciar um novo projeto em Go. Com foco em autenticação e gerenciamento de usuários, este projeto simplifica a inicialização de novas aplicações, permitindo que os desenvolvedores se concentrem em construir funcionalidades únicas para suas aplicações. 🥳
