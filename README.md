Please choose your language:

- [English](#jwt-auth-server)
- [Portuguese](#servidor-de-autentica%C3%A7%C3%A3o-jwt)

# JWT Auth Server

This is a simple authentication server in Golang using JWT (JSON Web Tokens) for user authentication. The project is made up of different routes and handlers for signing in, welcoming, refreshing token, and logging out.

## Methods

### `main()`

This is the entry point for the server. It sets up the log format, loads environment variables from the .env file, initializes a router, sets up routes, and starts the server.

### `Welcome(w http.ResponseWriter, r *http.Request)`

This handler is responsible for welcoming the user. It fetches a token from the cookie, verifies the token's validity, extracts the claims, and returns a welcome message to the user.

### `Signin(w http.ResponseWriter, r *http.Request)`

This handler is responsible for logging the user in. It reads credentials from the request body, validates these credentials, creates a JWT token, and returns it in a cookie.

### `Refresh(w http.ResponseWriter, r *http.Request)`

This handler is responsible for refreshing the user's JWT token. It fetches the token from the cookie, verifies its validity, creates a new token, and returns it in a cookie.

### `Logout(w http.ResponseWriter, r *http.Request)`

This handler is responsible for logging the user out. It simply removes the cookie, effectively invalidating the user's JWT token.

### `AuthenticationMiddleware(next http.Handler) http.Handler`

This middleware is responsible for verifying the user's authentication. It's used on routes that require authentication.

## Installation Instructions

1. First, clone the repository with `git clone`.
2. Then, at the root of the project, create a `.env` file with the secret key for the JWT like so: `JWT_KEY=yoursecretkey`.
3. Run the command `go build` to compile the code.
4. Finally, run the command `./yourfilename` to start the server.

## Testing Instructions

1. Use a tool like Postman or cURL to make a POST request to `localhost:8080/signin` with a JSON body containing `username` and `password`. You should receive a cookie with a JWT token.

```json
{
	"username": "admin",
	"password": "123456"
}
```

2. Make a GET request to `localhost:8080/welcome` with the JWT token you received. You should receive a welcome message with the username.

3. Make a POST request to `localhost:8080/refresh` with the JWT token to receive a new token.

4. Make a GET request to `localhost:8080/logout` to invalidate the token.

#### The following packages are used in this application:

* "net/http": This is a native Go package that provides HTTP client and server implementations.

* "github.com/golang-jwt/jwt/v4": This package provides the implementation of JWTs (JSON Web Tokens), which are used for user authentication.

* "github.com/gorilla/mux": This package is a powerful URL router and dispatcher. It's used to route incoming HTTP requests to their corresponding handler functions.

* "github.com/joho/godotenv": This package is used to read environment variables from a .env file, which is a common method for configuration in development environments.

* "encoding/json": This is a native Go package that's used for encoding and decoding JSON.

* "os" and "time": These are native Go packages used for operating system functionality and time manipulation respectively.

* "github.com/sirupsen/logrus": This package is a flexible logging library. It's used for logging debug information, errors, and general server events.

# Servidor de Autenticação JWT

Este é um servidor de autenticação simples em Golang que utiliza JWT (JSON Web Tokens) para autenticação de usuários. O projeto é composto por diferentes rotas e manipuladores para login, boas-vindas, atualização de token e logout.

## Métodos

### `main()`

Este é o ponto de entrada para o servidor. Configura o formato de log, carrega variáveis de ambiente do arquivo `.env`, inicializa um roteador, define rotas e inicia o servidor.

### `Welcome(w http.ResponseWriter, r *http.Request)`

Este manipulador é responsável por acolher o usuário. Ele recebe um token do cookie, verifica a validade do token, extrai as alegações e retorna uma mensagem de boas-vindas ao usuário.

### `Signin(w http.ResponseWriter, r *http.Request)`

Este manipulador é responsável por fazer login no usuário. Ele lê as credenciais do corpo da requisição, valida essas credenciais, cria um token JWT e o retorna no cookie.

### `Refresh(w http.ResponseWriter, r *http.Request)`

Este manipulador é responsável por atualizar o token JWT do usuário. Ele extrai o token do cookie, verifica a validade, cria um novo token e o retorna no cookie.

### `Logout(w http.ResponseWriter, r *http.Request)`

Este manipulador é responsável por fazer o logout do usuário. Ele simplesmente remove o cookie, efetivamente invalidando o token JWT do usuário.

### `AuthenticationMiddleware(next http.Handler) http.Handler`

Este middleware é responsável por verificar a autenticação do usuário. Ele é usado nas rotas que requerem autenticação.

## Instruções de instalação

1. Primeiramente, clone o repositório com `git clone`.
2. Depois, na raiz do projeto, crie um arquivo `.env` com a chave secreta para o JWT, assim: `JWT_KEY=suachavesecreta`.
3. Rode o comando `go build` para compilar o código.
4. Por fim, rode o comando `./nomedoseuarquivo` para iniciar o servidor.

## Instruções para testes

1. Use uma ferramenta como Postman ou cURL para fazer uma requisição POST para `localhost:8080/signin` com um corpo JSON contendo `username` e `password`. Você deve receber um cookie com um token JWT.

```json
{
	"username": "admin",
	"password": "123456"
}
```

2. Faça uma requisição GET para `localhost:8080/welcome` com o token JWT que você recebeu. Você deve receber uma mensagem de boas-vindas com o nome de usuário.

3. Faça uma requisição POST para `localhost:8080/refresh` com o token JWT para receber um novo token.

4. Faça uma requisição GET para `localhost:8080/logout` para invalidar o token.

### Os seguintes pacotes são usados nesta aplicação:

* "net/http": Este é um pacote nativo do Go que fornece implementações de cliente e servidor HTTP.

* "github.com/golang-jwt/jwt/v4": Este pacote fornece a implementação dos JWTs (JSON Web Tokens), que são usados para autenticação do usuário.

* "github.com/gorilla/mux": Este pacote é um poderoso roteador e despachante de URL. É usado para direcionar as solicitações HTTP recebidas para suas respectivas funções de manipulador.

* "github.com/joho/godotenv": Este pacote é usado para ler variáveis de ambiente de um arquivo .env, que é um método comum para configuração em ambientes de desenvolvimento.

* "encoding/json": Este é um pacote nativo do Go que é usado para codificar e decodificar JSON.

* "os" e "time": Estes são pacotes nativos do Go usados para funcionalidade do sistema operacional e manipulação de tempo, respectivamente.

* "github.com/sirupsen/logrus": Este pacote é uma biblioteca de log flexível. É usado para registrar informações de debug, erros e eventos gerais do servidor.