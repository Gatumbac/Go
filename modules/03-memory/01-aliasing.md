# 01 · Valores, referencias y aliasing

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Genéricos básicos](../02-design/09-generics.md) · Siguiente: [Stack y heap](02-stack-heap.md)

## Objetivos de aprendizaje

- Explicar valores, referencias y aliasing y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Genéricos básicos](../02-design/09-generics.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Dos nombres pueden acceder al mismo objeto: una mutación local puede tener efectos fuera de su función.

## Explicación conceptual

Aliasing ocurre cuando dos accesos permiten observar el mismo almacenamiento. Go pasa valores por copia, pero algunos valores dan acceso a otros datos: pointers, slices, maps, channels e interfaces con ciertos valores dinámicos. La copia del valor y la independencia de sus datos son preguntas distintas.

Ownership es una política de diseño: quién modifica, durante cuánto tiempo y cómo se comparte. El lenguaje no impone un borrow checker; debes documentar el contrato. Devolver un slice interno permite al caller modificar elementos del objeto aunque no exportes su campo. Devolver una copia de elementos protege la colección, pero no los objetos compartidos por elementos pointer.

Inmutabilidad puede simplificar compartir datos, pero un struct con un slice no se vuelve inmutable por convención de nombre. Una API puede copiar en entrada, copiar en salida o transferir ownership; cada opción tiene costes y condiciones. Escoge según necesidad y comprueba casos de mutación posterior.

En concurrencia, aliasing más escrituras sin sincronización puede producir carreras. Primero identifica el objeto compartido, luego el mecanismo de acceso. Una interfaz no borra esa necesidad. El modelo conceptual de flechas ayuda; no necesitas direcciones concretas ni unsafe para razonar.

## Modelo mental visual

```text
slice A ─┐
         ├──> [shared elements]
slice B ─┘
copy(elements) → [different elements]
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/01-aliasing/main.go):

```go
package main

import "fmt"

func main() { a := []int{3, 7}; b := a; c := append([]int(nil), a...); b[0] = 9; fmt.Println(a, b, c) }
```

Desde la raíz:

```sh
go run ./modules/03-memory/examples/01-aliasing
go build -o /tmp/go-from-zero-example ./modules/03-memory/examples/01-aliasing
```

**Resultado e interpretación:** Imprime `[9 7] [9 7] [3 7]`. c tiene elementos copiados de enteros.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Java comparte referencias a objetos; C expone direcciones manualmente. En Go la combinación de valores y acceso compartido exige mirar campos y contenedores.

## Errores comunes

Copiar solo el descriptor; llamar deep copy a copiar []*T; devolver buffers reutilizados sin avisar.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `Clone(values []int) []int`. Entrada: [1,2]; salida independiente [1,2]. Restricción: no devolver input. Aceptación: cambiar resultado no cambia input; nil puede producir nil.

[Enunciado y espacio de implementación](exercises/01-aliasing/guided/README.md) · [Pistas progresivas](exercises/01-aliasing/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Diseña `Snapshot` de un struct con []string y explica su contrato. Entrada: datos con dos strings. Salida: copia modificable sin afectar original. Restricción: sin reflect. Aceptación: prueba mutación posterior en ambos sentidos.

[Enunciado y espacio](exercises/01-aliasing/challenge/README.md) · [Pistas opcionales](exercises/01-aliasing/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Copiar []*T independiza los T?
2. ¿Qué decide ownership?

[Respuestas razonadas](self-assessment/01-aliasing.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Dos nombres pueden acceder al mismo objeto: una mutación local puede tener efectos fuera de su función. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Assignment](https://go.dev/ref/spec#Assignments)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
