# 04 · JSON

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Clientes y recursos](03-clients.md) · Siguiente: [Middleware y errores](05-middleware.md)

## Objetivos de aprendizaje

- Explicar json y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Clientes y recursos](03-clients.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

El contrato externo debe transformarse a tipos sin aceptar silenciosamente datos inesperados o cuerpos ilimitados.

## Explicación conceptual

Encoding/json convierte entre bytes JSON y valores Go. Solo campos exportados participan por defecto; tags controlan nombres y opciones. Un campo ausente suele conservar zero value, por lo que pointer o un DTO explícito puede distinguir ausencia de un cero legítimo.

Decoder permite leer un stream; DisallowUnknownFields rechaza campos no conocidos en structs. Un Decode exitoso no garantiza que el stream terminó: podría haber otro valor JSON después. Para un contrato de un único documento, intenta otro Decode y exige EOF. Limita cuerpo antes de decodificar y valida dominio después.

No uses JSON como copia profunda universal: cambia representaciones, puede perder datos o ignorar campos. Number tiene restricciones de precisión cuando se decodifica a interfaces float64; UseNumber puede ayudar si necesitas preservar análisis posterior. No infieras seguridad del parseo correcto.

Encoder puede agregar newline y puede fallar escribiendo. Si ya enviaste status y parte del cuerpo, no puedes reemplazar limpiamente toda respuesta con error. Para respuestas pequeñas, serializar antes de escribir puede facilitar elegir status. Esta lección usa encoding/json estable, no características experimentales de otra versión.

## Modelo mental visual

```text
limited bytes → decode DTO → validate domain → use
                unknown/trailing input → reject by contract
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/04-json/main.go):

```go
package main

import (
	"encoding/json"
	"fmt"
)

type City struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func main() {
	data, err := json.Marshal(City{Name: "Loja", Population: 200})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/04-json
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/04-json
```

**Resultado e interpretación:** Imprime `{"name":"Loja","population":200}`. El ejemplo no declara orden como garantía general para todos los documentos.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

JSON.parse produce datos sin validar el dominio TypeScript. En Go decodificar a struct añade forma tipada, pero las reglas de negocio siguen siendo tuyas.

## Errores comunes

Campos minúsculos; cero confundido con ausente; Decode una vez aceptando trailing JSON; cuerpos enormes.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa `DecodePerson(reader io.Reader)(Person,error)` con Name no vacío y Age 0..130, rechaza campos desconocidos y documentos extra. Entrada:{"name":"Ana","age":20}. Salida:Person válida. Aceptación: inválido, trailing JSON y unknown field con error.

[Enunciado y espacio de implementación](exercises/04-json/guided/README.md) · [Pistas progresivas](exercises/04-json/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Crea handler POST /settings con body limitado a 1024 bytes y campo enabled requerido usando *bool. Entrada:false explícito y ausencia. Salida:éxito y error respectivamente. Aceptación: distinguir false de no enviado.

[Enunciado y espacio](exercises/04-json/challenge/README.md) · [Pistas opcionales](exercises/04-json/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Un struct válido sintácticamente garantiza dominio?
2. ¿Por qué otro Decode para EOF?

[Respuestas razonadas](self-assessment/04-json.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

El contrato externo debe transformarse a tipos sin aceptar silenciosamente datos inesperados o cuerpos ilimitados. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[JSON](https://pkg.go.dev/encoding/json)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
