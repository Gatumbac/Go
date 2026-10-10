# 01 · Concurrencia y paralelismo

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Race detector](../04-tooling/07-race.md) · Siguiente: [Goroutines y scheduler](02-goroutines.md)

## Objetivos de aprendizaje

- Explicar concurrencia y paralelismo y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Race detector](../04-tooling/07-race.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Varias actividades pueden progresar mientras esperan I/O; eso no exige que usen CPU al mismo instante.

## Explicación conceptual

Concurrencia es organizar actividades con progreso potencialmente solapado. Paralelismo es ejecución simultánea, normalmente usando varios cores. Un sistema concurrente funciona incluso con una CPU, alternando trabajo; agregar CPUs no arregla una dependencia secuencial.

Un proceso tiene recursos y espacio de direcciones; threads de un proceso comparten memoria y son planificados por el sistema operativo. Goroutines son unidades administradas por el runtime Go y se ejecutan sobre threads. No confundas goroutine con proceso aislado: comparten objetos y pueden competir por datos.

El beneficio de concurrencia depende del trabajo. Mientras una petición espera red, otra puede ejecutar; si todas esperan una misma conexión o un lock, lanzar más goroutines solo añade cola y consumo. Las tareas CPU-bound se benefician de paralelismo hasta límites de cores, memoria y algoritmo.

Diseña primero una gráfica de dependencias y presupuesto de recursos. Backpressure limita trabajo admitido para no convertir sobrecarga en OOM. La ley práctica es simple: más actividades activas no equivale a más capacidad. Define qué se puede solapar y qué debe ordenarse.

## Modelo mental visual

```text
one core: A → B → A → B (concurrent scheduling)
two cores: core1 AAAAA
           core2 BBBBB (possible parallel execution)
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/01-concurrency/main.go):

```go
package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	go func() { ch <- "alpha" }()
	go func() { ch <- "beta" }()
	for range 2 {
		fmt.Println(<-ch)
	}
}
```

Desde la raíz:

```sh
go run ./modules/05-concurrency/examples/01-concurrency
go build -o /tmp/go-from-zero-example ./modules/05-concurrency/examples/01-concurrency
```

**Resultado e interpretación:** Imprime alpha y beta, en cualquier orden. No demuestra que se ejecutaron en paralelo.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Promise.all expresa concurrencia de actividades async, pero no vuelve CPU-bound JavaScript paralelo en un único thread. Goroutines permiten otro modelo sin quitar dependencias.

## Errores comunes

Orden de logs como prueba de paralelismo; goroutine=thread; concurrencia ilimitada para arreglar saturación.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Dibuja dos timelines para descargar dos recursos y luego combinarlos. Entrada: dependencias dadas. Salida: versión secuencial y concurrente con join. Aceptación: combinación ocurre después de ambas descargas.

[Enunciado y espacio de implementación](exercises/01-concurrency/guided/README.md) · [Pistas progresivas](exercises/01-concurrency/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Diseña un presupuesto para 10000 requests con límite de 20 operaciones activas. Entrada: ese escenario. Salida: política de cola/rechazo y métricas de espera. Restricción: no implementar pool todavía. Aceptación: no prometer capacidad ilimitada.

[Enunciado y espacio](exercises/01-concurrency/challenge/README.md) · [Pistas opcionales](exercises/01-concurrency/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Concurrencia exige varios cores?
2. ¿Qué no mejora más goroutines?

[Respuestas razonadas](self-assessment/01-concurrency.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Varias actividades pueden progresar mientras esperan I/O; eso no exige que usen CPU al mismo instante. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Concurrency is not parallelism](https://go.dev/blog/waza-talk)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
