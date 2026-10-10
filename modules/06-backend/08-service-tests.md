# 08 · Servicios y testing HTTP

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Database/sql](07-sql.md) · Siguiente: [Proyecto final](../../final-project/README.md)

## Objetivos de aprendizaje

- Explicar servicios y testing http y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Database/sql](07-sql.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un servicio debe poder probarse por contratos sin red externa ni una arquitectura desproporcionada.

## Explicación conceptual

Separa transporte HTTP de una operación de aplicación y sus colaboradores. El handler parsea y traduce; la operación valida reglas y coordina dependencias. Una interfaz pequeña para almacenamiento o consulta permite un fake determinista. No dupliques toda la API de una DB bajo una interfaz gigantesca.

httptest.NewRecorder prueba el handler en memoria; NewServer agrega una pila HTTP real local. El primero no prueba conexiones ni TLS. Decide cuál necesitas y anota el límite. Tests de integración con una DB real son otra categoría y no se sustituyen con un fake.

Prueba éxito, entrada inválida, ausencia, fallo de dependencia y cancelación. Comprueba status, headers y body relevantes. Un test que solo verifica 200 ignora datos incorrectos. En concurrencia los fakes también deben ser seguros si handlers comparten estado.

Mantén assembly en main o constructor explícito. No conectes DB en init: los tests importarían efectos externos. Errores internos se traducen en la frontera; no publiques secretos o SQL. Este ejemplo consulta un color, un dominio deliberadamente distinto de la cola de trabajos final.

## Modelo mental visual

```text
request → handler adapter → lookup capability
                              real implementation / fake
assert status + body + error translation
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/08-service-tests/main.go):

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Lookup interface {
	Color(context.Context) (string, error)
}
type fixed struct{}

func (fixed) Color(context.Context) (string, error) { return "blue", nil }
func colorHandler(dep Lookup) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, err := dep.Color(r.Context())
		if err != nil {
			http.Error(w, "unavailable", 503)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, name)
	})
}
func main() {
	rec := httptest.NewRecorder()
	colorHandler(fixed{}).ServeHTTP(rec, httptest.NewRequest("GET", "/color", nil))
	fmt.Println(rec.Code, rec.Body.String())
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/08-service-tests
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/08-service-tests
```

**Resultado e interpretación:** Imprime `200 blue`. Las pruebas adicionales cubren fallo sin filtrar su texto interno.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/08-service-tests/main_test.go):

```go
package main

import (
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

type broken struct{}

func (broken) Color(context.Context) (string, error) { return "", errors.New("private detail") }
func TestColor(t *testing.T) {
	for _, tc := range []struct {
		name string
		dep  Lookup
		want int
		body string
	}{{"success", fixed{}, 200, "blue"}, {"failure", broken{}, 503, "unavailable\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			colorHandler(tc.dep).ServeHTTP(rec, httptest.NewRequest("GET", "/color", nil))
			if rec.Code != tc.want || rec.Body.String() != tc.body {
				t.Fatalf("got %d %q", rec.Code, rec.Body.String())
			}
			if strings.Contains(rec.Body.String(), "private") {
				t.Fatal("leaked detail")
			}
		})
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -v ./modules/06-backend/examples/08-service-tests
```

## Comparación con otros lenguajes

Un service layer en Node puede servir igual; el objetivo es una frontera comprobable, no replicar carpetas de un framework.

## Errores comunes

Fake como integración DB; solo status; constructor con efectos ocultos; test que depende de Internet.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Diseña handler GET /books/{id} con una interfaz de búsqueda. Entrada:id positivo, ausente, inválido o fallo. Salida:200,404,400,503 respectivamente. Restricción: fake, no DB. Aceptación: body de éxito correcto y mensajes externos seguros.

[Enunciado y espacio de implementación](exercises/08-service-tests/guided/README.md) · [Pistas progresivas](exercises/08-service-tests/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Prueba cancelación y concurrencia del fake de tu servicio. Entrada:ctx cancelado y diez requests. Salida:errores definidos y datos sin carrera. Aceptación: -race limpio y ninguna dependencia externa.

[Enunciado y espacio](exercises/08-service-tests/challenge/README.md) · [Pistas opcionales](exercises/08-service-tests/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Recorder prueba TCP?
2. ¿Qué gana una interfaz pequeña?

[Respuestas razonadas](self-assessment/08-service-tests.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un servicio debe poder probarse por contratos sin red externa ni una arquitectura desproporcionada. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Httptest](https://pkg.go.dev/net/http/httptest) · [Module layout](https://go.dev/doc/modules/layout)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
