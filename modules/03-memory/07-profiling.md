# 07 · Asignaciones y profiling

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Garbage collection](06-gc.md) · Siguiente: [Go test](../04-tooling/01-testing.md)

## Objetivos de aprendizaje

- Explicar asignaciones y profiling y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Garbage collection](06-gc.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Optimizar sin evidencia suele mover complejidad y no resolver el cuello de botella.

## Explicación conceptual

Profiling atribuye muestras o eventos a funciones y líneas. Un perfil CPU muestrea actividad cuando el proceso usa CPU; no mide directamente toda latencia de red. Un perfil heap permite estudiar asignación o memoria retenida según la vista elegida. Block y mutex profiles ayudan con esperas y contención, pero tienen sus propios costes y configuración.

Para comenzar usa benchmarks con -benchmem. B/op y allocs/op describen el escenario medido, no necesariamente el request real. Captura perfiles en archivos y analiza con `go tool pprof`; su vista top muestra dónde se concentra una métrica, y list permite inspeccionar funciones.

Una optimización válida conserva resultados, mejora una métrica relevante y no empeora riesgos sin justificación. Preasignar capacidad puede reducir crecimiento si conoces un tamaño razonable; preasignar cantidades enormes puede aumentar memoria viva. Reusar buffers cambia ownership y puede causar datos corruptos si se comparten.

Mantén resultado visible al benchmark para evitar eliminación de trabajo. Usa varias mediciones, mismo hardware y configuración. No publiques el endpoint pprof sin control en una API pública: los perfiles pueden revelar información operativa.

## Modelo mental visual

```text
representative workload → measurement → profile
                              ↓
                         hypothesis → change → compare
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/07-profiling/main.go):

```go
package main

import (
	"fmt"
	"testing"
)

var sink []int

func main() {
	n := testing.AllocsPerRun(20, func() { sink = make([]int, 100) })
	fmt.Printf("allocations/run=%.1f\n", n)
}
```

Desde la raíz:

```sh
go run ./modules/03-memory/examples/07-profiling
go build -o /tmp/go-from-zero-example ./modules/03-memory/examples/07-profiling
```

**Resultado e interpretación:** Normalmente una asignación por llamada en este escenario, pero es una medición del toolchain. El sink conserva el resultado.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Node.js y Java también tienen profilers; cambiar de lenguaje no reemplaza definir una carga y una métrica representativas.

## Errores comunes

CPU profile para latencia de red; medir cargas diferentes; optimizar solo allocs/op y perder legibilidad.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe un programa que compare AllocsPerRun al construir un slice de tamaño 100 con y sin capacidad inicial. Entrada: mismo contenido. Salida: dos cifras y comprobación de igualdad. Aceptación: no imponer un ganador sin medir.

[Enunciado y espacio de implementación](exercises/07-profiling/guided/README.md) · [Pistas progresivas](exercises/07-profiling/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Después de benchmarks, perfila ese experimento. Entrada: benchmark con resultado usado. Salida: informe de top y una hipótesis. Aceptación: distinguir bytes asignados de memoria retenida.

[Enunciado y espacio](exercises/07-profiling/challenge/README.md) · [Pistas opcionales](exercises/07-profiling/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿CPU profile explica todo timeout?
2. ¿Menos asignaciones siempre mejora?

[Respuestas razonadas](self-assessment/07-profiling.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Optimizar sin evidencia suele mover complejidad y no resolver el cuello de botella. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Diagnostics](https://go.dev/doc/diagnostics) · [Pprof](https://pkg.go.dev/runtime/pprof)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
