# 04 · Interfaces y method sets

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Métodos](03-methods.md) · Siguiente: [Composición](05-composition.md)

## Objetivos de aprendizaje

- Explicar interfaces y method sets y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Métodos](03-methods.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Quieres que una operación use una capacidad sin conocer una implementación concreta ni crear una jerarquía.

## Explicación conceptual

Una interfaz describe métodos. Un tipo satisface la interfaz si su method set incluye esos métodos con firmas exactas; no declara implements. Esto permite que el consumidor defina una capacidad pequeña y acepte tipos existentes sin acoplarlos a su package.

El method set de T incluye métodos con receiver T; el de *T incluye los de T y *T, con reglas adicionales para embedding. Por eso un *T puede satisfacer una interfaz que T no satisface, aun cuando una llamada sobre una variable T sea cómoda por direccionabilidad.

Un valor de interfaz combina conceptualmente tipo dinámico y valor dinámico. Esta es una explicación semántica, no una promesa de layout de memoria. Solo es nil cuando no hay tipo ni valor dinámico. Una interfaz que contiene un pointer nil tiene tipo dinámico y no es igual a nil: invocar su método depende de cómo ese método maneje el receiver.

`any` es alias de interface{} y permite cualquier valor, pero desplaza comprobaciones hacia assertions o type switches. Úsalo cuando la heterogeneidad lo requiera, no para evitar diseñar tipos. Una interfaz de un solo método, como io.Reader, suele facilitar composición y testing más que una interfaz que copia todos los métodos de un struct.

### Method sets y llamadas

| Tipo del valor | Métodos declarados en T | Métodos declarados en *T |
| --- | --- | --- |
| T | Pertenecen a su method set. | No pertenecen por comodidad de llamada. |
| *T | Pertenecen también. | Pertenecen. |

Una variable direccionable puede llamar un método de *T como abreviatura de tomar su dirección. Una interfaz que requiere ese método sigue exigiendo el method set adecuado. Prueba tu predicción distinguiendo «la llamada compila» de «el valor satisface la interfaz».

Un typed nil frecuente aparece al devolver un *CustomError nil como error. La intención era «sin error», pero el retorno convierte el pointer a una interfaz con tipo dinámico. El caller observa error != nil. El contrato debe retornar nil como interfaz cuando no hubo fallo, no confiar en el pointer nil.

## Modelo mental visual

```text
consumer → interface capability ← concrete type
interface value: dynamic type + dynamic value
(nil, nil) differs from (*T, nil)
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/04-interfaces/main.go):

```go
package main

import "fmt"

type Describer interface{ Describe() string }
type Planet string

func (p Planet) Describe() string { return "planet " + string(p) }
func show(d Describer)            { fmt.Println(d.Describe()) }
func main()                       { show(Planet("Mars")) }
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/04-interfaces
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/04-interfaces
```

**Resultado e interpretación:** Imprime `planet Mars`; Planet nunca declaró implements.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Las interfaces TS son estructurales y se borran al ejecutar. Go también satisface implícitamente por estructura de métodos, pero usa valores de interfaz en runtime.

## Errores comunes

Interfaces enormes; typed nil; usar any para todo; comparar interfaces cuyo valor dinámico no es comparable.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Define `Sizer` con `Size() int` y dos tipos que lo satisfagan por valor, `Folder` y `Envelope`, con un campo de tamaño. Entrada: tamaños 4 y 9. Salida: `TotalSize([]Sizer)` devuelve 13. Restricción: sin type switch. Aceptación: caller depende solo de Sizer.

[Enunciado y espacio de implementación](exercises/04-interfaces/guided/README.md) · [Pistas progresivas](exercises/04-interfaces/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Investiga typed nil con tu propio tipo `Reporter` y un pointer nil asignado a una interfaz. Entrada: pointer nil. Salida: nota que distingue ambos chequeos nil sin invocar un método inseguro. Aceptación: explicación de tipo dinámico y valor.

[Enunciado y espacio](exercises/04-interfaces/challenge/README.md) · [Pistas opcionales](exercises/04-interfaces/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Por qué interface con pointer nil no es nil?
2. ¿Dónde definir una interfaz pequeña?

[Respuestas razonadas](self-assessment/04-interfaces.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Quieres que una operación use una capacidad sin conocer una implementación concreta ni crear una jerarquía. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Interfaces](https://go.dev/ref/spec#Interface_types) · [FAQ nil error](https://go.dev/doc/faq#nil_error)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
