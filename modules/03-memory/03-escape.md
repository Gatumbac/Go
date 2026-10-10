# 03 · Escape analysis

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Stack y heap](02-stack-heap.md) · Siguiente: [Slices y backing arrays](04-slices.md)

## Objetivos de aprendizaje

- Explicar escape analysis y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Stack y heap](02-stack-heap.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Los diagnósticos del compilador explican por qué un diseño puede aumentar asignaciones, pero deben leerse con contexto.

## Explicación conceptual

Escape analysis estudia si referencias a un objeto pueden sobrevivir a su scope de asignación o necesitar otra ubicación. Cuando el compilador no puede demostrar que es seguro mantenerlo en stack, puede moverlo al heap. La decisión preserva semántica, no es un error de programa.

`go build -gcflags=-m` muestra decisiones de inlining y escape; `-m=2` añade detalle. Usa flags del package del ejemplo para evitar ruido de toda la biblioteca estándar. Los mensajes dependen de versión y contexto. Una función puede cambiar de comportamiento de asignación al inlined en su caller.

Introducir una interfaz, closure o llamada a fmt puede modificar el análisis observado. Por eso un microexperimento que imprime una variable no mide automáticamente la versión real que no imprime. Distingue escape de una referencia, asignación efectiva y retención: no son sinónimos.

Comprueba con benchmarks `-benchmem` y perfiles. Un diagnóstico puede decir que algo escapa, pero no que eso domina el tiempo ni que quitarlo mejore el sistema. Evita cambiar una API clara por una oscura solo para reducir una cifra aislada.

## Modelo mental visual

```text
reference flow → compiler proof?
 yes: local storage possible
 no: longer-lived storage may be needed
measure → determine practical cost
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/03-escape/main.go):

```go
package main

import "fmt"

func newLabel() *string { s := "sample"; return &s }
func main()             { fmt.Println(*newLabel()) }
```

Desde la raíz:

```sh
go run ./modules/03-memory/examples/03-escape
go build -o /tmp/go-from-zero-example ./modules/03-memory/examples/03-escape
```

**Resultado e interpretación:** Imprime `sample`. Ejecuta además `go build -gcflags=-m` sobre este directorio y busca diagnósticos; no se fija una salida exacta de esos mensajes.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Verificación específica

Desde la raíz del repo:

```sh
go build -gcflags='-m=2' -o /tmp/go-course-escape ./modules/03-memory/examples/03-escape
```

## Comparación con otros lenguajes

C obliga al programador a controlar vida útil; Go analiza y administra. Eso reduce bugs de dangling pointers pero no elimina costes de asignación.

## Errores comunes

Escapes como fallo de calidad; medir fmt en vez del cálculo; memorizar mensajes de una versión.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea dos funciones con el mismo resultado lógico, una devuelve struct pequeño y otra pointer. Entrada: número 5. Salida: mismo campo 5. Aceptación: recoge -m y separa observación de hipótesis; no exigir un escape concreto.

[Enunciado y espacio de implementación](exercises/03-escape/guided/README.md) · [Pistas progresivas](exercises/03-escape/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Construye un benchmark de esas funciones después del módulo de herramientas; por ahora escribe el protocolo de medición. Entrada: ambas funciones. Salida: qué medir y cómo evitar eliminar el resultado. Aceptación: mismo trabajo semántico en ambas variantes.

[Enunciado y espacio](exercises/03-escape/challenge/README.md) · [Pistas opcionales](exercises/03-escape/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Escape implica cuello de botella?
2. ¿Inlining puede alterar el análisis?

[Respuestas razonadas](self-assessment/03-escape.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Los diagnósticos del compilador explican por qué un diseño puede aumentar asignaciones, pero deben leerse con contexto. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Diagnostics](https://go.dev/doc/diagnostics) · [Compiler](https://pkg.go.dev/cmd/compile)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
