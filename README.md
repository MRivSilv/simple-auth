# simple-auth

API de autenticación centralizada tipo "FusioAuth casero". Emite y valida tokens JWT asimétricos (RS256) para que múltiples proyectos compartan un mismo sistema de usuarios.

## Arquitectura

```
                    ┌──────────────┐
                    │  FusioAuth   │
                    │ (private.pem)│
                    └──────┬───────┘
                           │ login
                           ▼
  ┌──────────────┐  JWT   ┌──────────────┐
  │   App A      │◄───────│   Usuario    │
  │ (public.pem) │        └──────────────┘
  └──────────────┘
  ┌──────────────┐
  │   App B      │
  │ (public.pem) │
  └──────────────┘
```

- **FusioAuth** tiene `private.pem` + `public.pem` → firma tokens
- **Cada app** tiene solo `public.pem` → verifica tokens sin depender de FusioAuth en cada request

## Endpoints

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/register` | Crear usuario |
| POST | `/login` | Iniciar sesión, devuelve JWT |
| GET | `/protected` | Ejemplo de ruta protegida con JWT |

## Uso

### 1. Generar llaves (ya incluidas en `keys/`)

```bash
openssl genpkey -algorithm RSA -out keys/private.pem -pkeyopt rsa_keygen_bits:2048
openssl pkey -in keys/private.pem -pubout -out keys/public.pem
```

### 2. Configurar entorno

```env
JWT_PRIVATE_KEY=keys/private.pem
JWT_PUBLIC_KEY=keys/public.pem
```

### 3. Iniciar

```bash
go run main.go
```

### 4. Probar

```bash
# Registrar
curl -X POST http://localhost:8080/register \
  -d '{"username":"demo","password":"secreta","email":"demo@test.com"}'

# Login
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -d '{"username":"demo","password":"secreta"}' | jq -r .token)

# Ruta protegida
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/protected
```

## Integrar en otra app

Copia `keys/public.pem` en tu proyecto. Usa el mismo código de validación (o el cliente HTTP que prefieras):

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

## Pendiente / Por desarrollar

- [ ] **Base de datos persistente** — reemplazar store en memoria por PostgreSQL, MySQL o SQLite
- [ ] **Refresh tokens** — token de corta duración + refresh token rotable
- [ ] **Scopes / permisos** — claim `scope` en el JWT para autorización granular por proyecto
- [ ] **Roles de usuario** — admin, moderador, etc.
- [ ] **Revocación de tokens** — blacklist en Redis o DB
- [ ] **Email verification** — confirmar cuenta antes de poder loguearse
- [ ] **Recuperación de contraseña** — reset por email
- [ ] **Rate limiting** — evitar bruteforce en login/register
- [ ] **CORS** — permitir requests desde múltiples dominios (esencial para multi-proyecto)
- [ ] **Validación de entrada** — sanitizar email, username, password strength
- [ ] **Manejo de errores consistente** — códigos HTTP y mensajes uniformes
- [ ] **Logging estructurado** — en lugar de log.Println
- [ ] **Health check** — `GET /health`
- [ ] **Graceful shutdown** — capturar SIGTERM/SIGINT
- [ ] **HTTPS** — certificados TLS
- [ ] **Pruebas** — tests unitarios y de integración
- [ ] **Dockerfile / docker-compose** — para despliegue rápido
- [ ] **API keys para servicios** — autenticación máquina-a-máquina
