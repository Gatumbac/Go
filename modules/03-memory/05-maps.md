# 05 · Maps

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Slices y backing arrays](04-slices.md) · Siguiente: [Garbage collection](06-gc.md)

## Objetivos de aprendizaje

- Explicar maps y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Slices y backing arrays](04-slices.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

El mapa ofrece asociación rápida, pero sus reglas observables no incluyen orden ni coordinación concurrente.

## Explicación conceptual

El map se comporta como una asociación clave-valor mutable. Copiar el valor map permite acceder a la misma asociación; no duplica entradas. Nil map puede leerse y borrarse, pero no escribirse. Make crea una asociación inicializada; una capacidad sugerida no es límite de tamaño.

Las claves requieren comparabilidad. Un slice no sirve como clave; un array comparable sí puede servir. Con interfaces como claves, el tipo dinámico también debe poder compararse: una entrada que contiene un slice causa panic. Para límites claros suele convenir una clave concreta.

La especificación deja sin definir el orden de iteración. Si necesitas salida estable, extrae y ordena claves. Las entradas no son direccionables: no puedes obtener &m[key]; para modificar campos de un struct almacenado por valor, actualiza una copia y vuelve a asignarla, o elige pointers conscientemente.

El runtime puede cambiar hashing y layout entre versiones. No enseñes buckets como estructura eterna. El coste promedio esperado de lookup no justifica afirmar O(1) en todas las circunstancias ni eliminar medición. En concurrencia, protege la asociación y sus invariantes, no solo llamadas aisladas si el contrato exige leer-modificar-escribir.

## Modelo mental visual

```text
map value A ─┐
             ├─> association { key: value }
map value B ─┘
ordered output = collect keys + sort + lookup
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/05-maps/main.go):

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"b": 2, "a": 1}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
```

Desde la raíz:

```sh
go run ./modules/03-memory/examples/05-maps
go build -o /tmp/go-from-zero-example ./modules/03-memory/examples/05-maps
```

**Resultado e interpretación:** Imprime `a 1` y `b 2`, orden impuesto por el programa, no por map.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Los objetos JS y Map tienen reglas de orden distintas. No transportes esas expectativas a map Go.

## Errores comunes

Orden estable asumido; &m[key]; map con claves any que contienen slices; locks separados alrededor de un incremento lógico.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `SortedKeys(values map[string]int) []string`. Entrada: {z:1,a:2}. Salida: [a,z]. Restricción: no modificar map. Aceptación: nil produce longitud 0 y orden lexicográfico.

[Enunciado y espacio de implementación](exercises/05-maps/guided/README.md) · [Pistas progresivas](exercises/05-maps/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa `MergeCounts(left,right map[string]int) map[string]int` sumando claves compartidas. Entrada: {a:1},{a:2,b:1}. Salida: {a:3,b:1}. Restricción: no modificar inputs. Aceptación: nil y mapas vacíos.

[Enunciado y espacio](exercises/05-maps/challenge/README.md) · [Pistas opcionales](exercises/05-maps/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Copiar un map crea otro conjunto de entradas?
2. ¿Map tiene orden aleatorio garantizado?

[Respuestas razonadas](self-assessment/05-maps.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

El mapa ofrece asociación rápida, pero sus reglas observables no incluyen orden ni coordinación concurrente. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Maps specification](https://go.dev/ref/spec#Map_types)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
