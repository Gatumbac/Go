# 03 · Estructura de un programa

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Compilador, binario y runtime](02-compilation.md) · Siguiente: [Packages y modules](04-packages.md)

## Objetivos de aprendizaje

- Explicar estructura de un programa y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Compilador, binario y runtime](02-compilation.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Antes de pensar en APIs necesitas reconocer dónde comienza la ejecución y cómo se conectan nombres, imports y funciones.

## Explicación conceptual

Todo archivo Go comienza con una declaración de package. Los archivos normales de un mismo directorio pertenecen al mismo package. `package main` identifica una aplicación, y `func main()` es su punto de entrada: no recibe argumentos ni devuelve un código de salida como en C.

Los imports vinculan nombres de paquetes utilizados. Un import sin uso es error; también lo es una variable local declarada y no usada. Estas reglas obligan a retirar ruido, pero no prueban calidad de diseño. `fmt.Println` es una función exportada del package fmt; la primera letra mayúscula determina exportación, no un modificador `public`.

Go usa llaves para bloques, pero normalmente omite punto y coma gracias a reglas léxicas de inserción. Por eso la llave de apertura de funciones y control va en la misma línea. `gofmt` organiza el estilo; no necesitas discutir espacios con cada colaborador.

Hay inicialización antes de main: valores de package y funciones init. Deja los efectos laterales fuera de init cuando puedas para hacer la construcción de dependencias visible. Al retornar main, termina el proceso; no espera goroutines sueltas. Aquí aún no usamos concurrencia: la secuencia de llamadas es directa.

## Modelo mental visual

```text
package → imports → declarations
                       |
                       v
                 func main()
                       |
                       v
                 fmt.Println
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/03-program/main.go):

```go
package main

import "fmt"

func greeting(name string) string { return "Welcome, " + name }
func main()                       { fmt.Println(greeting("Ada")) }
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/03-program
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/03-program
```

**Resultado e interpretación:** Imprime `Welcome, Ada`. Primero se evalúa greeting y luego se entrega su resultado a Println.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

En Node.js el código top-level de un módulo se ejecuta al cargarlo. En Go las declaraciones top-level definen elementos; los pasos ordinarios se escriben dentro de funciones.

## Errores comunes

Añadir argumentos a main; usar imports como scripts; pensar que todas las declaraciones se ejecutan en orden como instrucciones.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea `main.go` para imprimir dos líneas: `Learning Go` y `One step at a time`. Entrada: ninguna. Restricción: solo fmt y main. Aceptación: salida exacta con dos saltos de línea y explicación de package e import.

[Enunciado y espacio de implementación](exercises/03-program/guided/README.md) · [Pistas progresivas](exercises/03-program/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Crea un programa que salude a un nombre fijo mediante una función `message(name string) string`. Entrada: "Lin". Salida: `Hello Lin`. Restricción: la función devuelve el texto, no imprime. Aceptación: main controla la salida y puedes separar declaración de llamada.

[Enunciado y espacio](exercises/03-program/challenge/README.md) · [Pistas opcionales](exercises/03-program/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Package main basta para ejecutar?
2. ¿Qué significa exportar en Go?

[Respuestas razonadas](self-assessment/03-program.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Antes de pensar en APIs necesitas reconocer dónde comienza la ejecución y cómo se conectan nombres, imports y funciones. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Program initialization](https://go.dev/ref/spec#Program_initialization_and_execution)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
