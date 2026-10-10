# 04 · Benchmarks

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Coverage y fuzzing](03-coverage-fuzz.md) · Siguiente: [Gofmt y go vet](05-format-vet.md)

## Objetivos de aprendizaje

- Explicar benchmarks y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Coverage y fuzzing](03-coverage-fuzz.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Necesitas comparar coste de una operación bajo una carga controlada, no cronometrar todo un comando.

## Explicación conceptual

BenchmarkXxx recibe *testing.B. En Go moderno `for b.Loop()` expresa las iteraciones y administra el temporizador del loop; la forma clásica con b.N sigue disponible y debes reconocerla. Preparación fuera del loop evita incluir costes que no pertenecen a la operación.

`go test -run '^$' -bench . -benchmem` selecciona benchmarks sin tests ordinarios. ns/op, B/op y allocs/op describen ese escenario. No equivalen a latencia p99 de un servidor ni a throughput bajo contención. Un benchmark concurrente necesita un diseño acorde al sistema.

Evita que el compilador elimine trabajo, por ejemplo conservando un resultado observado. No uses una entrada constante trivial si la carga real tiene variedad. Compara resultados correctos primero. Repite mediciones en un equipo estable y anota toolchain, CPU, flags y tamaño.

Para profiling añade -cpuprofile o -memprofile y analiza con go tool pprof. La recogida añade coste y el resultado es evidencia para localizar, no una clasificación moral de funciones. Optimizaciones pequeñas no justifican complicar la API si no mejoran una necesidad real.

## Modelo mental visual

```text
same workload → variant A / variant B
                  ↓ measure repeated runs
          correctness + meaningful improvement
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/04-benchmarks/main.go):

```go
package main

import "fmt"

func total(values []int) int {
	n := 0
	for _, v := range values {
		n += v
	}
	return n
}
func main() { fmt.Println(total([]int{2, 3, 4})) }
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/04-benchmarks
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/04-benchmarks
```

**Resultado e interpretación:** Imprime `9`. BenchmarkTotal mide la suma, excluyendo la creación de la entrada.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/04-benchmarks/main_test.go):

```go
package main

import "testing"

var result int

func BenchmarkTotal(b *testing.B) {
	values := make([]int, 100)
	for i := range values {
		values[i] = i
	}
	b.ReportAllocs()
	for b.Loop() {
		result = total(values)
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -run '^$' -bench . -benchmem ./modules/04-tooling/examples/04-benchmarks
go test -run '^$' -bench . -cpuprofile=/tmp/go-course-cpu.prof -o /tmp/go-course-bench.test ./modules/04-tooling/examples/04-benchmarks
go tool pprof -top /tmp/go-course-cpu.prof
```

## Comparación con otros lenguajes

console.time mide intervalos pero no administra calentamiento, repetición o optimización como un harness. JMH también evita trampas de microbenchmarks.

## Errores comunes

Benchmarks sin verificar resultado; stdout dentro del loop; confundir ns/op con p99; optimizaciones eliminadas.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `JoinWords(words []string) string` con separador coma y mide dos estrategias tuyas. Entrada:["a","b"]. Salida:"a,b". Restricción: mismo contrato y tamaño de entrada. Aceptación: tests primero y reporte benchmem.

[Enunciado y espacio de implementación](exercises/04-benchmarks/guided/README.md) · [Pistas progresivas](exercises/04-benchmarks/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Compara búsqueda lineal y lookup en map para claves fijas generadas antes del loop. Entrada: 1000 claves, mitad presentes. Salida: mismo conteo. Aceptación: declarar si incluyes construcción del map y justificarlo.

[Enunciado y espacio](exercises/04-benchmarks/challenge/README.md) · [Pistas opcionales](exercises/04-benchmarks/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Por qué preparar fuera del loop?
2. ¿Un benchmark prueba rendimiento productivo?

[Respuestas razonadas](self-assessment/04-benchmarks.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Necesitas comparar coste de una operación bajo una carga controlada, no cronometrar todo un comando. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Testing benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
