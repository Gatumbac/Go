# 05 · Variables, constantes y tipos

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Packages y modules](04-packages.md) · Siguiente: [Operadores y control de flujo](06-flow.md)

## Objetivos de aprendizaje

- Explicar variables, constantes y tipos y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Packages y modules](04-packages.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Los datos externos llegan como bytes; antes de usarlos debes definir qué operaciones y rangos admite tu programa.

## Explicación conceptual

Go tiene tipos básicos, compuestos y definidos por el usuario. `var n int` obtiene el zero value 0. `var ready bool` obtiene false; strings vacíos, pointers nil y structs con campos inicializados recursivamente. No hay variables ordinarias sin inicializar como en C.

`:=` declara e inicializa variables locales y requiere al menos una variable nueva en ese bloque. No se usa al nivel de package. Una asignación con `=` modifica variables existentes. El shadowing crea otra variable en un bloque interior: no es reasignar la exterior.

Los tipos son explícitos incluso con inferencia. Go no convierte implícitamente int a float64 en expresiones ordinarias. Una conversión puede perder información: convertir un float a int trunca hacia cero; convertir entre enteros no valida un rango de negocio. `int` depende de arquitectura; para formatos estables usa tipos de tamaño explícito cuando corresponde.

Las constantes no son variables de solo lectura: sus valores se evalúan como constantes y pueden ser untyped hasta que un contexto exige tipo. Esto permite expresar precisión y comprobar representabilidad. `iota` ayuda a enumerar constantes, pero no verifica automáticamente que un entero externo pertenezca al conjunto. Un tipo definido distinto puede impedir mezclas accidentales; un alias no introduce esa distinción.

## Modelo mental visual

```text
declaration → type + zero value
initializer → inferred or explicit type
conversion → new representation (check information loss)
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/05-types/main.go):

```go
package main

import "fmt"

type Celsius float64

func main() {
	var count int
	const ratio = 1.5
	temp := Celsius(21.5)
	fmt.Println(count, float64(count)+ratio, temp)
}
```

Desde la raíz:

```sh
go run ./modules/01-fundamentals/examples/05-types
go build -o /tmp/go-from-zero-example ./modules/01-fundamentals/examples/05-types
```

**Resultado e interpretación:** Imprime `0 1.5 21.5`. La conversión explícita habilita la suma; Celsius es un tipo definido.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

TypeScript borra la mayoría de tipos al emitir JS. Go mantiene restricciones al compilar y representaciones concretas en ejecución. `const` en JS fija un binding; no equivale a una constante Go.

## Errores comunes

Usar := fuera de funciones; pensar que int siempre tiene 64 bits; usar conversiones como validación; ignorar shadowing.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `Minutes(hours int) int` en package practice. Entrada: 0, 2, 5. Salida: 0, 120, 300. Restricción: dominio 0..1000 y sin I/O. Aceptación: tabla de casos y explicación de var, := y const usada para el factor.

[Enunciado y espacio de implementación](exercises/05-types/guided/README.md) · [Pistas progresivas](exercises/05-types/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Define `type Meters float64` y `type Seconds float64`; escribe `Speed(distance Meters, elapsed Seconds) float64` para elapsed positivo. Entrada: 100, 20. Salida: 5. Restricción: sin aliases. Aceptación: demuestra en tus notas por qué no intercambias argumentos de tipos distintos.

[Enunciado y espacio](exercises/05-types/challenge/README.md) · [Pistas opcionales](exercises/05-types/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Zero value implica dato válido de negocio?
2. ¿Convertir valida?

[Respuestas razonadas](self-assessment/05-types.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Los datos externos llegan como bytes; antes de usarlos debes definir qué operaciones y rangos admite tu programa. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Types](https://go.dev/ref/spec#Types) · [Constants](https://go.dev/ref/spec#Constants)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
