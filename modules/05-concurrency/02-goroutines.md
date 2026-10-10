# 02 · Goroutines y scheduler

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Concurrencia y paralelismo](01-concurrency.md) · Siguiente: [Channels](03-channels.md)

## Objetivos de aprendizaje

- Explicar goroutines y scheduler y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Concurrencia y paralelismo](01-concurrency.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Quieres iniciar trabajo independiente y garantizar que termina antes de abandonar el proceso.

## Explicación conceptual

`go f()` inicia una goroutine que ejecuta la llamada. Los argumentos se evalúan en la goroutine que inicia; después el runtime decide cuándo corre el trabajo. No tienes garantía de orden, fairness exacta ni thread fijo. Main debe esperar explícitamente si necesita sus resultados.

El modelo G/M/P ayuda a explicar scheduler: G representa goroutine, M un thread del sistema y P recursos de ejecución para código Go. Es un detalle del runtime, no una interfaz que tu aplicación deba manipular. GOMAXPROCS limita paralelismo de ejecución de código Go, no el número de goroutines ni todos los threads.

El runtime puede mover actividades entre threads y manejar esperas. Ciertas llamadas y cgo interactúan de otras formas; evita simplificar toda espera como «ningún thread bloqueado». La preemption evita que un trabajo coopere solo en puntos voluntarios, pero no garantiza plazos de negocio.

Las goroutines no tienen un destructor automático de trabajo: si bloquean sin posibilidad de desbloquear, pueden permanecer vivas. Diseña terminación, señal de cancelación y espera del caller. Un channel de resultado en este ejemplo demuestra el join; no necesitas Sleep para dar tiempo.

## Modelo mental visual

```text
main → launch G → receive result → return
             ↓
       runtime schedules on threads
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/02-goroutines/main.go):

```go
package main

import "fmt"

func main() { done := make(chan int); go func(n int) { done <- n * n }(5); fmt.Println(<-done) }
```

Desde la raíz:

```sh
go run ./modules/05-concurrency/examples/02-goroutines
go build -o /tmp/go-from-zero-example ./modules/05-concurrency/examples/02-goroutines
```

**Resultado e interpretación:** Imprime `25`; main espera el resultado mediante comunicación.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Un thread Java suele tener más coste de creación/stack; una goroutine es más ligera pero no gratuita. Async Python depende de cooperación y su runtime específico.

## Errores comunes

Sleep para esperar; main retorna antes del worker; GOMAXPROCS como límite de conexiones.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea un programa con tres goroutines que calculen el doble de 1,2,3 y entreguen resultados por channel. Entrada: esos números. Salida: suma 12. Restricción: no sleeps. Aceptación: main espera tres resultados y no exige orden.

[Enunciado y espacio de implementación](exercises/02-goroutines/guided/README.md) · [Pistas progresivas](exercises/02-goroutines/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe scheduler-notes.md explicando una carga con 1000 goroutines y GOMAXPROCS=2. Entrada: esa configuración. Salida: qué limita, qué no limita y qué medir. Aceptación: no afirmar que existen solo dos threads.

[Enunciado y espacio](exercises/02-goroutines/challenge/README.md) · [Pistas opcionales](exercises/02-goroutines/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Main espera goroutines al retornar?
2. ¿GOMAXPROCS limita número de workers?

[Respuestas razonadas](self-assessment/02-goroutines.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Quieres iniciar trabajo independiente y garantizar que termina antes de abandonar el proceso. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Runtime](https://pkg.go.dev/runtime) · [Go statements](https://go.dev/ref/spec#Go_statements)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
