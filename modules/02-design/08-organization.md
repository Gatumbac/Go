# 08 · Organización e idiomas

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Defer, panic y recover](07-defer.md) · Siguiente: [Genéricos básicos](09-generics.md)

## Objetivos de aprendizaje

- Explicar organización e idiomas y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Defer, panic y recover](07-defer.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una carpeta por capa no garantiza un buen diseño; las dependencias y contratos sí determinan cuánto cuesta cambiarlo.

## Explicación conceptual

Organiza por responsabilidades comprensibles y evita packages genéricos llamados utils que crecen sin una frontera. El consumidor necesita una API pequeña, no todas las funciones exportadas. Un package no debería requerir conocer sus detalles para usarlo correctamente.

El dominio modela reglas; el transporte interpreta HTTP; una dependencia de almacenamiento persiste. Esa separación no obliga a introducir diez capas. Empieza con estructura simple y extrae cuando aparece una frontera concreta. `internal` limita imports desde fuera del árbol permitido; no es una palabra reservada del source.

Inyectar dependencias por parámetros o constructors hace visible qué requiere una operación. Variables globales mutables o init con I/O dificultan tests, configuración y shutdown. Para errores, conserva causas relevantes hasta la frontera que decide cómo responder.

El Go idiomático favorece nombres claros, early returns para fallos y interfaces pequeñas. No copies automáticamente patrones de Java, contenedores de DI o jerarquías abstractas. Un diseño idiomático debe simplificar el contrato real, no demostrar cuántos patrones conoces.

## Modelo mental visual

```text
HTTP adapter → application → domain
                    ↓
               storage capability
(no reverse import from domain to HTTP)
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/08-organization/main.go):

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func announce(w io.Writer, text string) error { _, err := fmt.Fprintln(w, text); return err }
func main() {
	if err := announce(os.Stdout, "ready"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/08-organization
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/08-organization
```

**Resultado e interpretación:** Imprime `ready`. announce acepta una capacidad estándar, no depende de os.Stdout globalmente.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Express facilita middleware y handlers, pero las reglas pueden existir fuera del framework. En Go evita recrear un contenedor IoC solo por familiaridad.

## Errores comunes

Packages utils omnipresentes; exports innecesarios; imports cíclicos; init que conecta a una DB.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe una herramienta que calcule el precio con descuento en package `pricing` y muestre resultado desde `cmd`. Entrada: 1000 cents y 10%. Salida: 900. Restricción: pricing sin fmt/os y dominio 0..100 para porcentaje. Aceptación: dependencia solo de cmd hacia pricing.

[Enunciado y espacio de implementación](exercises/08-organization/guided/README.md) · [Pistas progresivas](exercises/08-organization/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Dibuja en architecture.md un grafo para una API de biblioteca con catálogo y almacenamiento. Entrada: esos dos conceptos. Salida: imports propuestos y una interfaz de consumidor mínima. Restricción: no implementar API. Aceptación: sin ciclos y explicar qué no necesita interfaz.

[Enunciado y espacio](exercises/08-organization/challenge/README.md) · [Pistas opcionales](exercises/08-organization/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Por qué inyectar io.Writer?
2. ¿Más capas siempre facilitan cambios?

[Respuestas razonadas](self-assessment/08-organization.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una carpeta por capa no garantiza un buen diseño; las dependencias y contratos sí determinan cuánto cuesta cambiarlo. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Effective Go](https://go.dev/doc/effective_go) · [Layout](https://go.dev/doc/modules/layout)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
