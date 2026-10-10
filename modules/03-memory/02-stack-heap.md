# 02 · Stack y heap

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Valores, referencias y aliasing](01-aliasing.md) · Siguiente: [Escape analysis](03-escape.md)

## Objetivos de aprendizaje

- Explicar stack y heap y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Valores, referencias y aliasing](01-aliasing.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Entender vida útil y coste de almacenamiento ayuda a interpretar perfiles sin atribuir magia a los pointers.

## Explicación conceptual

La stack de una goroutine almacena parte de su estado de ejecución; el heap permite objetos cuya vida útil o uso no encaja en ese almacenamiento. Es un modelo útil, no una regla de source: el compilador puede eliminar una variable, mantenerla en registros o colocarla en stack/heap.

Las stacks de goroutines son gestionadas por el runtime y pueden crecer. No son necesariamente stacks de tamaño fijo de threads del sistema. El scheduler multiplexa goroutines sobre threads; eso tampoco elimina límites de memoria o costes de miles de tareas activas.

El scope léxico y la vida útil no son iguales. Una variable puede dejar de estar en scope y su objeto seguir alcanzable mediante un pointer. Go preserva semántica de acceso válido; escape analysis ayuda a decidir ubicación. Usar new no es una promesa de heap y devolver &local no es un dangling pointer.

Asignar en heap puede aumentar trabajo del GC; stack no significa coste cero. No optimices cambiando APIs de valores a pointers antes de medir: podría causar escapes y complicar propiedad. Busca primero retención accidental y asignaciones grandes. Consulta perfiles con un escenario representativo.

## Modelo mental visual

```text
goroutine call frames → short-lived state
reachable object ─────> lifetime may outlive call
compiler decides storage; source semantics stay valid
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/02-stack-heap/main.go):

```go
package main

import "fmt"

func value() *int { n := 42; return &n }
func main()       { p := value(); fmt.Println(*p) }
```

Desde la raíz:

```sh
go run ./modules/03-memory/examples/02-stack-heap
go build -o /tmp/go-from-zero-example ./modules/03-memory/examples/02-stack-heap
```

**Resultado e interpretación:** Imprime `42`; el acceso es válido aunque value ya retornó. La salida no revela ubicación física.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

En C la dirección de una local puede invalidarse al retornar. En Go no debes inferir ese comportamiento desde una similitud sintáctica.

## Errores comunes

Scope=vida útil; new=heap siempre; stack=sin coste; interpretar una dirección como prueba de ubicación.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe un programa que devuelva un pointer a un struct local `Sample` con N=8 y lea N tras retornar. Entrada: ninguna. Salida: 8. Aceptación: explica validez sin afirmar dónde fue asignado.

[Enunciado y espacio de implementación](exercises/02-stack-heap/guided/README.md) · [Pistas progresivas](exercises/02-stack-heap/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe memory-notes.md comparando dos diseños: devolver valor pequeño o pointer compartido. Entrada: ambos contratos. Salida: semántica, riesgo de mutación y medición propuesta. Restricción: no decidir velocidad sin medición.

[Enunciado y espacio](exercises/02-stack-heap/challenge/README.md) · [Pistas opcionales](exercises/02-stack-heap/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿New garantiza heap?
2. ¿Scope determina cuándo GC libera?

[Respuestas razonadas](self-assessment/02-stack-heap.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Entender vida útil y coste de almacenamiento ayuda a interpretar perfiles sin atribuir magia a los pointers. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[GC guide](https://go.dev/doc/gc-guide) · [FAQ allocation](https://go.dev/doc/faq#stack_or_heap)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
