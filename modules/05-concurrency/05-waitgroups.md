# 05 · WaitGroups

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Select](04-select.md) · Siguiente: [Mutex y RWMutex](06-mutex.md)

## Objetivos de aprendizaje

- Explicar waitgroups y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Select](04-select.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Esperar resultados y esperar terminación son necesidades distintas; algunos workers no producen un valor por cada salida.

## Explicación conceptual

WaitGroup cuenta tareas pendientes. Add incrementa, Done decrementa y Wait bloquea hasta cero. Registra tareas antes de permitir que Wait observe un contador cero; llamar Add dentro del worker puede dejar una ventana en que el caller retorna prematuramente.

En Go 1.26 puedes usar WaitGroup.Go, incorporado en Go 1.25, para registrar y arrancar una tarea. La función no debe hacer panic según su contrato. Reconoce también el patrón clásico Add antes de go y defer Done: aparece en mucho código existente.

WaitGroup no protege datos compartidos ni devuelve errores. Si todos escriben un contador sin lock, esperar al final no elimina sus carreras durante la ejecución. Usa channel, mutex o independencia por índice con un contrato correcto.

No copies un WaitGroup después de usarlo. Reutilizarlo requiere que la ronda anterior haya terminado y no mezclar generaciones de trabajo. Combina espera con cancelación: cancelar pide salir, Wait demuestra que efectivamente salieron.

## Modelo mental visual

```text
Add before launch → worker work → Done
                           main Wait → all finished
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/05-waitgroups/main.go):

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	values := make([]int, 3)
	var wg sync.WaitGroup
	for i := range values {
		wg.Go(func() { values[i] = i + 1 })
	}
	wg.Wait()
	fmt.Println(values)
}
```

Desde la raíz:

```sh
go run ./modules/05-concurrency/examples/05-waitgroups
go build -o /tmp/go-from-zero-example ./modules/05-concurrency/examples/05-waitgroups
```

**Resultado e interpretación:** Imprime `[1 2 3]`. Cada worker escribe un elemento distinto y main lee después de Wait.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/05-waitgroups/main_test.go):

```go
package main

import (
	"sync"
	"testing"
)

func TestJoin(t *testing.T) {
	values := make([]int, 8)
	var wg sync.WaitGroup
	for i := range values {
		wg.Go(func() { values[i] = i })
	}
	wg.Wait()
	for i, n := range values {
		if n != i {
			t.Fatalf("index %d got %d", i, n)
		}
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -race -count=1 ./modules/05-concurrency/examples/05-waitgroups
```

## Comparación con otros lenguajes

Promise.all une resultados/rechazos; WaitGroup solo une terminación y necesita otro contrato para errores.

## Errores comunes

Add dentro de goroutine; WaitGroup como mutex; olvidar Done en una rama; copiarlo.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea cinco workers que escriban cuadrados en índices distintos de un slice. Entrada: índices 0..4. Salida:[0,1,4,9,16]. Restricción: WaitGroup, sin sleeps. Aceptación: main lee tras join y -race limpio.

[Enunciado y espacio de implementación](exercises/05-waitgroups/guided/README.md) · [Pistas progresivas](exercises/05-waitgroups/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa dos rondas secuenciales de trabajo y registra cuándo se puede reutilizar WaitGroup. Entrada: dos lotes pequeños. Salida: conteos por lote. Aceptación: no mezclar Add de una ronda con Wait anterior pendiente.

[Enunciado y espacio](exercises/05-waitgroups/challenge/README.md) · [Pistas opcionales](exercises/05-waitgroups/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿WaitGroup evita carreras durante el trabajo?
2. ¿Por qué Add antes de go?

[Respuestas razonadas](self-assessment/05-waitgroups.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Esperar resultados y esperar terminación son necesidades distintas; algunos workers no producen un valor por cada salida. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[WaitGroup](https://pkg.go.dev/sync#WaitGroup)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
