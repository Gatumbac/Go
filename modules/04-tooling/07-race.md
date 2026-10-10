# 07 · Race detector

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Debugging con Delve](06-delve.md) · Siguiente: [Concurrencia y paralelismo](../05-concurrency/01-concurrency.md)

## Objetivos de aprendizaje

- Explicar race detector y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Debugging con Delve](06-delve.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Los resultados pueden parecer correctos mientras varias goroutines acceden sin sincronización a los mismos datos.

## Explicación conceptual

Una data race combina accesos concurrentes a la misma ubicación, al menos una escritura, sin la relación de sincronización exigida. Puede producir resultados impredecibles aun cuando en algunos runs el número esperado aparezca. Una race condition de negocio es más amplia: puede existir incluso con accesos individuales protegidos.

`go test -race` instrumenta accesos y observa las ejecuciones. El reporte identifica stacks de los accesos y creación de goroutines; busca el objeto y su política de ownership. No significa que toda función del stack sea defectuosa. El overhead cambia timing y coste.

Un run limpio solo cubre los caminos ejecutados. Aumenta escenarios relevantes y utiliza coordinación determinista en tests. Sleeps arbitrarios no establecen una relación happens-before. Un mutex, comunicación por channel o primitivas documentadas sí pueden hacerlo.

Este módulo presenta la herramienta antes de la teoría detallada de concurrencia. El ejemplo usa dos goroutines pero coordina únicamente la impresión de resultados propios, sin datos escritos en común. La lección de fallos mostrará un contador defectuoso separado.

## Modelo mental visual

```text
goroutine A write ──?── goroutine B read
                 no synchronization → data race
lock/channel establishes ordered access
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/07-race/main.go):

```go
package main

import "fmt"

func main() {
	results := make(chan int, 2)
	go func() { results <- 4 }()
	go func() { results <- 6 }()
	a, b := <-results, <-results
	fmt.Println(a + b)
}
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/07-race
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/07-race
```

**Resultado e interpretación:** Imprime `10`. El orden de llegada puede variar, pero la suma y ausencia de escrituras compartidas no dependen de él.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/07-race/main_test.go):

```go
package main

import "testing"

func TestTransfer(t *testing.T) {
	ch := make(chan int)
	go func() { ch <- 17 }()
	if got := <-ch; got != 17 {
		t.Fatalf("got %d", got)
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -race -count=1 ./modules/04-tooling/examples/07-race
```

## Comparación con otros lenguajes

Node.js de un thread reduce ciertos accesos simultáneos a memoria, pero aún tiene race conditions entre operaciones async; workers añaden otras posibilidades.

## Errores comunes

Run limpio como prueba absoluta; Sleep como lock; creer que ++ es atómico.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe race-notes.md sobre dos requests que comprueban stock y luego reservan. Entrada: stock=1, dos compras. Salida: interleaving inválido posible. Aceptación: diferenciar data race de regla de negocio rota.

[Enunciado y espacio de implementación](exercises/07-race/guided/README.md) · [Pistas progresivas](exercises/07-race/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Tras el módulo concurrente, crea un contador protegido y tests bajo -race. Entrada: 4 goroutines, 100 incrementos cada una. Salida:400. Aceptación: no sleeps, espera terminación y explica la sincronización.

[Enunciado y espacio](exercises/07-race/challenge/README.md) · [Pistas opcionales](exercises/07-race/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Un incremento es atómico?
2. ¿Race detector verde prueba todo el programa?

[Respuestas razonadas](self-assessment/07-race.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Los resultados pueden parecer correctos mientras varias goroutines acceden sin sincronización a los mismos datos. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Race detector](https://go.dev/doc/articles/race_detector) · [Memory model](https://go.dev/ref/mem)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
