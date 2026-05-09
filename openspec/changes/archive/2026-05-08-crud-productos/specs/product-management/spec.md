## ADDED Requirements

### Requirement: Crear producto
El sistema SHALL permitir crear un nuevo producto proporcionando nombre, descripción, precio y stock.

#### Scenario: Creación exitosa
- **WHEN** se envía un POST a `/api/productos` con nombre, descripción, precio y stock válidos
- **THEN** el sistema guarda el producto en la base de datos y devuelve el producto creado con su ID asignado y HTTP 201

#### Scenario: Campos requeridos faltantes
- **WHEN** se envía un POST a `/api/productos` sin nombre o sin precio
- **THEN** el sistema devuelve HTTP 400 con un mensaje de error

---

### Requirement: Listar todos los productos
El sistema SHALL devolver la lista completa de productos almacenados.

#### Scenario: Lista con productos
- **WHEN** se envía un GET a `/api/productos`
- **THEN** el sistema devuelve HTTP 200 con un array JSON de todos los productos

#### Scenario: Lista vacía
- **WHEN** se envía un GET a `/api/productos` y no hay productos registrados
- **THEN** el sistema devuelve HTTP 200 con un array vacío `[]`

---

### Requirement: Obtener producto por ID
El sistema SHALL devolver un único producto dado su identificador.

#### Scenario: Producto encontrado
- **WHEN** se envía un GET a `/api/productos/:id` con un ID existente
- **THEN** el sistema devuelve HTTP 200 con el producto correspondiente

#### Scenario: Producto no encontrado
- **WHEN** se envía un GET a `/api/productos/:id` con un ID que no existe
- **THEN** el sistema devuelve HTTP 404 con un mensaje de error

---

### Requirement: Actualizar producto
El sistema SHALL permitir actualizar los campos de un producto existente.

#### Scenario: Actualización exitosa
- **WHEN** se envía un PUT a `/api/productos/:id` con un ID existente y campos válidos
- **THEN** el sistema actualiza el producto y devuelve HTTP 200 con el producto actualizado

#### Scenario: Producto no encontrado al actualizar
- **WHEN** se envía un PUT a `/api/productos/:id` con un ID que no existe
- **THEN** el sistema devuelve HTTP 404 con un mensaje de error

---

### Requirement: Eliminar producto
El sistema SHALL permitir eliminar un producto existente por su ID.

#### Scenario: Eliminación exitosa
- **WHEN** se envía un DELETE a `/api/productos/:id` con un ID existente
- **THEN** el sistema elimina el producto y devuelve HTTP 204 sin cuerpo

#### Scenario: Producto no encontrado al eliminar
- **WHEN** se envía un DELETE a `/api/productos/:id` con un ID que no existe
- **THEN** el sistema devuelve HTTP 404 con un mensaje de error
