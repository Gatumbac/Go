# 03 · Coverage y fuzzing

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Table-driven tests y subtests](02-tables.md) · Siguiente: [Benchmarks](04-benchmarks.md)

## Objetivos de aprendizaje

- Explicar coverage y fuzzing y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Table-driven tests y subtests](02-tables.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Quieres detectar caminos no ejecutados y entradas que no se te ocurrieron, sin confundir métricas con corrección.

## Explicación conceptual

Coverage registra qué statements se ejecutan. `go test -cover` resume y -coverprofile permite un reporte. Un 100% no prueba combinaciones, condiciones, mensajes correctos ni ausencia de carreras. Usa la métrica para preguntar qué camino falta, no para inflar tests sin propósito.

Fuzzing genera entradas a partir de seeds y mutaciones para buscar violaciones de una propiedad. Define una propiedad útil: ida/vuelta conserva datos, parser no hace panic en entradas arbitrarias, o resultado cumple una invariante. Si la función acepta cualquier string, no filtres tanto que excluyas los casos interesantes.

Las funciones FuzzXxx usan *testing.F; Add registra seeds y Fuzz un callback con tipos admitidos. En ejecución normal go test prueba seeds; para explorar usa -fuzz y -fuzztime. Un caso descubierto puede conservarse como regresión, pero debes revisar qué contrato expuso.

Fuzz tests deben ser deterministas por entrada, rápidos y sin dependencia de servicios persistentes. Limita recursos; generar tamaños desde bytes externos sin límite puede medir OOM en lugar de lógica. Este libro usa un roundtrip de bytes, no un parser de negocio incompleto.

## Modelo mental visual

```text
seeds → mutated inputs → property check
                          ↓ violation
                    minimal regression input
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/03-coverage-fuzz/main.go):

```go
package main

import (
	"encoding/hex"
	"fmt"
)

func roundtrip(data []byte) ([]byte, error) { return hex.DecodeString(hex.EncodeToString(data)) }
func main()                                 { out, err := roundtrip([]byte{0, 255}); fmt.Println(out, err) }
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/03-coverage-fuzz
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/03-coverage-fuzz
```

**Resultado e interpretación:** Imprime `[0 255] <nil>`. FuzzRoundtrip comprueba bytes incluyendo valores que no son texto.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/03-coverage-fuzz/main_test.go):

```go
package main

import (
	"bytes"
	"testing"
)

func FuzzRoundtrip(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{0, 255})
	f.Fuzz(func(t *testing.T, in []byte) {
		out, err := roundtrip(in)
		if err != nil || !bytes.Equal(in, out) {
			t.Fatalf("roundtrip: %v", err)
		}
	})
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -coverprofile=/tmp/go-course-cover.out ./modules/04-tooling/examples/03-coverage-fuzz
go tool cover -func=/tmp/go-course-cover.out
go test -fuzz FuzzRoundtrip -fuzztime=3s -parallel=2 ./modules/04-tooling/examples/03-coverage-fuzz
```

## Comparación con otros lenguajes

Property-based testing en otros ecosistemas persigue invariantes similares. Fuzzing integrado no decide por ti cuál es una propiedad correcta.

## Errores comunes

Coverage como garantía; fuzz que compara consigo mismo; excluir todo input difícil; runs infinitos sin presupuesto.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe tests con coverage para `SumPositive(values []int) int`, dominio -100..100 y longitud <=100. Entrada:[-1,2,3]. Salida:5. Aceptación: cero, vacío y mezcla probados; interpreta qué no demuestra el porcentaje.

[Enunciado y espacio de implementación](exercises/03-coverage-fuzz/guided/README.md) · [Pistas progresivas](exercises/03-coverage-fuzz/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Crea un fuzz test de una transformación reversible escrita por ti, con strings acotados a 1024 bytes. Entrada: bytes arbitrarios. Salida: roundtrip igual. Restricción: sin servicios externos. Aceptación: seeds vacíos y multibyte, run de duración limitada.

[Enunciado y espacio](exercises/03-coverage-fuzz/challenge/README.md) · [Pistas opcionales](exercises/03-coverage-fuzz/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿100% coverage garantiza ausencia de bugs?
2. ¿Qué aporta un seed?

[Respuestas razonadas](self-assessment/03-coverage-fuzz.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Quieres detectar caminos no ejecutados y entradas que no se te ocurrieron, sin confundir métricas con corrección. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Fuzzing](https://go.dev/doc/fuzz/) · [Coverage](https://go.dev/blog/cover)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
