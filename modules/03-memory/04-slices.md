# 04 · Slices y backing arrays

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Escape analysis](03-escape.md) · Siguiente: [Maps](05-maps.md)

## Objetivos de aprendizaje

- Explicar slices y backing arrays y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Escape analysis](03-escape.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un append puede compartir almacenamiento o crear otro; necesitas predecir ambos casos sin confiar en un factor de crecimiento.

## Explicación conceptual

Modela un slice como acceso a un array, longitud y capacidad. No es una garantía de layout del ABI, pero explica operaciones. `s[a:b]` crea otra ventana del mismo almacenamiento; sus elementos pueden seguir compartidos. Un slice de longitud pequeña puede retener un array grande.

Si append cabe en capacidad, reutiliza el backing array. Si no, obtiene otro suficientemente grande y copia los elementos. Guardar el resultado es necesario en ambos casos porque cambia la longitud. El crecimiento exacto depende de runtime y tipos; nunca pruebes una regla universal de duplicación.

La forma `s[a:b:c]` limita capacidad a c-a. Puede forzar que un append posterior obtenga otro array, pero no independiza los elementos ya compartidos. Para independencia debes copiar los elementos; para slices de objetos pointer esa copia sigue sin ser profunda.

Reducir longitud no elimina necesariamente referencias de la parte oculta del array. Si gestionas una colección de pointers que retiene objetos innecesarios, limpiar posiciones retiradas puede evitar retención. Evalúa coste y ownership antes de modificar buffers compartidos.

## Modelo mental visual

```text
backing array [1 2 3 4]
a = [0:2] → len 2 cap 4
b = [0:2:2] → len 2 cap 2
append(b, ...) requires different storage
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/04-slices/main.go):

```go
package main

import "fmt"

func main() { a := []int{1, 2, 3, 4}; b := a[:2:2]; b = append(b, 9); b[0] = 7; fmt.Println(a, b) }
```

Desde la raíz:

```sh
go run ./modules/03-memory/examples/04-slices
go build -o /tmp/go-from-zero-example ./modules/03-memory/examples/04-slices
```

**Resultado e interpretación:** Imprime `[1 2 3 4] [7 2 9]`: append superó capacidad limitada y creó almacenamiento diferente.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

ArrayList oculta capacidad de backing storage; Go expone len/cap y ventanas, así que aliasing debe ser parte del contrato.

## Errores comunes

Reslice como copia; cap como longitud; full slice como deep copy; append sin guardar resultado.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea un experimento con array [4]int y dos ventanas: una con cap suficiente y otra limitada. Entrada: append del mismo valor. Salida: tablas antes/después de mutar. Aceptación: justificar cuándo se comparte almacenamiento.

[Enunciado y espacio de implementación](exercises/04-slices/guided/README.md) · [Pistas progresivas](exercises/04-slices/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `HeadCopy(values []int, n int) []int` con 0<=n<=len. Entrada: [1,2,3],2. Salida independiente [1,2]. Restricción: no solo reslicing. Aceptación: mutación del resultado no cambia input.

[Enunciado y espacio](exercises/04-slices/challenge/README.md) · [Pistas opcionales](exercises/04-slices/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Limitar cap elimina aliasing actual?
2. ¿Por qué un slice pequeño retiene mucho?

[Respuestas razonadas](self-assessment/04-slices.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un append puede compartir almacenamiento o crear otro; necesitas predecir ambos casos sin confiar en un factor de crecimiento. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Slices internals](https://go.dev/blog/slices-intro)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
