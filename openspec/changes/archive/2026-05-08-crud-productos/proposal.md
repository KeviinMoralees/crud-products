## Why

El proyecto necesita una API REST en Go para gestionar un catálogo de productos. Es la funcionalidad base de la aplicación: sin ella no hay nada sobre qué construir.

## What Changes

- Definición del struct `Product` en la capa de dominio (entidad).
- Interfaz `Repository` en el dominio (contrato de acceso a datos).
- Interfaz `Service` + struct con lógica de negocio, con cada operación en su propio archivo.
- Implementación SQL del repositorio en la capa de infraestructura.
- Handler HTTP con Gin: el handler encapsula el router y registra todas las rutas internamente.
- `cmd/main.go` inicializa el handler (que a su vez configura las rutas) y arranca el servidor.

## Capabilities

### New Capabilities

- `product-management`: Ciclo de vida completo de un producto — crear, listar, obtener por ID, actualizar y eliminar — siguiendo Clean Architecture (dominio → repositorio → transporte HTTP con Gin).

### Modified Capabilities

## Impact

- Proyecto nuevo: no hay código existente afectado.
- Se añaden dependencias: `github.com/gin-gonic/gin` como router HTTP y un driver SQL para persistencia.
- `cmd/main.go` es el único punto de wiring: crea el repositorio, lo inyecta en el servicio, crea el handler (que recibe el servicio y registra las rutas), y llama a `handler.Run()`.
- La API quedará expuesta en `/api/productos` y será la base de toda funcionalidad futura.
