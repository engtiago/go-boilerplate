# 🚀 Go-Boilerplate

Go-Boilerplate is a robust base repository for back-end applications in Go. This project provides an initial structure for the development of robust web applications, focusing on authentication and user management.

## 📚 Description

The project is organized into several directories, each with a specific purpose in the application lifecycle:

- 🎛️ `controllers`: This directory contains the business logic for user authentication operations (loginController.go) and user management (usersController.go). The functions in these files deal with obtaining, creating, updating, and deleting users.

- 🚀 `initializers`: This directory is responsible for initializing the components necessary for the application to function.

- 🛡️ `middleware`: This directory manages the middlewares used in the application.

- 📄 `models`: This directory contains the user model (userModel.go), which defines the structure of a 'User' in the database with fields for 'Name', 'Email', and 'Password'.

- 🌐 `routers`: This directory manages the application's routes.

## 📦 Dependencies

The project uses several high-level Go libraries:

- `github.com/gofiber/fiber/v2`: For creating an efficient and flexible web application.

- `github.com/golang-jwt/jwt/v5`: Used for generating JWT tokens for user authentication.

- `golang.org/x/crypto/bcrypt`: Used for password hashing when creating new users.

- `gorm.io/gorm`: Used for interacting with the database.

Additionally, the project also includes Dockerfile and Makefile files, indicating that the project can be containerized for deployment and uses `make` to automate common tasks.

## 🎉 Conclusion

Go-Boilerplate is a solid foundation for any developer who wants to start a new project in Go. With a focus on authentication and user management, this project simplifies the initiation of new applications, allowing developers to focus on building unique features for their applications. 🥳
