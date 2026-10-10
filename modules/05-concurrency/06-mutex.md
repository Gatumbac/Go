# 06 · Mutex y RWMutex

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [WaitGroups](05-waitgroups.md) · Siguiente: [Carreras, deadlocks y leaks](07-failures.md)

## Objetivos de aprendizaje

- Explicar mutex y rwmutex y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [WaitGroups](05-waitgroups.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una invariante compartida debe cambiar como una unidad, no como accesos sueltos protegidos por separado.

## Explicación conceptual

Mutex ofrece exclusión mutua. Lock entra a la sección crítica y Unlock sale; la sincronización documentada también establece visibilidad entre accesos. Protege la relación entre datos, no «una línea porque parece peligrosa». Un incremento requiere leer y escribir bajo la misma exclusión.

RWMutex permite lectores simultáneos o un escritor exclusivo. Solo sirve si las operaciones de lectura realmente no mutan y la carga lo justifica. No siempre supera Mutex: bookkeeping y contención pueden hacerlo más costoso. No se debe intentar upgrade de RLock a Lock manteniendo lectura.

Mantén la sección crítica pequeña, pero no la dividas rompiendo una invariante. Evita I/O lento bajo el lock cuando puedas obtener una copia y liberar antes. Define un orden fijo si una operación necesita varios locks para reducir deadlocks. Los locks no son reentrantes.

Usa pointer receivers en objetos con Mutex y no copies el lock tras uso. Devolver pointers a datos protegidos puede permitir accesos sin lock; devuelve snapshots o proporciona operaciones. Atomics sirven para contratos específicos, no sustituyen automáticamente invariantes de varios campos.

### Una carrera de negocio sin accesos sueltos desprotegidos

Con stock=1, si A lee bajo lock y lo libera, B puede leer también 1. Después ambos adquieren el lock para decrementar y aceptan la compra. Todos los accesos individuales podrían estar protegidos, pero la decisión «si hay stock, reservar» no lo está como unidad.

| Paso | A | B | Stock |
| --- | --- | --- | --- |
| 1 | Lee 1 y libera. | Espera. | 1 |
| 2 | Decide reservar. | Lee 1 y libera. | 1 |
| 3 | Decrementa bajo lock. | Decide reservar. | 0 |
| 4 | Termina. | Decrementa bajo lock. | -1 |

Race detector puede no denunciar esta secuencia: los accesos de memoria están sincronizados. El test debe comprobar la invariante de stock y admisión, no solo ausencia de data race.

## Modelo mental visual

```text
Lock → read current → update invariant → Unlock
other goroutines wait; then observe ordered state
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/06-mutex/main.go):

```go
package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu sync.Mutex
	n  int
}

func (c *counter) add()       { c.mu.Lock(); defer c.mu.Unlock(); c.n++ }
func (c *counter) value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.n }
func main() {
	var c counter
	var wg sync.WaitGroup
	for range 4 {
		wg.Go(func() {
			for range 100 {
				c.add()
			}
		})
	}
	wg.Wait()
	fmt.Println(c.value())
}
```

Desde la raíz:

```sh
go run ./modules/05-concurrency/examples/06-mutex
go build -o /tmp/go-from-zero-example ./modules/05-concurrency/examples/06-mutex
```

**Resultado e interpretación:** Imprime `400`. El contador se modifica bajo la misma exclusión en todas las tareas.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/06-mutex/main_test.go):

```go
package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c counter
	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 100 {
				c.add()
			}
		})
	}
	wg.Wait()
	if got := c.value(); got != 800 {
		t.Fatalf("got %d want 800", got)
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -race -count=1 ./modules/05-concurrency/examples/06-mutex
```

## Comparación con otros lenguajes

Un lock en Java también debe proteger todas las rutas. En código async JS puede haber invariantes rotas sin memoria simultánea; aquí también hay data races reales.

## Errores comunes

Leer sin lock porque solo se escribe poco; RWMutex en cada struct por rendimiento supuesto; copiar locks; I/O bajo lock sin evaluar.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Define `Inventory` con stock y método `Reserve() bool` protegido: si hay stock decrementa una vez. Entrada:stock 1 y dos llamadas concurrentes. Salida:exactamente una true. Aceptación: -race limpio y nunca negativo.

[Enunciado y espacio de implementación](exercises/06-mutex/guided/README.md) · [Pistas progresivas](exercises/06-mutex/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa snapshot de un mapa protegido que devuelve una copia. Entrada:dos claves. Salida:map independiente. Restricción: no exponer map interno. Aceptación: cambiar resultado no afecta inventario.

[Enunciado y espacio](exercises/06-mutex/challenge/README.md) · [Pistas opcionales](exercises/06-mutex/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Lock por cada acceso basta para stock?
2. ¿RWMutex siempre mejora lecturas?

[Respuestas razonadas](self-assessment/06-mutex.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una invariante compartida debe cambiar como una unidad, no como accesos sueltos protegidos por separado. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Sync](https://pkg.go.dev/sync) · [Memory model](https://go.dev/ref/mem)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
