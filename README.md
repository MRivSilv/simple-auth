# simple-auth

Centralized auth API — a homegrown "FusioAuth-like" service. Issues and validates RS256 JWT tokens so multiple projects can share a single user system.

## Architecture

```
                    ┌──────────────┐
                    │  simple-auth │
                    │ (private.pem)│
                    └──────┬───────┘
                           │ login
                           ▼
  ┌──────────────┐  JWT   ┌──────────────┐
  │   App A      │◄───────│   User       │
  │ (public.pem) │        └──────────────┘
  └──────────────┘
  ┌──────────────┐
  │   App B      │
  │ (public.pem) │
  └──────────────┘
```

- **simple-auth** holds `private.pem` + `public.pem` → signs tokens
- **Each app** only gets `public.pem` → validates tokens without calling the auth server on every request

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/register` | Create a user |
| POST | `/login` | Authenticate, returns JWT |
| GET | `/protected` | Example protected route (requires JWT) |

## Usage

### 1. Generate keys (already included in `keys/`)

```bash
openssl genpkey -algorithm RSA -out keys/private.pem -pkeyopt rsa_keygen_bits:2048
openssl pkey -in keys/private.pem -pubout -out keys/public.pem
```

### 2. Configure environment

```env
JWT_PRIVATE_KEY=keys/private.pem
JWT_PUBLIC_KEY=keys/public.pem
```

### 3. Start

```bash
go run main.go
```

### 4. Try it

```bash
# Register
curl -X POST http://localhost:8080/register \
  -d '{"username":"demo","password":"secreta","email":"demo@test.com"}'

# Login
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -d '{"username":"demo","password":"secreta"}' | jq -r .token)

# Protected route
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/protected
```

## Integrating into another app

Copy `keys/public.pem` to your project. Validate tokens locally:

```go
import "github.com/golang-jwt/jwt/v5"

var publicKey *rsa.PublicKey

func init() {
    data, _ := os.ReadFile("public.pem")
    publicKey, _ = jwt.ParseRSAPublicKeyFromPEM(data)
}

func validateToken(tokenStr string) (*jwt.Token, error) {
    return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        return publicKey, nil
    })
}
```

## Roadmap / TODO

- [ ] **Persistent database** — replace in-memory store with PostgreSQL, MySQL or SQLite
- [ ] **Refresh tokens** — short-lived access token + rotatable refresh token
- [ ] **Scopes / permissions** — `scope` claim in JWT for granular per-project authorization
- [ ] **User roles** — admin, moderator, etc.
- [ ] **Token revocation** — blacklist via Redis or database
- [ ] **Email verification** — confirm account before login
- [ ] **Password reset** — email-based recovery
- [ ] **Rate limiting** — prevent brute force on login/register
- [ ] **CORS** — allow requests from multiple domains (essential for multi-project setups)
- [ ] **Input validation** — sanitize email, username, enforce password strength
- [ ] **Consistent error handling** — uniform HTTP codes and error messages
- [ ] **Structured logging** — replace log.Println
- [ ] **Health check** — `GET /health`
- [ ] **Graceful shutdown** — handle SIGTERM/SIGINT
- [ ] **HTTPS** — TLS certificates
- [ ] **Tests** — unit and integration tests
- [ ] **Dockerfile / docker-compose** — easy deployment
- [ ] **API keys for services** — machine-to-machine authentication
