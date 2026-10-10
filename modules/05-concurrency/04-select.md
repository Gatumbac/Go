# 04 · Select

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Channels](03-channels.md) · Siguiente: [WaitGroups](05-waitgroups.md)

## Objetivos de aprendizaje

- Explicar select y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Channels](03-channels.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una actividad puede necesitar esperar datos o cancelar sin quedar atrapada en una sola recepción.

## Explicación conceptual

Select elige entre operaciones de channel que puedan progresar. Si hay varios cases listos, elige uno mediante selección pseudoaleatoria uniforme según especificación; no garantiza prioridad a cancelación ni a un case escrito primero. Si ninguno está listo, bloquea salvo que exista default.

Default convierte la operación en no bloqueante; dentro de un loop sin pausa puede crear busy spin. No lo añadas por miedo a bloqueo si la espera es el comportamiento deseado. Un channel nil nunca está listo y permite deshabilitar dinámicamente un case.

Al cerrar un channel, su recepción queda lista de forma permanente después de drenar. En un select repetido, debes terminar o deshabilitarlo para no leer ceros indefinidamente. Comprueba el bool cuando necesitas distinguir dato de fin.

Timeouts suelen representarse con context o timers. `time.After` sirve para esperas simples; crear uno en cada iteración puede asignar trabajo innecesario. Un timeout no interrumpe por fuerza la función arbitraria que no escucha. Diseña salida cooperativa y comprueba que las goroutines terminan.

## Modelo mental visual

```text
wait
 ├─ data ready → process
 ├─ cancel ready → return
 └─ neither → block (unless default)
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/04-select/main.go):

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 8
	select {
	case n := <-ch:
		fmt.Println(n)
	default:
		fmt.Println("empty")
	}
	select {
	case <-ch:
		fmt.Println("data")
	default:
		fmt.Println("empty")
	}
}
```

Desde la raíz:

```sh
go run ./modules/05-concurrency/examples/04-select
go build -o /tmp/go-from-zero-example ./modules/05-concurrency/examples/04-select
```

**Resultado e interpretación:** Imprime `8` y `empty`. Los cases listos son deterministas en este escenario.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Promise.race devuelve una primera resolución, pero no cancela automáticamente las demás actividades. Select tampoco garantiza terminar el trabajo del perdedor.

## Errores comunes

Default en busy loop; prioridad por orden source; channel cerrado leído sin ok; creer que timeout mata goroutine.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `TryReceive(ch <-chan int) (int,bool)` como recepción no bloqueante para channel abierto. Entrada: canal con 7 y canal vacío. Salida:(7,true),(0,false). Aceptación: no sleeps; documenta closed channel como recepción disponible con dato zero si decides ampliar contrato.

[Enunciado y espacio de implementación](exercises/04-select/guided/README.md) · [Pistas progresivas](exercises/04-select/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Diseña un consumidor que espera datos o señal done y termina al cerrar cualquiera según contrato. Entrada: secuencia de dos datos y cierre. Salida: conteo recibido. Aceptación: no loop infinito con channel cerrado.

[Enunciado y espacio](exercises/04-select/challenge/README.md) · [Pistas opcionales](exercises/04-select/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿El primer case tiene prioridad?
2. ¿Default evita leaks?

[Respuestas razonadas](self-assessment/04-select.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una actividad puede necesitar esperar datos o cancelar sin quedar atrapada en una sola recepción. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Select](https://go.dev/ref/spec#Select_statements)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
