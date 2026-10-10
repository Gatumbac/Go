# 01 · Structs

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Strings, bytes y runes](../01-fundamentals/09-text.md) · Siguiente: [Punteros y valores](02-pointers.md)

## Objetivos de aprendizaje

- Explicar structs y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Strings, bytes y runes](../01-fundamentals/09-text.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Los datos relacionados necesitan nombres y reglas; una lista de argumentos sueltos no expresa el dominio.

## Explicación conceptual

Un struct agrupa campos con tipos. El zero value inicializa cada campo, pero no garantiza invariantes de negocio. Un constructor es una función convencional, no sintaxis especial: puede devolver un error si el objeto necesita validación. No añadas constructor cuando el zero value sea suficiente.

Los literals con nombres de campos resisten mejor cambios de orden que los posicionales. La exportación usa mayúsculas también en campos. Un campo no exportado puede mantener detalles dentro del package; eso es encapsulación, no una garantía de que ninguna copia comparta recursos.

Asignar un struct copia sus campos. Si contiene arrays se copian sus valores; si contiene slices, maps o pointers se copian esos valores y puede persistir aliasing. Por eso «struct se copia» no equivale a «deep copy». Define propiedad de colecciones antes de devolver objetos mutables.

Dos structs no son comparables si alguno de sus campos no es comparable, por ejemplo un slice. No recurras inmediatamente a reflect: a menudo una comparación explícita del dominio es más legible. Los tags son metadatos que bibliotecas como encoding/json interpretan; no cambian por sí solos el comportamiento de asignación.

## Modelo mental visual

```text
Book value
  Title string
  Pages int
assignment → field values copied
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/01-structs/main.go):

```go
package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func main() {
	a := Book{Title: "Systems", Pages: 120}
	b := a
	b.Pages = 90
	fmt.Println(a.Pages, b.Pages)
}
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/01-structs
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/01-structs
```

**Resultado e interpretación:** Imprime `120 90`: cambiar el campo int de b no cambia a.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Una clase Java se suele compartir mediante referencias. Un struct Go ordinario es un valor; no incluye herencia ni una identidad obligatoria.

## Errores comunes

Usar literales posicionales frágiles; suponer deep copy; añadir validaciones que solo se ejecutan en un constructor pero permitir después campos inválidos.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Define `Product` con Name string y PriceCents int. Implementa `Valid(p Product) bool`: nombre no vacío y precio >=0. Entradas: {"pen",100}, {"",1}, {"pen",-1}. Salidas: true,false,false. Restricción: no I/O. Aceptación: documentar si precio cero es válido (sí).

[Enunciado y espacio de implementación](exercises/01-structs/guided/README.md) · [Pistas progresivas](exercises/01-structs/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Define `Order` con Items []string y escribe en notes.md qué comparte una copia. Entrada: un Order con dos items. Salida: evidencia de cambiar un campo y un elemento. Restricción: tu programa debe demostrar ambos comportamientos sin soluciones de deep copy. Aceptación: explicación correcta del aliasing.

[Enunciado y espacio](exercises/01-structs/challenge/README.md) · [Pistas opcionales](exercises/01-structs/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Copiar un struct independiza todos sus datos?
2. ¿Un constructor es obligatorio?

[Respuestas razonadas](self-assessment/01-structs.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Los datos relacionados necesitan nombres y reglas; una lista de argumentos sueltos no expresa el dominio. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Struct types](https://go.dev/ref/spec#Struct_types)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
