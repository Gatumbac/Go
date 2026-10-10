# 06 · Errores y wrapping

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Composición](05-composition.md) · Siguiente: [Defer, panic y recover](07-defer.md)

## Objetivos de aprendizaje

- Explicar errores y wrapping y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Composición](05-composition.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Los fallos esperables forman parte del contrato: el caller necesita datos para decidir, no solo mensajes impresos.

## Explicación conceptual

`error` es una interfaz con `Error() string`. El patrón `(T, error)` devuelve error nil en éxito y uno no nil en fallo. Debes documentar si T puede usarse cuando hay error: a menudo no, pero existen APIs de lectura con datos parciales.

Usa errors.New para errores simples y fmt.Errorf con `%w` para añadir contexto conservando la causa. `%v` la convierte en texto y pierde la cadena inspeccionable. `errors.Is` busca una causa equivalente en la cadena y `errors.As` extrae un tipo de error; comparar strings es frágil ante contexto o traducciones.

Un sentinel puede representar una condición estable; un error tipado puede transportar campos. Escoge lo mínimo para decisiones reales del caller. No exportes todos los detalles internos accidentalmente al wrappear si tu API no quiere prometer esas causas.

Propaga con contexto útil, o maneja en una frontera. Registrar y devolver repetidamente suele duplicar logs. `panic` no reemplaza errores esperables de validación o red. Un servicio convierte errores a una respuesta externa segura; el mensaje interno no tiene por qué mostrarse a clientes.

## Modelo mental visual

```text
failure → context wrapping → caller classification
               preserves cause        ↓
                                  retry / reject
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/06-errors/main.go):

```go
package main

import (
	"errors"
	"fmt"
)

var errMissing = errors.New("missing data")

func load() error { return fmt.Errorf("load config: %w", errMissing) }
func main()       { err := load(); fmt.Println(errors.Is(err, errMissing)) }
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/06-errors
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/06-errors
```

**Resultado e interpretación:** Imprime `true`. El contexto no destruye la causa.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

En JS/Java las excepciones pueden saltar frames; Go normalmente mantiene el camino del error explícito en los retornos.

## Errores comunes

Comparar mensajes; descartar error con _; loguear en cada capa; panic para entrada inválida.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `ParseAge(text string) (int, error)` usando strconv: dominio 0..130. Entrada: "18","x","-1". Salida: 18,nil; error; error. Restricción: no imprimir ni panic. Aceptación: errores con contexto, valor no usable en fallo.

[Enunciado y espacio de implementación](exercises/06-errors/guided/README.md) · [Pistas progresivas](exercises/06-errors/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Define un sentinel `ErrEmpty` y función `First(values []int) (int,error)`. Entrada: vacío y [9]. Salida: error identificable con errors.Is; 9,nil. Restricción: el caller no compara texto. Aceptación: documentar causa y zero value en fallo.

[Enunciado y espacio](exercises/06-errors/challenge/README.md) · [Pistas opcionales](exercises/06-errors/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Qué aporta %w?
2. ¿Es suficiente loguear un error?

[Respuestas razonadas](self-assessment/06-errors.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Los fallos esperables forman parte del contrato: el caller necesita datos para decidir, no solo mensajes impresos. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Errors](https://pkg.go.dev/errors) · [Wrapping](https://go.dev/blog/go1.13-errors)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
