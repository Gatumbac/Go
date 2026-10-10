# 05 · Gofmt y go vet

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Benchmarks](04-benchmarks.md) · Siguiente: [Debugging con Delve](06-delve.md)

## Objetivos de aprendizaje

- Explicar gofmt y go vet y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Benchmarks](04-benchmarks.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una convención automática reduce ruido visual; un analizador detecta patrones sospechosos antes de ejecutar.

## Explicación conceptual

Gofmt produce formato canónico de Go. Formato no es corrección: un programa bien formateado puede tener errores de ownership o negocio. `gofmt -l` lista archivos distintos del formato esperado sin reescribirlos; -w los modifica. En tus ejercicios decide explícitamente cuándo aplicar formato.

Go vet analiza patrones, como formatos incorrectos de Printf, locks copiados o construcciones sospechosas. Puede descubrir bugs sin ejecutar el camino, pero no demuestra ausencia de fallos. Un diagnóstico exige revisar el contrato y la herramienta; no lo suprimas solo para obtener verde.

Build comprueba tipos y sintaxis; test comprueba escenarios; vet busca patrones. Son evidencias complementarias, no intercambiables. Ejecutarlas por package permite localizar problemas. `go fix` sí transforma source; no forma parte de la revisión no mutante de tus soluciones.

El formato automático del workspace está desactivado para respetar tu edición manual. Los ejemplos del libro se entregan con formato canónico. En un proyecto profesional suele habilitarse formatOnSave como convenio del equipo; aquí la autorización de cambios tiene otra finalidad pedagógica.

## Modelo mental visual

```text
source → gofmt: layout
       → build: types/syntax
       → vet: suspicious patterns
       → tests: behavior scenarios
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/05-format-vet/main.go):

```go
package main

import "fmt"

func main() { name := "Nia"; fmt.Printf("name=%s\n", name) }
```

Desde la raíz:

```sh
go run ./modules/04-tooling/examples/05-format-vet
go build -o /tmp/go-from-zero-example ./modules/04-tooling/examples/05-format-vet
```

**Resultado e interpretación:** Imprime `name=Nia`. Cambiar mentalmente %s a %d permite anticipar qué señalaría vet sin editar el ejemplo.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Verificación específica

Desde la raíz del repo:

```sh
gofmt -l modules/04-tooling/examples/05-format-vet/main.go
go vet ./modules/04-tooling/examples/05-format-vet
```

## Comparación con otros lenguajes

Prettier formatea; ESLint combina reglas variadas. Go separa una herramienta de formato canónico y analizadores de patrones.

## Errores comunes

Vet como prueba completa; -w sin permiso sobre soluciones; formatear para ocultar un diff de lógica.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Crea un programa que imprima un entero y un string mediante Printf. Entrada:7 y "ready". Salida:`7 ready`. Aceptación: build/test/vet sin diagnósticos y explica los verbos elegidos; formatea solo cuando tú decidas.

[Enunciado y espacio de implementación](exercises/05-format-vet/guided/README.md) · [Pistas progresivas](exercises/05-format-vet/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe checks.md clasificando tres fallos: tipo incompatible, formato incorrecto, algoritmo incorrecto. Salida: qué herramienta puede detectarlos y un límite por herramienta. Aceptación: no atribuir corrección semántica a gofmt.

[Enunciado y espacio](exercises/05-format-vet/challenge/README.md) · [Pistas opcionales](exercises/05-format-vet/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Gofmt -l modifica archivos?
2. ¿Vet reemplaza tests?

[Respuestas razonadas](self-assessment/05-format-vet.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una convención automática reduce ruido visual; un analizador detecta patrones sospechosos antes de ejecutar. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Gofmt](https://pkg.go.dev/cmd/gofmt) · [Vet](https://pkg.go.dev/cmd/vet)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
