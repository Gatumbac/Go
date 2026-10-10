# 08 · Arrays, slices y maps

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Funciones y retornos](07-functions.md) · Siguiente: [Strings, bytes y runes](09-text.md)

## Objetivos de aprendizaje

- Explicar arrays, slices y maps y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Funciones y retornos](07-functions.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

No todos los contenedores tienen las mismas reglas de copia, tamaño u ownership: elegirlos cambia el comportamiento del programa.

## Explicación conceptual

Un array `[N]T` contiene N elementos y N es parte del tipo. Asignar un array copia sus elementos. Un slice `[]T` describe una ventana de un array subyacente: tiene longitud y capacidad, y copiarlo no copia necesariamente los elementos. Su zero value es nil y se puede usar con append.

`len` cuenta elementos visibles; `cap` limita hasta dónde puede extenderse el slice sin un nuevo backing array. `append` devuelve el slice resultante: debes guardarlo porque puede cambiar longitud y almacenamiento. No prometas un factor fijo de crecimiento. `copy` copia elementos y devuelve cuántos se copiaron; los objetos apuntados por esos elementos pueden seguir compartidos.

Un map asocia claves comparables con valores. Su lectura de una clave ausente devuelve el zero value; usa `value, ok` para distinguir ausencia de un valor legítimamente cero. Leer un nil map funciona, escribir en él causa panic: inicialízalo con make o literal. Delete de una clave ausente es válido.

Los maps no garantizan orden de iteración, y no son seguros para escritura concurrente sin coordinación. No necesitas elegir un layout interno para usarlos correctamente. Primero escribe una política de propiedad: quién puede modificar y quién observa las colecciones.

### Comparación de contratos

| Contenedor | Tipo/tamaño | Asignación | Zero value |
| --- | --- | --- | --- |
| `[3]int` | Tamaño parte del tipo. | Copia tres enteros. | Tres ceros. |
| `[]int` | Longitud variable. | Copia acceso a una ventana. | nil; append válido. |
| `map[string]int` | Asociación mutable. | Comparte asociación. | nil; lectura válida, escritura inválida. |

Al copiar `[3]*Item` se copian tres pointers, no tres objetos Item. El tipo de los elementos importa incluso cuando el contenedor sí copia sus elementos. Esa observación evita convertir una regla de arrays en una promesa de independencia universal.

## Modelo mental visual

```text
array: [10 20 30 40]
              ↑        ↑
slice: pointer + len + cap
map: key ─────> value (no iteration order contract)
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/08-collections/main.go):

```go
package main

import "fmt"

func main() {
	a := [3]int{4, 5, 6}
	b := a
	b[0] = 99
	s := a[:2]
	s[1] = 8
	counts := map[string]int{"red": 2}
	_, found := counts["blue"]
	fmt.Println(a, b, s, found)
}
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/08-collections
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/08-collections
```

**Resultado e interpretación:** Imprime `[4 8 6] [99 5 6] [4 8] false`. b es una copia; s comparte los elementos de a.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Los arrays JS son contenedores dinámicos de referencias. Un slice Go no equivale a un array JS ni a ArrayList: su descriptor y capacidad explican aliasing.

## Errores comunes

Olvidar asignar append; escribir nil maps; tratar valor cero como ausencia; esperar deep copy de un slice.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `Unique(values []string) []string` conservando primera aparición. Entrada: ["a","b","a"]. Salida: ["a","b"]; vacío: longitud 0. Restricción: no ordenar ni mutar entrada. Aceptación: orden estable y distinguir conjunto de lista de salida.

[Enunciado y espacio de implementación](exercises/08-collections/guided/README.md) · [Pistas progresivas](exercises/08-collections/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `Frequencies(values []string) map[string]int`. Entrada: ["x","x","y"]. Salida: x=2,y=1, sin exigir orden de impresión. Restricción: resultado no nil. Aceptación: vacío produce map vacío y el input no cambia.

[Enunciado y espacio](exercises/08-collections/challenge/README.md) · [Pistas opcionales](exercises/08-collections/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Qué se copia al asignar un slice?
2. ¿Cómo distinguir ausencia de cero en map?

[Respuestas razonadas](self-assessment/08-collections.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

No todos los contenedores tienen las mismas reglas de copia, tamaño u ownership: elegirlos cambia el comportamiento del programa. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Slices](https://go.dev/blog/slices-intro) · [Maps](https://go.dev/blog/maps)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
