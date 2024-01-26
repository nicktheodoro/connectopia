# Connectopia API

Connectopia API is a simple social media application, built in Go, that provides basic functionality similar to Twitter. It allows users to register, log in, create and manage publications, follow and unfollow other users, and update their profile information.
The API is designed to demonstrate the implementation of essential features in a social media platform.

## Features

- **User Management:**
  - Register a new user
  - Log in and obtain authentication tokens
  - Retrieve user information by ID, name, or username
  - Update user profile and password
  - Delete user account

- **Publication Management:**
  - Create new publications
  - Update and delete own publications
  - View publications in user's feed
  - Retrieve a single publication by ID
  - View all publications of a specific user
  - Like and unlike publications

- **Follow/Unfollow:**
  - Follow and unfollow other users
  - Retrieve followers and following lists

## API Endpoints

- **Users:**
  - `POST /users` - Create a new user
  - `POST /login` - Log in and obtain authentication token
  - `GET /users/{id}` - Retrieve user information by ID
  - `GET /users/search` - Search users by name or username
  - `PUT /users/{id}` - Update user profile
  - `PUT /users/{id}/password` - Update user password
  - `DELETE /users/{id}` - Delete user account
  - `GET /users/{id}/followers` - Retrieve followers of a user
  - `GET /users/{id}/following` - Retrieve users followed by a user

- **Publications:**
  - `POST /publications` - Create a new publication
  - `PUT /publications/{id}` - Update a publication
  - `DELETE /publications/{id}` - Delete a publication
  - `GET /publications` - Retrieve publications in the user's feed
  - `GET /publications/{id}` - Retrieve a single publication by ID
  - `GET /users/{id}/publications` - Retrieve all publications of a specific user
  - `POST /publications/{id}/like` - Like a publication
  - `DELETE /publications/{id}/unlike` - Unlike a publication

- **Follow/Unfollow:**
  - `POST /users/{id}/follow` - Follow a user
  - `DELETE /users/{id}/unfollow` - Unfollow a user

## Setup

1. Clone the repository.
2. Run `go get` to install dependencies.
3. Set up the database and environment variables using the provided `.env.example` file.
4. Run the application with `go run main.go`.

## Dependencies

- [github.com/badoux/checkmail](https://github.com/badoux/checkmail) - Email validation
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL database driver
- [github.com/gorilla/mux](https://github.com/gorilla/mux) - HTTP router
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Environment variable loading
- [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go) - JSON Web Token (JWT) implementation

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
