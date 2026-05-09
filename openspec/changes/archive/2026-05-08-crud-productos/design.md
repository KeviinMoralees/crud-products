## Context

Proyecto nuevo en Go que implementa una API REST para gestión de productos. Se parte de una estructura de carpetas ya definida que sigue Clean Architecture: dominio, repositorio (interfaz + implementación SQL) y transporte HTTP.

## Goals / Non-Goals

**Goals:**
- API REST funcional con las 5 operaciones CRUD sobre productos.
- Separación clara entre dominio, infraestructura y transporte.
- Wiring manual en `main.go` sin frameworks de inyección de dependencias.

**Non-Goals:**
- Autenticación o autorización.
- Paginación ni filtros avanzados.
- Migraciones automáticas de base de datos.

## Decisions

### Gin como framework HTTP
**Decisión:** usar `github.com/gin-gonic/gin`.
**Alternativas consideradas:** `net/http` estándar (más verboso, sin router paramétrico integrado), `chi` (más ligero pero menor ecosistema).
**Razón:** Gin tiene sintaxis familiar, manejo de parámetros de ruta directo (`c.Param`), binding de JSON integrado (`c.ShouldBindJSON`) y un ecosistema amplio para crecer la app.

### Handler que encapsula el router y las rutas
**Decisión:** en `pkg/services/rest/routes/routes.go` se define un `Handler` struct que contiene el `*gin.Engine` y el `Service`. Su constructor registra todas las rutas. `main.go` solo llama `handler.Run(":8080")`.
**Razón:** Mantiene `main.go` limpio y evita que la configuración de rutas se disperse.

### SQLite con `database/sql` + `mattn/go-sqlite3`
**Decisión:** SQLite como motor de base de datos, usando el driver CGO `mattn/go-sqlite3`.
**Alternativas consideradas:** Postgres (requiere servidor externo), `modernc.org/sqlite` (sin CGO, más simple de compilar).
**Razón:** SQLite es ideal para desarrollo local y aprendizaje: cero configuración, un solo archivo, suficiente para el alcance de este proyecto.

### Cada operación del servicio en su propio archivo
**Decisión:** `create_product.go`, `get_all_products.go`, `get_product_by_id.go`, `update_product.go`, `delete_product.go`.
**Razón:** La estructura ya está definida en el repo. Escala mejor que un único archivo de servicio: cada colaborador puede trabajar en una operación sin conflictos de merge.

### Inyección de dependencias manual
**Decisión:** sin frameworks (wire, dig, fx). Todo se construye en `main.go`.
**Razón:** Proyecto de aprendizaje — ver explícitamente cómo fluye la dependencia de abajo hacia arriba (repo → servicio → handler) es más didáctico.

## Risks / Trade-offs

- **CGO requerido para sqlite3** → Hace más lenta la compilación y requiere gcc. Mitigación: documentar el requisito en el README.
- **Sin migraciones** → La tabla de productos debe crearse manualmente o con un script simple en `main.go`. Mitigación: crear la tabla con `CREATE TABLE IF NOT EXISTS` al arrancar.
- **Sin validaciones avanzadas** → Los campos del producto no tienen validaciones de longitud ni formato. Mitigación: aceptable para el alcance actual; se puede añadir con struct tags de Gin (`binding:"required"`).
