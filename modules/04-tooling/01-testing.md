# 01 · Go test

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Asignaciones y profiling](../03-memory/07-profiling.md) · Siguiente: [Table-driven tests y subtests](02-tables.md)

## Objetivos de aprendizaje

- Explicar go test y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Asignaciones y profiling](../03-memory/07-profiling.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una prueba automatiza un contrato para detectar regresiones sin revisar manualmente cada ejecución.

## Explicación conceptual

Go descubre funciones TestXxx en archivos `_test.go`. Reciben *testing.T y reportan fallos con Error/Fatal y sus variantes. Error marca fallo y permite seguir; Fatal detiene la goroutine del test actual. No invoques Fatal desde un worker esperando que termine correctamente el test padre.

El test prepara datos, ejecuta comportamiento y compara resultado. Comprueba un contrato observable, no cada paso privado de la implementación. Un test que calcula el esperado usando la misma función que prueba no es evidencia independiente. Nombres claros explican qué escenario falla.

`go test` compila el package y un binario de pruebas. Puede reutilizar resultados de tests en ciertas condiciones; `-count=1` fuerza nueva ejecución cuando importa. Usa `-run` para seleccionar un test y `-v` para ver nombres. El package del test puede ser el mismo o uno externo con sufijo _test, según necesites ver detalles internos o API exportada.

t.TempDir y t.Cleanup administran recursos temporales y limpieza. Tests independientes no requieren orden; si dependen de estado global compartido, paralelizarlos puede esconder o crear fallos. Primero elimina acoplamiento, luego considera t.Parallel.

## Modelo mental visual

```text
arrange inputs → act → assert observable result
                         ↓
                  failure identifies contract
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/01-testing/main.go):

```go
package main

import "fmt"

func twice(n int) int { return n * 2 }
func main()           { fmt.Println(twice(6)) }
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/01-testing
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/01-testing
```

**Resultado e interpretación:** Imprime `12`. Ejecuta también go test en este directorio: TestTwice comprueba resultado sin interceptar stdout.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/01-testing/main_test.go):

```go
package main

import "testing"

func TestTwice(t *testing.T) {
	if got := twice(6); got != 12 {
		t.Fatalf("twice(6)=%d, want 12", got)
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -v -count=1 ./modules/04-tooling/examples/01-testing
```

## Comparación con otros lenguajes

Jest y JUnit tienen runners externos; testing está integrado. El objetivo sigue siendo el contrato, no el porcentaje por sí mismo.

## Errores comunes

Expected calculado con la misma función; tests que dependen del orden; Fatal desde una goroutine auxiliar.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `Abs(n int) int` en package practice, con dominio -1000..1000. Entrada: -4,0,7. Salida: 4,0,7. Escribe tus tests antes de la función. Aceptación: fallo útil para cada caso y go test verde.

[Enunciado y espacio de implementación](exercises/01-testing/guided/README.md) · [Pistas progresivas](exercises/01-testing/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Prueba una función `IsEven(n int) bool` escrita por ti. Entrada: -2,-1,0,3. Salida: true,false,true,false. Restricción: tests sin I/O y sin dependencia entre casos. Aceptación: prueba negativa demuestra que el test detecta un cambio incorrecto.

[Enunciado y espacio](exercises/01-testing/challenge/README.md) · [Pistas opcionales](exercises/01-testing/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Test verde demuestra comprensión?
2. ¿Cuándo usar Fatal?

[Respuestas razonadas](self-assessment/01-testing.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una prueba automatiza un contrato para detectar regresiones sin revisar manualmente cada ejecución. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Testing](https://pkg.go.dev/testing)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
