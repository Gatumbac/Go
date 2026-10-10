# 06 · Debugging con Delve

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Gofmt y go vet](05-format-vet.md) · Siguiente: [Race detector](07-race.md)

## Objetivos de aprendizaje

- Explicar debugging con delve y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Gofmt y go vet](05-format-vet.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Cuando una salida sorprende, necesitas observar el estado antes y después del punto donde cambia.

## Explicación conceptual

Un debugger controla ejecución y permite inspeccionar variables, stack frames y goroutines. Delve está especializado en Go y entiende su runtime. No lo necesitas para leer este curso, pero puede reducir tiempo de investigación cuando prints no bastan.

`dlv debug ./ruta` construye un binario depurable y abre una sesión. Pon un breakpoint por nombre de función con `break main.discount`, usa continue para llegar, args y locals para inspeccionar, next para avanzar sin entrar, step para entrar y stack para ver llamadas. Breakpoints por símbolo resisten mejor cambios de línea.

Las optimizaciones pueden ocultar variables o cambiar pasos observables; Delve suele usar flags de build adecuados para debug. No concluyas que una variable no existe semánticamente porque el debugger no puede mostrarla. Añade una reproducción mínima y una hipótesis antes de recorrer cientos de frames.

En concurrencia, pausar altera timing. Un bug que desaparece al depurar puede depender de coordinación: complementa con race detector y tests controlados. No arregles al azar porque viste un valor incorrecto; encuentra cuándo se desvió del contrato.

## Modelo mental visual

```text
reproduce → hypothesis → breakpoint before mutation
                         ↓ inspect before/after
                    explain cause → verify repair
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/06-delve/main.go):

```go
package main

import "fmt"

func discount(cents int) int { result := cents - cents/10; return result }
func main()                  { fmt.Println(discount(200)) }
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/06-delve
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/06-delve
```

**Resultado e interpretación:** Imprime `180`. Traza esperada: args cents=200, después del cálculo result=180. En Delve usa `break main.discount`, `continue`, `args`, `next`, `locals`, `quit`.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Verificación específica

Desde la raíz del repo:

```sh
dlv debug ./modules/04-tooling/examples/06-delve
```

Este comando es opcional y requiere instalar Delve por tu cuenta. La traza descrita arriba sirve para estudiar sin esa herramienta.

## Comparación con otros lenguajes

Debuggers de VS Code para JS o Java ofrecen pasos similares; goroutines y runtime Go requieren soporte específico.

## Errores comunes

Cambiar código sin hipótesis; exigir una variable optimizada; inferir timing productivo al pausar.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Investiga el bug intencional descrito en docs/debugging.md. Entrada: [-8,-3,-5]. Salida esperada:-3. Restricción: crea tu reproducción, no edites ejemplos normales. Aceptación: traza que identifique el valor inicial incorrecto y tu reparación probada.

[Enunciado y espacio de implementación](exercises/06-delve/guided/README.md) · [Pistas progresivas](exercises/06-delve/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Depura una función tuya que modifica un slice compartido. Entrada: dos alias. Salida: traza de quién cambia el elemento. Aceptación: explicar la causa sin atribuirla al debugger.

[Enunciado y espacio](exercises/06-delve/challenge/README.md) · [Pistas opcionales](exercises/06-delve/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Next y step son equivalentes?
2. ¿Un bug que desaparece al pausar queda resuelto?

[Respuestas razonadas](self-assessment/06-delve.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Cuando una salida sorprende, necesitas observar el estado antes y después del punto donde cambia. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Delve commands](https://github.com/go-delve/delve/tree/master/Documentation/cli)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
