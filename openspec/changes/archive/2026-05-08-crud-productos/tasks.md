## 1. Dependencias y módulo

- [x] 1.1 Agregar `github.com/gin-gonic/gin` y `github.com/mattn/go-sqlite3` al `go.mod` con `go get`
- [x] 1.2 Verificar que `go mod tidy` deja el módulo limpio

## 2. Entidad de dominio

- [x] 2.1 Definir el struct `Product` en `pkg/domains/entities/product.go` con campos: ID, Nombre, Descripcion, Precio, Stock

## 3. Interfaz del repositorio

- [x] 3.1 Definir la interfaz `Repository` en `pkg/domains/repository/product/product.go` con los métodos: Create, GetAll, GetByID, Update, Delete

## 4. Implementación SQL del repositorio

- [x] 4.1 Completar el struct `repository` en `pkg/repositories/product/repository.go` con el campo `*sql.DB`
- [x] 4.2 Implementar el constructor `NewRepository(db *sql.DB) Repository`
- [x] 4.3 Implementar el método `Create`
- [x] 4.4 Implementar el método `GetAll`
- [x] 4.5 Implementar el método `GetByID`
- [x] 4.6 Implementar el método `Update`
- [x] 4.7 Implementar el método `Delete`

## 5. Servicio de dominio

- [x] 5.1 Definir la interfaz `Service` y el struct `service` en `pkg/domains/product/service.go` con inyección del `Repository`
- [x] 5.2 Implementar `CreateProduct` en `pkg/domains/product/create_product.go`
- [x] 5.3 Implementar `GetAllProducts` en `pkg/domains/product/get_all_products.go`
- [x] 5.4 Implementar `GetProductByID` en `pkg/domains/product/get_product_by_id.go`
- [x] 5.5 Implementar `UpdateProduct` en `pkg/domains/product/update_product.go`
- [x] 5.6 Implementar `DeleteProduct` en `pkg/domains/product/delete_product.go`

## 6. Handler HTTP con Gin

- [x] 6.1 Definir el struct `Handler` en `pkg/services/rest/routes/routes.go` con `*gin.Engine` y `Service`
- [x] 6.2 Implementar el constructor `NewHandler(service product.Service) *Handler` que registra todas las rutas en `/api/productos`
- [x] 6.3 Implementar el método `Run(addr string)` que arranca el servidor Gin
- [x] 6.4 Implementar el handler `CreateProduct`
- [x] 6.5 Implementar el handler `GetAllProducts`
- [x] 6.6 Implementar el handler `GetProductByID`
- [x] 6.7 Implementar el handler `UpdateProduct`
- [x] 6.8 Implementar el handler `DeleteProduct`

## 7. Wiring en main.go

- [x] 7.1 Abrir la conexión SQLite y crear la tabla `productos` con `CREATE TABLE IF NOT EXISTS`
- [x] 7.2 Instanciar el repositorio, el servicio y el handler en orden
- [x] 7.3 Llamar a `handler.Run(":8080")` para arrancar la API
