# 09 · Genéricos básicos

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Organización e idiomas](08-organization.md) · Siguiente: [Valores, referencias y aliasing](../03-memory/01-aliasing.md)

## Objetivos de aprendizaje

- Explicar genéricos básicos y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Organización e idiomas](08-organization.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Reutilizar un algoritmo para varios tipos no debe obligarte a perder comprobaciones usando any.

## Explicación conceptual

Los parámetros de tipo representan tipos elegidos al instanciar una función o tipo genérico. Una constraint define qué tipos admite y qué operaciones se permiten en su cuerpo. `any` admite cualquiera, pero no autoriza operaciones arbitrarias; `comparable` permite igualdad para tipos que satisfacen sus reglas.

Una constraint con `~int` incluye tipos cuyo underlying type es int. Sin ~, una lista int admite ese tipo concreto, no todos los tipos definidos con esa base. Uniones expresan conjuntos; no equivalen a valores runtime. La inferencia puede evitar argumentos de tipo explícitos cuando los parámetros de valores los determinan.

Usa genéricos para un algoritmo realmente común, como transformar una colección o comparar tipos compatibles. Usa interfaces ordinarias cuando el comportamiento es la frontera, como Read o Write. Generics no reemplaza polimorfismo de capacidades ni convierte todos los algoritmos en más eficientes.

No prometas un mecanismo de compilación fijo por instanciación ni ausencia de costes: son detalles del toolchain. Mantén constraints mínimas. Métodos ordinarios no pueden declarar sus propios parámetros de tipo adicionales; el receiver puede ser de un tipo genérico. Comparado con C, no necesitas preprocesador para conservar tipos.

## Modelo mental visual

```text
generic algorithm + constraint
       ↓ choose T
checked operations for T → concrete call
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/09-generics/main.go):

```go
package main

import "fmt"

func same[T comparable](a, b T) bool { return a == b }
func main()                          { fmt.Println(same(3, 3), same("a", "b")) }
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/09-generics
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/09-generics
```

**Resultado e interpretación:** Imprime `true false`. El compilador comprueba ambas instanciaciones.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Los generics TS describen tipos sobre JS emitido; Go ejecuta código tipado compilado. Templates C++ tienen otras reglas; no asumas la misma expansión.

## Errores comunes

Any como permiso para sumar; genéricos en APIs sin algoritmo común; constraint demasiado restrictiva.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `Contains[T comparable](values []T, target T) bool`. Entrada: [2,4],4 y ["a"],"b". Salida: true,false. Restricción: un recorrido, sin reflect. Aceptación: vacío false y pruebas con dos tipos.

[Enunciado y espacio de implementación](exercises/09-generics/guided/README.md) · [Pistas progresivas](exercises/09-generics/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Define un tipo `Queue[T any]` solo para operaciones secuenciales Push y Pop con bool de presencia. Entrada: insertar 1,2 y extraer tres veces. Salida: 1,true;2,true;zero,false. Restricción: no concurrencia ni solución del proyecto final. Aceptación: orden FIFO y documentar ausencia.

[Enunciado y espacio](exercises/09-generics/challenge/README.md) · [Pistas opcionales](exercises/09-generics/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Any permite == en el cuerpo?
2. ¿Interfaz o genérico para io.Reader?

[Respuestas razonadas](self-assessment/09-generics.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Reutilizar un algoritmo para varios tipos no debe obligarte a perder comprobaciones usando any. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Generics tutorial](https://go.dev/doc/tutorial/generics) · [Type parameters](https://go.dev/ref/spec#Type_parameter_declarations)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
