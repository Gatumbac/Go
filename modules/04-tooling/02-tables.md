# 02 · Table-driven tests y subtests

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Go test](01-testing.md) · Siguiente: [Coverage y fuzzing](03-coverage-fuzz.md)

## Objetivos de aprendizaje

- Explicar table-driven tests y subtests y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Go test](01-testing.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Repetir tests casi iguales hace difícil detectar qué caso falta y mantener mensajes consistentes.

## Explicación conceptual

Un table-driven test reúne entradas y esperados en una colección de casos. El loop reduce boilerplate, pero la tabla debe reflejar particiones del contrato: vacío, fronteras, valores típicos y errores. No significa probar todos los enteros.

`t.Run` crea subtests nombrados. Puedes seleccionar un caso con -run y recibir un mensaje localizado. Cada subtest conserva preparación y assertions propias. Si hay recursos compartidos, piensa antes de llamar Parallel: tiempos y globales pueden cambiar la independencia.

Los rangos modernos con := crean variables por iteración; no enseñes `tc := tc` como una necesidad eterna. Si asignas a una variable previa con =, cambia el caso. También puede haber aliasing en campos slice o pointer aunque el valor de caso sea una copia.

Usa cmp detallado solo si lo necesitas; la biblioteca estándar cubre muchas pruebas pequeñas. No introduzcas un framework por costumbre. El test debe indicar entrada, obtenido y esperado sin requerir reproducir el fallo para entenderlo.

## Modelo mental visual

```text
table rows → subtest names → independent assertions
empty / boundary / normal / invalid
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/02-tables/main.go):

```go
package main

import "fmt"

func sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
func main() { fmt.Println(sign(-9), sign(0), sign(9)) }
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/02-tables
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/02-tables
```

**Resultado e interpretación:** Imprime `-1 0 1`; los subtests comprueban esas tres particiones.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/02-tables/main_test.go):

```go
package main

import "testing"

func TestSign(t *testing.T) {
	cases := []struct {
		name     string
		in, want int
	}{{"negative", -2, -1}, {"zero", 0, 0}, {"positive", 3, 1}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := sign(tc.in); got != tc.want {
				t.Fatalf("sign(%d)=%d want %d", tc.in, got, tc.want)
			}
		})
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -v -run 'TestSign/negative' ./modules/04-tooling/examples/02-tables
```

## Comparación con otros lenguajes

Describe.each de Jest y tests parametrizados de JUnit resuelven una necesidad similar; Go lo expresa con structs y un loop normal.

## Errores comunes

Muchos casos redundantes sin fronteras; nombres vacíos; Parallel sobre recursos comunes; capturas antiguas enseñadas sin versión.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa y prueba `Clamp(n,low,high int) int`, low<=high. Entrada: (-1,0,10),(5,0,10),(12,0,10). Salida:0,5,10. Restricción: tabla con fronteras incluidas. Aceptación: subtests con nombres únicos.

[Enunciado y espacio de implementación](exercises/02-tables/guided/README.md) · [Pistas progresivas](exercises/02-tables/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe tests de `Category(age int) string`: menores de 18 "minor", resto "adult", dominio 0..130. Entrada:0,17,18,130. Salida:minor,minor,adult,adult. Aceptación: detectar el bug de usar >18 en vez de >=18 sin revelar tu implementación.

[Enunciado y espacio](exercises/02-tables/challenge/README.md) · [Pistas opcionales](exercises/02-tables/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Una tabla grande garantiza cobertura de casos?
2. ¿Copiar el caso evita todo aliasing?

[Respuestas razonadas](self-assessment/02-tables.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Repetir tests casi iguales hace difícil detectar qué caso falta y mantener mensajes consistentes. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Subtests](https://go.dev/blog/subtests)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
