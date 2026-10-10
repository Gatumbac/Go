# Glosario

[Inicio](README.md)

| Término | Significado y desarrollo |
| --- | --- |
| Zero value | Valor inicial de un tipo; no implica validez de negocio. [Variables, constantes y tipos](modules/01-fundamentals/05-types.md) |
| Module | Unidad de dependencias con go.mod y ruta base. [Packages y modules](modules/01-fundamentals/04-packages.md) |
| Package | Archivos de un directorio que comparten declaraciones. [Packages y modules](modules/01-fundamentals/04-packages.md) |
| Rune | Alias de int32 usado para un code point Unicode, no necesariamente un grapheme. [Strings, bytes y runes](modules/01-fundamentals/09-text.md) |
| Slice | Ventana de elementos con longitud y capacidad, potencialmente compartida. [Arrays, slices y maps](modules/01-fundamentals/08-collections.md) |
| Receiver | Parámetro que asocia un método a un tipo. [Métodos](modules/02-design/03-methods.md) |
| Method set | Conjunto de métodos relevante para llamadas e interfaces. [Interfaces y method sets](modules/02-design/04-interfaces.md) |
| Typed nil | Interfaz no nil cuyo valor dinámico es un pointer nil. [Interfaces y method sets](modules/02-design/04-interfaces.md) |
| Wrapping | Contexto de error que preserva la causa mediante %w. [Errores y wrapping](modules/02-design/06-errors.md) |
| Constraint | Conjunto de tipos y operaciones permitidas para un parámetro de tipo. [Genéricos básicos](modules/02-design/09-generics.md) |
| Aliasing | Dos accesos al mismo almacenamiento. [Valores, referencias y aliasing](modules/03-memory/01-aliasing.md) |
| Ownership | Contrato de quién puede modificar datos y durante qué ciclo de vida. [Valores, referencias y aliasing](modules/03-memory/01-aliasing.md) |
| Escape analysis | Análisis del compilador sobre vida útil y ubicación posible. [Escape analysis](modules/03-memory/03-escape.md) |
| Backing array | Almacenamiento de elementos al que accede un slice. [Slices y backing arrays](modules/03-memory/04-slices.md) |
| Heap | Almacenamiento administrado para objetos según decisiones de compilador/runtime. [Stack y heap](modules/03-memory/02-stack-heap.md) |
| GC | Recuperación de memoria no alcanzable; no cierre determinista de recursos. [Garbage collection](modules/03-memory/06-gc.md) |
| Allocation | Reserva de almacenamiento; su coste debe medirse por escenario. [Asignaciones y profiling](modules/03-memory/07-profiling.md) |
| Coverage | Statements ejecutados por tests, no garantía de corrección. [Coverage y fuzzing](modules/04-tooling/03-coverage-fuzz.md) |
| Fuzzing | Exploración automatizada de entradas para comprobar una propiedad. [Coverage y fuzzing](modules/04-tooling/03-coverage-fuzz.md) |
| Benchmark | Medición controlada de una operación y carga definidas. [Benchmarks](modules/04-tooling/04-benchmarks.md) |
| Goroutine | Actividad de ejecución administrada por el runtime. [Goroutines y scheduler](modules/05-concurrency/02-goroutines.md) |
| Happens-before | Relación de orden y visibilidad establecida por sincronización. [Mutex y RWMutex](modules/05-concurrency/06-mutex.md) |
| Data race | Accesos concurrentes sin sincronización a una ubicación, al menos uno escritura. [Carreras, deadlocks y leaks](modules/05-concurrency/07-failures.md) |
| Race condition | Regla cuyo resultado depende de un interleaving no controlado. [Carreras, deadlocks y leaks](modules/05-concurrency/07-failures.md) |
| Deadlock | Actividades que esperan un progreso que no puede ocurrir. [Carreras, deadlocks y leaks](modules/05-concurrency/07-failures.md) |
| Leak | Trabajo o memoria retenidos más allá de su uso lógico. [Carreras, deadlocks y leaks](modules/05-concurrency/07-failures.md) |
| Backpressure | Limitar admisión para conservar un presupuesto de recursos. [Worker pools y shutdown](modules/05-concurrency/09-workers.md) |
| Context | Propagación de deadline, cancelación y datos de request. [Context y cancelación](modules/05-concurrency/08-context.md) |
| Graceful shutdown | Parar admisión y terminar dentro de una política/plazo definidos. [Worker pools y shutdown](modules/05-concurrency/09-workers.md) |
| Handler | Contrato de atender un request y producir una respuesta HTTP. [Servidores y routing](modules/06-backend/02-servers.md) |
| Middleware | Composición alrededor de handlers para políticas compartidas. [Middleware y errores](modules/06-backend/05-middleware.md) |
| Pool SQL | Conjunto administrado de conexiones de database/sql. [Database/sql](modules/06-backend/07-sql.md) |
| Idempotencia | Repetir una operación produce el mismo efecto definido; importante para retries. [TCP y HTTP](modules/06-backend/01-protocols.md) |
| Pointer | Valor que identifica una variable; copiarlo no copia el objeto. [Desarrollo](modules/02-design/02-pointers.md) |
| Interface | Contrato de métodos satisfecho implícitamente por un method set. [Desarrollo](modules/02-design/04-interfaces.md) |
| Struct | Valor compuesto de campos tipados; su copia no implica copia profunda. [Desarrollo](modules/02-design/01-structs.md) |
| Map | Asociación mutable sin orden de iteración garantizado. [Desarrollo](modules/03-memory/05-maps.md) |
| Closure | Función que usa variables de su entorno léxico. [Desarrollo](modules/01-fundamentals/07-functions.md) |
| Runtime | Servicios que acompañan la ejecución, como scheduler y administración de memoria. [Desarrollo](modules/01-fundamentals/02-compilation.md) |
| Stack | Almacenamiento de estado de ejecución administrado para una goroutine. [Desarrollo](modules/03-memory/02-stack-heap.md) |
| Channel | Comunicación tipada que puede bloquear y establecer sincronización. [Desarrollo](modules/05-concurrency/03-channels.md) |
| Mutex | Exclusión mutua para proteger accesos e invariantes compartidas. [Desarrollo](modules/05-concurrency/06-mutex.md) |
| WaitGroup | Coordinación para esperar terminación de tareas, sin proteger sus datos. [Desarrollo](modules/05-concurrency/05-waitgroups.md) |
| Deadline | Instante límite para una operación, propagado mediante context. [Desarrollo](modules/05-concurrency/08-context.md) |
| Backoff | Espera controlada entre intentos para no reintentar de forma agresiva. [Desarrollo](final-project/05-resilience.md) |
| Snapshot | Representación consistente e independiente del estado observado. [Desarrollo](final-project/README.md) |
| DTO | Tipo de datos para el contrato de transporte, distinto de reglas de dominio. [Desarrollo](modules/06-backend/04-json.md) |
| TCP | Stream ordenado de bytes que no preserva mensajes de aplicación. [Desarrollo](modules/06-backend/01-protocols.md) |
| RoundTripper | Capacidad que realiza una transacción HTTP para un Client. [Desarrollo](modules/06-backend/03-clients.md) |
| Transaction | Grupo de operaciones sobre tx que termina mediante commit o rollback. [Desarrollo](modules/06-backend/07-sql.md) |
