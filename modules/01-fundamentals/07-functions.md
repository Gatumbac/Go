# 07 · Funciones y retornos

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Operadores y control de flujo](06-flow.md) · Siguiente: [Arrays, slices y maps](08-collections.md)

## Objetivos de aprendizaje

- Explicar funciones y retornos y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Operadores y control de flujo](06-flow.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una función permite expresar un contrato pequeño y comprobable, y separar cálculos de I/O.

## Explicación conceptual

La firma define parámetros y resultados. Los argumentos se pasan por valor: se copian sus valores, aunque un valor como un slice pueda permitir acceder a almacenamiento compartido. No digas que slices se pasan «por referencia» para explicar eso; profundizaremos en su descriptor.

Los múltiples retornos expresan datos relacionados sin objetos auxiliares. El patrón `(value, error)` comunica fallo explícito; el caller debe decidir qué hacer. Los nombres de resultados pueden mejorar documentación, pero naked returns en funciones largas obligan a recordar demasiado estado.

Las funciones son valores: puedes pasarlas como argumentos, devolverlas y capturar variables en closures. La captura comparte la variable según su scope; no siempre es una fotografía inmutable de su valor. En concurrencia requiere más cuidado. Una función variádica recibe el último parámetro como slice; puedes expandir un slice con `...`.

Separar una función pura de main permite probar resultados sin interceptar terminal. Una función que calcula y además imprime acopla responsabilidades y oculta la composición. Pureza no es una obligación universal, pero mantener I/O en fronteras facilita pruebas y análisis de errores.

## Modelo mental visual

```text
input → small function → value, status
caller → decide how to print / retry / reject
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/07-functions/main.go):

```go
package main

import "fmt"

func bounds(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
func main() { low, high := bounds(9, 3); fmt.Println(low, high) }
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/07-functions
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/07-functions
```

**Resultado e interpretación:** Imprime `3 9`. El caller recibe dos resultados explícitos, no un array accidental.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Python permite desempaquetar tuplas; Go declara múltiples resultados en la firma. TypeScript suele devolver un objeto o tuple type. La función Go es comprobada por sus parámetros y retornos.

## Errores comunes

Ignorar retornos que señalan fallo; funciones que imprimen cuando deberían devolver; closures que dependen de estado mutable sin documentarlo.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `DivideExact(a, b int) (int, bool)` para enteros no negativos: devuelve cociente y true solo si b>0 y división exacta; en otro caso 0,false. Entrada: (12,3),(7,2),(5,0). Salida: (4,true),(0,false),(0,false). Restricción: sin panic ni I/O. Aceptación: documentar el bool.

[Enunciado y espacio de implementación](exercises/07-functions/guided/README.md) · [Pistas progresivas](exercises/07-functions/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `Apply(values []int, transform func(int) int) []int`. Entrada: [1,2] y función que triplica. Salida: [3,6]. Restricción: callback no nil, resultado independiente, no modificar entrada. Aceptación: comprobar input vacío y explicar el paso por valor.

[Enunciado y espacio](exercises/07-functions/challenge/README.md) · [Pistas opcionales](exercises/07-functions/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Pasar un pointer evita paso por valor?
2. ¿Por qué separar cálculo de impresión?

[Respuestas razonadas](self-assessment/07-functions.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una función permite expresar un contrato pequeño y comprobable, y separar cálculos de I/O. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Functions](https://go.dev/ref/spec#Function_types)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
