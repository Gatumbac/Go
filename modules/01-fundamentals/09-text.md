# 09 · Strings, bytes y runes

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Arrays, slices y maps](08-collections.md) · Siguiente: [Structs](../02-design/01-structs.md)

## Objetivos de aprendizaje

- Explicar strings, bytes y runes y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Arrays, slices y maps](08-collections.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un byte no es un carácter visible. Confundir esas unidades rompe validación, truncamiento y procesamiento de texto internacional.

## Explicación conceptual

Un string es una secuencia inmutable de bytes; puede contener bytes que no formen UTF-8 válido. `len(s)` mide bytes, e indexar devuelve un byte. La convención de texto es UTF-8, pero el tipo no lo valida por ti.

`rune` es un alias de int32 usado para code points Unicode; no equivale necesariamente a un símbolo visible. Un grapheme puede tener varios code points, como una letra y un acento combinante. `range` sobre un string decodifica UTF-8 y entrega offset en bytes y rune. Secuencias inválidas producen RuneError y avanzan según reglas de decodificación.

`[]byte(s)` permite trabajar con bytes mutables y `[]rune(s)` con code points. Expresar estas conversiones como «siempre cero coste» sería incorrecto: pueden requerir asignaciones y copias; las optimizaciones dependen del contexto. Para contar runes usa unicode/utf8, para normalización avanzada necesitarías herramientas adicionales que este recorrido no introduce.

Cortar `s[:n]` corta por bytes y puede partir una codificación UTF-8. Tampoco cortar por runes garantiza conservar graphemes. Define tu contrato con precisión: bytes para protocolos, code points para ciertos ejercicios y graphemes para una interfaz humana que lo necesite.

## Modelo mental visual

```text
"ñA" → bytes: C3 B1 41 (len=3)
       runes: U+00F1 U+0041 (count=2)
       offsets from range: 0, 2
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/09-text/main.go):

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ñA"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	for offset, r := range s {
		fmt.Printf("%d %U\n", offset, r)
	}
}
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/09-text
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/09-text
```

**Resultado e interpretación:** Primera línea `3 2`; después `0 U+00F1` y `2 U+0041`. Los offsets son bytes, no índices de runes.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

JavaScript indexa unidades UTF-16; Java char tampoco representa todo Unicode. En C una cadena común es bytes terminados por cero; Go conserva longitud y admite bytes cero.

## Errores comunes

Llamar caracteres a len; asumir ASCII; truncar UTF-8 por bytes; confundir rune y grapheme.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `CountRunes(text string) int` para UTF-8 válido. Entrada: "abc", "ñA", "". Salida: 3,2,0. Restricción: no usar len como conteo de code points. Aceptación: explicar por qué la segunda entrada tiene tres bytes.

[Enunciado y espacio de implementación](exercises/09-text/guided/README.md) · [Pistas progresivas](exercises/09-text/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `ReverseRunes(text string) string` para UTF-8 válido. Entrada: "Goñ". Salida: "ñoG". Restricción: no prometer reversión correcta de graphemes. Aceptación: input vacío, texto multibyte y documentación del límite.

[Enunciado y espacio](exercises/09-text/challenge/README.md) · [Pistas opcionales](exercises/09-text/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Un string garantiza UTF-8 válido?
2. ¿Una rune es un símbolo visual completo?

[Respuestas razonadas](self-assessment/09-text.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un byte no es un carácter visible. Confundir esas unidades rompe validación, truncamiento y procesamiento de texto internacional. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Strings, bytes, runes](https://go.dev/blog/strings) · [UTF-8](https://pkg.go.dev/unicode/utf8)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
