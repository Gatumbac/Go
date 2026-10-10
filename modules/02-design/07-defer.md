# 07 · Defer, panic y recover

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Errores y wrapping](06-errors.md) · Siguiente: [Organización e idiomas](08-organization.md)

## Objetivos de aprendizaje

- Explicar defer, panic y recover y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Errores y wrapping](06-errors.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Los recursos deben liberarse aun cuando una función retorna por una rama de error.

## Explicación conceptual

Defer registra una llamada para ejecutarse cuando la función que lo contiene retorna; múltiples llamadas se ejecutan en orden inverso. Los argumentos y el receiver de la llamada se evalúan al registrar defer. Una closure diferida puede en cambio leer variables cuando se ejecuta: distingue esos dos momentos.

Defer simplifica pairing de adquisición y liberación, como Lock/Unlock o Open/Close. Regístralo después de comprobar que adquiriste el recurso. Dentro de un loop, las llamadas quedan pendientes hasta el retorno de la función, no hasta el fin de la iteración; extraer una función por iteración puede liberar antes.

Panic interrumpe el camino ordinario y desenrolla la pila de la goroutine ejecutando defers. Recover puede detener ese desenrollado cuando se llama directamente dentro de una función diferida en esa goroutine. No captura automáticamente panics de otras goroutines ni debe ocultar todos los bugs.

Usa error para fallos esperables. Recover se reserva para fronteras deliberadas que pueden contener una falla y definir qué estado queda. Un retorno normal tras recover no prueba que la operación haya cumplido su contrato. Errores de Close también importan según recurso: cerrar un escritor puede ser el momento en que falla el flush.

## Modelo mental visual

```text
defer A → defer B → return
                     ↓
                 execute B
                 execute A
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/07-defer/main.go):

```go
package main

import "fmt"

func main() {
	n := 1
	defer fmt.Println("snapshot", n)
	defer func() { fmt.Println("later", n) }()
	n = 2
	fmt.Println("body", n)
}
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/07-defer
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/07-defer
```

**Resultado e interpretación:** Salida: `body 2`, `later 2`, `snapshot 1`. Las llamadas defer son LIFO y sus argumentos se capturaron antes.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

try/finally libera al terminar un bloque try; defer se asocia al retorno de la función. No son idénticos cuando hay loops.

## Errores comunes

Defer antes de comprobar error; acumular archivos abiertos en un loop; recover que silencia corrupción de estado.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa una función `Trace() []string` que mediante defers devuelva secuencia ["body","second","first"]. Entrada: ninguna. Restricción: usar resultado nombrado y defers, sin prints. Aceptación: explica cuándo se evalúan expresiones y cuándo cambia el slice.

[Enunciado y espacio de implementación](exercises/07-defer/guided/README.md) · [Pistas progresivas](exercises/07-defer/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe un programa que abra y cierre tres archivos temporales uno por uno mediante una función auxiliar por archivo. Entrada: contenido fijo. Salida: tres tamaños correctos. Restricción: errores explícitos, limpieza de temporales. Aceptación: ningún recurso queda pendiente hasta el final del loop exterior.

[Enunciado y espacio](exercises/07-defer/challenge/README.md) · [Pistas opcionales](exercises/07-defer/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Defer se ejecuta al salir de un bloque?
2. ¿Recover en main captura panic de otro worker?

[Respuestas razonadas](self-assessment/07-defer.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Los recursos deben liberarse aun cuando una función retorna por una rama de error. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Defer, panic, recover](https://go.dev/blog/defer-panic-and-recover)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
