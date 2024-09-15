# GoLangExpress: JWT Authentication and Protected Routes

## About This Project

Hey there! 👋 I'm an experienced developer with 14+ years under my belt, primarily working with languages like Python, JavaScript, and Java. Recently, I've been diving into Go, and this project is my playground for learning how to implement authentication and protected routes in a Go web application.

## Why Go?

After years of working with dynamically typed languages and the JVM, I was intrigued by Go's simplicity, strong typing, and excellent performance. Its built-in concurrency features and fast compilation times were the cherries on top that convinced me to give it a shot.

## What I've Learned

This project has been an excellent learning experience. Here are some key takeaways:

1. Go's syntax and structure (coming from other languages, it's refreshingly straightforward)
2. Working with Gorilla Mux for routing (it's like Express.js for Go!)
3. Implementing JWT authentication in Go
4. Structuring a Go project (it's different from what I'm used to, but I'm starting to see the benefits)
5. Go's approach to middleware

## Project Structure

```
goLangExpress/
├── main.go
├── auth/
│   └── auth.go
├── middleware/
│   └── jwt.go
├── routes/
│   └── routes.go
├── handlers/
│   └── handlers.go
└── README.md
```

## How to Run

1. Clone this repository
2. Navigate to the project directory:
   ```
   cd goLangExpress
   ```
3. Run the application:
   ```
   go run main.go
   ```
4. The server will start on `http://localhost:8080`

## API Endpoints

- `POST /login`: Authenticate and receive a JWT
- `POST /logout`: (Simulated) logout
- `GET /me`: Get user info (protected)
- `GET /protected1` to `/protected5`: Sample protected routes

To test protected routes, include the JWT in the Authorization header:
```
Authorization: Bearer <your_jwt_here>
```

## Future Improvements

As I continue learning Go, I plan to enhance this project:

1. Implement proper error handling (I know my current approach is not idiomatic Go)
2. Add unit tests (I'm used to TDD, and I miss it in this project)
3. Use a real database for user management
4. Implement refresh tokens
5. Add logging
6. Containerize the application with Docker

## Disclaimer

I'm new to Go, so this code might not be the most idiomatic or efficient. If you're an experienced Go developer and spot areas for improvement, I'd love to hear your feedback! Learning is a journey, and I'm excited to be on this path with Go.

Happy coding! 🚀