# 03 · Métodos

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Punteros y valores](02-pointers.md) · Siguiente: [Interfaces y method sets](04-interfaces.md)

## Objetivos de aprendizaje

- Explicar métodos y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Punteros y valores](02-pointers.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

El comportamiento asociado a un tipo debe tener un contrato de mutación claro, sin ocultarlo bajo clases.

## Explicación conceptual

Un método es una función con receiver. Puedes definirlo para un tipo definido en tu mismo package; no para tipos externos arbitrarios. El receiver se recibe por valor como cualquier parámetro. Un value receiver ve una copia; un pointer receiver puede modificar la variable compartida.

Escoge receiver según semántica. Un objeto mutable con métodos que cambian estado suele necesitar pointers de forma consistente. Tipos pequeños e inmutables pueden usar valores. No copies tipos que contengan Mutex una vez usados: su estado de sincronización no debe duplicarse como datos normales.

Go permite llamar a un método con pointer receiver sobre una variable direccionable usando automáticamente su dirección. Esto facilita llamadas, pero no cambia las reglas del method set al satisfacer interfaces. Un valor temporal no siempre es direccionable; las posiciones de maps tampoco lo son.

Los métodos no son virtuales por herencia. El despacho mediante interfaces depende del valor dinámico. Un nombre bien elegido comunica qué cambia y qué devuelve; llamar Set a una operación que no altera nada será confuso aunque compile.

## Modelo mental visual

```text
value receiver → copy → local changes
pointer receiver → address copy → shared changes
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/03-methods/main.go):

```go
package main

import "fmt"

type Counter struct{ N int }

func (c *Counter) Add(n int) { c.N += n }
func (c Counter) Value() int { return c.N }
func main()                  { c := Counter{}; c.Add(3); fmt.Println(c.Value()) }
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/03-methods
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/03-methods
```

**Resultado e interpretación:** Imprime `3`. c es direccionable; Go facilita la llamada a Add usando &c.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Los métodos Java reciben implícitamente una referencia this. Go permite elegir explícitamente el tipo del receiver.

## Errores comunes

Value receiver esperando mutación persistente; copiar un Mutex; inferir satisfacción de interfaz por una llamada que usa conversión automática.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Define `Rectangle` con Width, Height int y método `Area() int`. Entrada: 3,4. Salida: 12. Restricción: dimensiones >=0 y receiver por valor. Aceptación: ninguna mutación y caso con cero.

[Enunciado y espacio de implementación](exercises/03-methods/guided/README.md) · [Pistas progresivas](exercises/03-methods/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Define `Balance` con Amount int y método `Deposit(n int)` para n>=0. Entrada: saldo 10, depósito 5. Salida: saldo 15. Restricción: pointer receiver, sin concurrencia. Aceptación: explicar diferencia frente a receiver por valor.

[Enunciado y espacio](exercises/03-methods/challenge/README.md) · [Pistas opcionales](exercises/03-methods/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Una llamada automática agrega métodos al value type?
2. ¿Por qué no copiar Mutex?

[Respuestas razonadas](self-assessment/03-methods.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

El comportamiento asociado a un tipo debe tener un contrato de mutación claro, sin ocultarlo bajo clases. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Methods](https://go.dev/ref/spec#Method_declarations)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
