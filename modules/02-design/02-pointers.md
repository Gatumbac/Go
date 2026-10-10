# 02 · Punteros y valores

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Structs](01-structs.md) · Siguiente: [Métodos](03-methods.md)

## Objetivos de aprendizaje

- Explicar punteros y valores y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Structs](01-structs.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una función necesita modificar un objeto del caller o compartirlo sin copiarlo entero; eso exige entender dirección, valor y vida útil.

## Explicación conceptual

Un pointer `*T` es un valor que identifica una variable de tipo T; `&x` obtiene su dirección y `*p` accede a la variable señalada. El zero value de un pointer es nil y dereferenciar nil causa panic. Comprueba nil cuando sea parte del contrato, no como defensa ritual contra cualquier error de diseño.

Todos los argumentos pasan por valor, incluido el pointer. Copiar un pointer crea otro acceso a la misma variable; reasignar ese pointer local no reasigna el pointer del caller. Modificar `*p` sí modifica la variable compartida. Esta diferencia explica gran parte de los errores al portar código desde Java o C.

Es válido devolver la dirección de una variable local: Go preserva su vida útil según necesidad. No interpretes esa legalidad como una promesa exacta de heap; el compilador decide asignación y optimización. La semántica garantiza que el objeto siga siendo válido mientras se usa.

Go ordinario no permite aritmética de pointers como C. Evita unsafe en el aprendizaje inicial. Elige pointers por semántica, identidad o mutación necesaria, no porque parezcan «siempre más rápidos»: pueden aumentar aliasing, escapes y coste de GC. Go 1.26 amplía new a expresiones, pero &x y new(T) siguen siendo bases útiles.

## Modelo mental visual

```text
caller p ──┐
           ├──> shared integer
callee q ──┘
q = other changes q; *q = value changes pointee
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/02-pointers/main.go):

```go
package main

import "fmt"

func increase(p *int) { *p += 2 }
func main()           { n := 10; increase(&n); fmt.Println(n) }
```

Desde la raíz:

```sh
go run ./modules/02-design/examples/02-pointers
go build -o /tmp/go-from-zero-example ./modules/02-design/examples/02-pointers
```

**Resultado e interpretación:** Imprime `12`; se pasó una copia de la dirección de n.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

En C podrías devolver una dirección local inválida; Go administra la vida útil. En Java compartir objetos es habitual sin escribir *; Go hace visible la elección.

## Errores comunes

Dereferenciar nil; confundir pointer local con objeto; escoger pointer por rendimiento sin medir.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `Reset(value *int)` para un pointer no nil. Entrada: variable con 7. Salida: variable con 0. Restricción: sin globales. Aceptación: el caller observa el cambio y explicas qué valor se copió.

[Enunciado y espacio de implementación](exercises/02-pointers/guided/README.md) · [Pistas progresivas](exercises/02-pointers/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `CopyValue(value *int) (int, bool)`: nil produce 0,false; otro pointer produce su valor,true sin cambiarlo. Entrada: nil y pointer a 9. Salida: (0,false),(9,true). Aceptación: ausencia diferenciada de un cero legítimo.

[Enunciado y espacio](exercises/02-pointers/challenge/README.md) · [Pistas opcionales](exercises/02-pointers/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Reasignar p cambia el pointer del caller?
2. ¿Devolver &local produce dangling pointer?

[Respuestas razonadas](self-assessment/02-pointers.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una función necesita modificar un objeto del caller o compartirlo sin copiarlo entero; eso exige entender dirección, valor y vida útil. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Pointers](https://go.dev/ref/spec#Pointer_types) · [Go 1.26](https://go.dev/doc/go1.26)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
