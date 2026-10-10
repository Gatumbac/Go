# 06 · Operadores y control de flujo

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Variables, constantes y tipos](05-types.md) · Siguiente: [Funciones y retornos](07-functions.md)

## Objetivos de aprendizaje

- Explicar operadores y control de flujo y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Variables, constantes y tipos](05-types.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

El comportamiento depende de condiciones; hacerlo explícito evita rutas que olvidan errores o mezclan casos incompatibles.

## Explicación conceptual

Go usa if, switch y for. No tiene while separado: `for condition` expresa esa repetición, `for {}` una repetición sin condición, y `for init; condition; post` el recorrido clásico. Break sale del loop más cercano; continue pasa a la siguiente iteración. Las labels existen, pero a menudo extraer funciones simplifica control anidado.

Un if puede tener una instrucción corta antes de su condición: la variable vive en las ramas y no fuera. Es útil para comprobar resultado y error sin ampliar su scope. Las condiciones deben ser bool, no valores truthy. `&&` y `||` evalúan con cortocircuito: úsalo al comprobar un pointer antes de acceder a su campo.

Switch no cae automáticamente al siguiente case. `fallthrough` es explícito y no reevalúa la condición siguiente; no lo uses para imitar Java por costumbre. Un switch sin expresión organiza condiciones como una cadena de if.

Range cambia según el dato: para slices entrega índice y valor; el valor es una copia. Para maps no promete orden. Go moderno crea variables nuevas por iteración cuando se declaran con :=, pero si se reutilizan con = el comportamiento es distinto. Evita enseñar el viejo bug de captura como una regla universal de Go 1.26.

## Modelo mental visual

```text
condition? ── true ──> branch A
    |
   false
    v
 branch B → continuation
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/06-flow/main.go):

```go
package main

import "fmt"

func main() {
	total := 0
	for _, n := range []int{2, 4, 6} {
		if n > 3 {
			total += n
		}
	}
	fmt.Println(total)
}
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/06-flow
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/06-flow
```

**Resultado e interpretación:** Imprime `10`: solo 4 y 6 satisfacen la condición. El underscore descarta el índice.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Java y C suelen necesitar break en switch; Go no. JS usa coerción truthy; Go exige booleanos. El for de Go cubre varias formas sintácticas.

## Errores comunes

Usar asignación en lugar de comparación; esperar orden de map; modificar el valor de range pensando que modifica el slice; asumir variables por iteración sin mirar := o =.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `CountPositive(values []int) int`. Entrada: [-2,0,4,7]. Salida: 2; vacío: 0. Restricción: un recorrido, sin ordenar ni modificar entrada. Aceptación: negativos, cero y positivos tratados según contrato.

[Enunciado y espacio de implementación](exercises/06-flow/guided/README.md) · [Pistas progresivas](exercises/06-flow/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `Classify(score int) string`: 0..49 "retry", 50..79 "pass", 80..100 "strong", fuera de 0..100 "invalid". Entrada: -1,49,50,80,101. Salida: invalid,retry,pass,strong,invalid. Restricción: no mapas. Aceptación: probar fronteras sin ramas solapadas.

[Enunciado y espacio](exercises/06-flow/challenge/README.md) · [Pistas opcionales](exercises/06-flow/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿El valor de range es el elemento original?
2. ¿Switch ejecuta varios cases?

[Respuestas razonadas](self-assessment/06-flow.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

El comportamiento depende de condiciones; hacerlo explícito evita rutas que olvidan errores o mezclan casos incompatibles. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[For statements](https://go.dev/ref/spec#For_statements) · [Switch](https://go.dev/ref/spec#Switch_statements)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
