# 05 · Composición

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Interfaces y method sets](04-interfaces.md) · Siguiente: [Errores y wrapping](06-errors.md)

## Objetivos de aprendizaje

- Explicar composición y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Interfaces y method sets](04-interfaces.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Reutilizar comportamiento no requiere afirmar que un tipo es una subclase de otro.

## Explicación conceptual

Composición significa construir un objeto con colaboradores o valores más pequeños. Puedes usar campos nombrados y delegar explícitamente, o embedding para promover campos y métodos. Embedding no convierte el tipo exterior en el interior ni crea una jerarquía de subtipos.

La promoción permite escribir outer.Method sin nombrar el campo embebido cuando no hay ambigüedad. Puede ser útil, pero también expone accidentalmente una API extensa. Dos miembros promovidos con el mismo nombre pueden producir ambigüedad; elegir explícitamente el componente conserva claridad.

Un método del componente trabaja sobre su propio receiver: no se reemplaza mágicamente por un método con igual nombre del tipo exterior. Eso difiere de expectativas de despacho virtual por herencia. Elige campos nombrados si necesitas restringir qué operaciones presenta el objeto.

Para dependencias como reloj o almacenamiento, una interfaz pequeña permite sustituir un colaborador durante tests. No añadas interfaces preventivamente para cada struct: introduce una frontera cuando hay una necesidad concreta de sustituir o separar comportamiento.

## Modelo mental visual

```text
Report
 ├─ metadata (value)
 └─ formatter (collaborator)
explicit delegation → controlled public behavior
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/05-composition/main.go):

```go
package main

import "fmt"

type Label struct{ Text string }

func (l Label) Render() string { return "[" + l.Text + "]" }

type Badge struct {
	Label
	Number int
}

func main() { b := Badge{Label: Label{Text: "guest"}, Number: 7}; fmt.Println(b.Render(), b.Number) }
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/05-composition
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/05-composition
```

**Resultado e interpretación:** Imprime `[guest] 7`; Render se promueve desde Label. Badge no se convirtió en Label.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Java/C# usan herencia para reutilizar y despachar. Go permite composición y promoción, pero sin override virtual por jerarquía.

## Errores comunes

Tratar embedding como extends; exponer métodos que no quieres prometer; interfaces para todo sin necesidad.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Define `Profile` con campo nombrado Address (City string) y Name string. Implementa `Location() string` delegando explícitamente al campo. Entrada: Name="Lee", City="Quito". Salida: "Quito". Restricción: sin embedding. Aceptación: API exterior no expone datos por promoción.

[Enunciado y espacio de implementación](exercises/05-composition/guided/README.md) · [Pistas progresivas](exercises/05-composition/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Modela `Document` que usa una interfaz `Formatter` con `Format(string) string`. Entrada: "draft" y un formatter escrito por ti. Salida: texto formateado. Restricción: dos implementaciones distintas. Aceptación: Document no decide el formato por type switch.

[Enunciado y espacio](exercises/05-composition/challenge/README.md) · [Pistas opcionales](exercises/05-composition/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Embedding crea subtipado?
2. ¿Por qué campos nombrados?

[Respuestas razonadas](self-assessment/05-composition.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Reutilizar comportamiento no requiere afirmar que un tipo es una subclase de otro. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Embedding](https://go.dev/doc/effective_go#embedding)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
