# Fuentes y compatibilidad

[Inicio](../README.md)

Base de ejecución: Go 1.26, verificado con Go 1.26.8 el 2026-10-10. La documentación online puede describir versiones posteriores; comprueba las marcas de versión y `go doc` en tu equipo. Los detalles de runtime no son un contrato del lenguaje.

| Fuente | Uso |
| --- | --- |
| [Especificación](https://go.dev/ref/spec) | Tipos, asignación, evaluación, interfaces y semántica. |
| [Go 1.26](https://go.dev/doc/go1.26) | Cambios del lenguaje y toolchain de la base elegida. |
| [Modules](https://go.dev/ref/mod) y [layout](https://go.dev/doc/modules/layout) | Imports, dependencias y organización. |
| [Effective Go](https://go.dev/doc/effective_go) | Estilo y modelos; complementar con fuentes recientes para genéricos y modules. |
| [Modelo de memoria](https://go.dev/ref/mem) | Sincronización y visibilidad entre goroutines. |
| [GC](https://go.dev/doc/gc-guide) | Alcanzabilidad y costes del recolector. |
| [Diagnostics](https://go.dev/doc/diagnostics) | Profiling, tracing y debugging. |
| [Testing](https://pkg.go.dev/testing) | Tests, fuzzing, benchmarks y subtests. |
| [Race detector](https://go.dev/doc/articles/race_detector) | Detección dinámica y limitaciones. |
| [HTTP](https://pkg.go.dev/net/http) | Routing, clientes, servidores y shutdown. |
| [SQL](https://go.dev/doc/database/) | Pool, consultas, transacciones y cancelación. |

No usar cifras universales para crecimiento de slices, coste de goroutines o GC. En Go 1.26 Green Tea está habilitado por defecto, pero el libro enseña alcanzabilidad y tradeoffs, no un layout dependiente de ese recolector. La salida de escape analysis y profiling depende de versión, flags, optimizaciones y hardware. Benchmarks son mediciones de un escenario, no promesas.
