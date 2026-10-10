# 01 · Filosofía e historia

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Método](../../docs/study-method.md) · Siguiente: [Compilador, binario y runtime](02-compilation.md)

## Objetivos de aprendizaje

- Explicar filosofía e historia y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Método](../../docs/study-method.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un backend necesita coordinar I/O, fallos y cambios sin que cada equipo invente un ecosistema entero. Go busca reducir complejidad de construcción y mantenimiento, no eliminar el diseño.

## Explicación conceptual

Go nació en Google: Robert Griesemer, Rob Pike y Ken Thompson iniciaron su diseño en 2007 y se anunció públicamente en 2009. Go 1 apareció en 2012 con un compromiso de compatibilidad. Su sintaxis pequeña, compilación, biblioteca estándar y herramientas integradas responden a problemas de ingeniería de equipos grandes.

Es estáticamente tipado: el compilador comprueba tipos antes de ejecutar. Es compilado, pero incluye runtime y garbage collection. Esas propiedades no se excluyen. La simplicidad de la sintaxis tampoco hace simples los problemas de concurrencia: ownership, coordinación y errores siguen siendo decisiones de diseño.

Las interfaces implícitas permiten separar consumidores de implementaciones sin jerarquías de clases. Las goroutines facilitan expresar actividades concurrentes; no garantizan que una consulta SQL, una red lenta o un algoritmo ineficiente mejoren por cambiar de lenguaje. Mide el cuello de botella antes de migrar.

Go encaja en servicios, herramientas CLI y sistemas de infraestructura. No es una elección universal: el control exacto de memoria o ecosistemas especializados pueden favorecer otros lenguajes. La biblioteca estándar permite aprender sin escoger primero un framework. `fmt` imprime valores; la sección siguiente profundiza en cómo se construye el programa.

## Modelo mental visual

```text
problema → modelo → código → compilación → medición
                 ↑                         |
                 +------- revisión --------+
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/01-philosophy/main.go):

```go
package main

import "fmt"

func main() {
	fmt.Println("Go: tipos estáticos, herramientas integradas")
}
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/01-philosophy
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/01-philosophy
```

**Resultado e interpretación:** Una línea: `Go: tipos estáticos, herramientas integradas`. Es una demostración de ejecución, no una prueba de rendimiento.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

TypeScript añade tipos a un entorno JavaScript; Go genera código máquina y conserva servicios de runtime. Java también combina compilación y runtime, pero su pipeline habitual usa bytecode y JVM. C ofrece más control manual de recursos; Go reduce esa carga con GC.

## Errores comunes

Confundir lenguaje pequeño con sistemas simples; prometer que goroutines sustituyen decisiones de arquitectura.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe `comparison.md` con una tabla de dos decisiones: API dominada por latencia SQL y CLI de transformación de datos. Entrada: esos dos escenarios. Salida: por escenario, beneficio posible de Go, límite y medición necesaria. No uses benchmarks inventados. Aceptación: separa lenguaje, infraestructura y algoritmo.

[Enunciado y espacio de implementación](exercises/01-philosophy/guided/README.md) · [Pistas progresivas](exercises/01-philosophy/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe `decision.md` para una migración hipotética desde Node.js. Entrada: servicio con Redis, workers y timeouts. Salida: tres hipótesis verificables y un argumento para conservar la implementación actual. Aceptación: ninguna hipótesis afirma que cambiar de lenguaje arregla por sí solo consistencia o latencia.

[Enunciado y espacio](exercises/01-philosophy/challenge/README.md) · [Pistas opcionales](exercises/01-philosophy/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Compilado significa sin runtime?
2. ¿Qué no arregla migrar desde Node.js?

[Respuestas razonadas](self-assessment/01-philosophy.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un backend necesita coordinar I/O, fallos y cambios sin que cada equipo invente un ecosistema entero. Go busca reducir complejidad de construcción y mantenimiento, no eliminar el diseño. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[FAQ](https://go.dev/doc/faq) · [Go 1](https://go.dev/doc/go1)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
