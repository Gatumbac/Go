# 04 · Packages y modules

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Estructura de un programa](03-program.md) · Siguiente: [Variables, constantes y tipos](05-types.md)

## Objetivos de aprendizaje

- Explicar packages y modules y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Estructura de un programa](03-program.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un archivo es una unidad de edición; necesitas otras unidades para compartir código y manejar dependencias reproducibles.

## Explicación conceptual

Un package reúne archivos de un directorio y ofrece nombres exportados. Un module reúne packages bajo una ruta base y un go.mod. Importas packages, no el archivo .go ni el go.mod. La ruta del module es un prefijo: `example.com/go-from-zero/util` identifica el package util dentro de este module.

El nombre declarado de un package y su import path son conceptos diferentes. Por convención coinciden con el último segmento, pero puedes usar un alias local. Los ciclos de imports están prohibidos: evita que dominio y transporte se necesiten mutuamente. Define dependencias en una dirección y coloca interfaces pequeñas en el consumidor cuando corresponda.

`go.mod` declara module, versión de lenguaje mínima y dependencias. La directiva go no es simplemente una nota: afecta características del lenguaje y requisitos de toolchain. `go.sum` registra hashes; no sustituye a go.mod ni significa que todos los módulos listados se importen directamente.

Este libro usa un solo module principal para reducir fricción. Cada example directory es un package main independiente. El laboratorio SQL usa un module aparte para aislar un driver externo. No necesitas go.work para el recorrido: los workspaces son útiles cuando desarrollas varios modules juntos, pero añadirlos ahora escondería la base.

## Modelo mental visual

```text
module example.com/go-from-zero (go.mod)
  ├─ package main: example A
  ├─ package main: example B
  └─ package reusable: exported names
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/04-packages/main.go):

```go
package main

import (
	"fmt"
	"path"
)

func main() { fmt.Println(path.Join("course", "examples")) }
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/04-packages
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/04-packages
```

**Resultado e interpretación:** Imprime `course/examples`. path es un package estándar; no requiere descargar un module externo.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Un npm package suele ser la unidad publicada y puede exportar varios paths. En Go module y package cumplen responsabilidades separadas; no traduzcas automáticamente cada directorio a una dependencia publicable.

## Errores comunes

Crear go.mod en cada carpeta sin necesidad; importar un path de archivo; aceptar ciclos; confundir GOPATH con ubicación obligatoria.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea un package `labels` dentro de guided y un subdirectorio `cmd` con main. Define `Title(name string) string` que devuelva `Topic: ` seguido del nombre. Entrada: "errors". Salida: `Topic: errors`. Restricción: usa el module existente, sin otro go.mod. Aceptación: main importa el package por su ruta completa y no hay ciclos.

[Enunciado y espacio de implementación](exercises/04-packages/guided/README.md) · [Pistas progresivas](exercises/04-packages/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Diseña y escribe dos packages, `domain` y `display`, dentro de challenge. Entrada: nombre de un producto; salida: representación textual. Restricción: display puede importar domain, domain no importa display. Aceptación: go list muestra dos packages y puedes justificar la dirección.

[Enunciado y espacio](exercises/04-packages/challenge/README.md) · [Pistas opcionales](exercises/04-packages/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Importas un module o un package?
2. ¿Por qué no crear muchos go.mod?

[Respuestas razonadas](self-assessment/04-packages.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un archivo es una unidad de edición; necesitas otras unidades para compartir código y manejar dependencias reproducibles. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Modules reference](https://go.dev/ref/mod) · [Module layout](https://go.dev/doc/modules/layout)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
