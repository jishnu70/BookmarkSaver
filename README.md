# 📚 Bookmark Saver API

A minimal, secure, and extensible backend API for saving and managing bookmarks with support for user authentication, tagging, and protected endpoints. Built with **Go**, **Gin**, **PostgreSQL**, and **GORM**.

---

## 🚀 Features

- 🔐 JWT-based authentication
- 📝 CRUD operations for bookmarks
- 🏷️ Tag support with many-to-many relationships
- 🧑‍💼 User registration and login
- 🔒 Protected routes with middleware
- 📦 Environment-based configuration

---

## 📁 Project Structure

```
bookmarksaver/
├── controllers/        # Route handler functions
├── initializers/       # Environment and DB setup
├── middleware/         # Auth middleware
├── models/             # GORM models
├── main.go             # Route definitions and server start
└── .env                # Secret keys and DB URL
```

---

## ⚙️ Tech Stack

- **Go**
- **Gin** (HTTP Router)
- **GORM** (ORM)
- **PostgreSQL**
- **JWT** for authentication
- **bcrypt** for password hashing

---

## 📦 Setup & Run Locally

### 1. Clone the repository

```bash
git clone https://github.com/jishnu70/BookmarkSaver.git
cd BookmarkSaver
```

### 2. Create `.env` file

```env
SECRET=your_jwt_secret
DB_URL=postgres://username:password@localhost:5432/bookmarkdb
```

> Make sure PostgreSQL is running and the DB (`bookmarkdb`) exists.

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run migrations

```bash
go run migrate.go
```

### 5. Start the server

```bash
go run main.go
```

> Server will run on `http://localhost:8080`

---

## 🛠️ API Endpoints

### 🔓 Public
- `POST /api/register` – Register new user
- `POST /api/login` – Login and receive JWT token

### 🔐 Protected (JWT required in `Authorization: Bearer <token>` header)
- `GET /auth/bookmarks/` – Get all bookmarks
- `GET /auth/bookmarks/:id` – Get a specific bookmark
- `POST /auth/bookmarks/` – Create a new bookmark
- `PUT /auth/bookmarks/:id` – Update a bookmark
- `DELETE /auth/bookmarks/:id` – Delete a bookmark

---

## 📌 Example Bookmark JSON

```json
{
  "title": "My Portfolio",
  "url": "https://myportfolio.com",
  "tags": ["personal", "work"]
}
```

---

## 🏁 Future Improvements

- 🔄 Refresh tokens
- 🔍 Search/filter by tags
- 📄 Swagger docs
- 🚀 Docker deployment
