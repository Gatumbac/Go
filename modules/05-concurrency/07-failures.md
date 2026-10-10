# 07 · Carreras, deadlocks y leaks

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Mutex y RWMutex](06-mutex.md) · Siguiente: [Context y cancelación](08-context.md)

## Objetivos de aprendizaje

- Explicar carreras, deadlocks y leaks y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Mutex y RWMutex](06-mutex.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un programa concurrente puede fallar sin error explícito: datos corruptos, esperas circulares o tareas que nunca terminan.

## Explicación conceptual

Data race es un fallo de sincronización de memoria. Race condition de negocio es dependencia de un interleaving que viola una regla; incluso operaciones atómicas aisladas pueden formar un check-then-act incorrecto. Protege la unidad lógica.

Deadlock ocurre cuando actividades esperan eventos que ninguna puede producir. Un ciclo de locks es un caso, no el único. Una recepción sin productor también puede bloquear indefinidamente. El runtime detecta algunos casos globales de bloqueo, pero un servidor con otras goroutines activas puede mantener una tarea bloqueada sin panic global.

Goroutine leak es trabajo retenido que no puede completar o que ya nadie necesita. Un productor bloqueado enviando cuando su consumidor salió es común. Aumentar el buffer puede aplazar el problema sin resolver cancelación. El diseño debe permitir abortar envíos y demostrar terminación.

Investiga con stacks, race detector y tests con deadlines. No uses sleeps para reparar ordering. En docs/debugging hay casos intencionalmente defectuosos separados del recorrido. Esta lección normal muestra una actividad que sí termina. El perfil experimental goroutineleak de Go 1.26 es opcional, no requisito ni detector absoluto.

## Modelo mental visual

```text
A waits for B
B waits for A → deadlock
producer waits for vanished consumer → potential leak
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/07-failures/main.go):

```go
package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func() { defer close(done) }()
	<-done
	fmt.Println("joined")
}
```

Desde la raíz:

```sh
go run ./modules/05-concurrency/examples/07-failures
go build -o /tmp/go-from-zero-example ./modules/05-concurrency/examples/07-failures
```

**Resultado e interpretación:** Imprime `joined`; el cierre comunica terminación del único productor.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Timeouts de requests JS no eliminan automáticamente trabajo externo; goroutines tampoco desaparecen porque el caller perdió interés.

## Errores comunes

Buffer como arreglo universal; sleeps como coordinación; esperar que runtime detecte todo deadlock.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Reproduce el contador defectuoso de docs/debugging.md en un temporal, ejecuta con -race y escribe el objeto compartido y las dos stacks relevantes. Entrada:1000 incrementos por tarea. Salida:diagnóstico, no cifra exacta. Aceptación: explicar acceso y plantear reparación tuya.

[Enunciado y espacio de implementación](exercises/07-failures/guided/README.md) · [Pistas progresivas](exercises/07-failures/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Diseña un test de terminación de un productor cuyo consumidor deja de recibir. Entrada:cancelación. Salida:join antes del deadline del test. Restricción: canal de cancelación explícito. Aceptación: prueba tanto éxito como cancelación.

[Enunciado y espacio](exercises/07-failures/challenge/README.md) · [Pistas opcionales](exercises/07-failures/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Buffer suficiente elimina necesidad de cancelación?
2. ¿Runtime detecta todo deadlock?

[Respuestas razonadas](self-assessment/07-failures.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un programa concurrente puede fallar sin error explícito: datos corruptos, esperas circulares o tareas que nunca terminan. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Pipelines](https://go.dev/blog/pipelines) · [Go 1.26](https://go.dev/doc/go1.26)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
