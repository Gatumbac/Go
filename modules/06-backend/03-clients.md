# 03 · Clientes y recursos

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Servidores y routing](02-servers.md) · Siguiente: [JSON](04-json.md)

## Objetivos de aprendizaje

- Explicar clientes y recursos y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Servidores y routing](02-servers.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Llamar otra API exige controlar recursos y fallos de transporte separados de los de aplicación.

## Explicación conceptual

http.Client usa un Transport para conexiones y puede reutilizarse entre llamadas concurrentes. Crear un cliente/transport por request puede desperdiciar pools. Ajusta timeout y política de redirects según contrato; un timeout cero no impone límite total.

Do devuelve error para problemas de transporte/política, no automáticamente para status no 2xx. Comprueba status y luego interpreta body según el API. Body es un stream que debe cerrarse tras una respuesta válida. Leer hasta EOF y cerrar facilita reutilización, con límites para no drenar datos ilimitados de un peer no confiable.

NewRequestWithContext conecta cancelación al request. El contexto puede cancelar una actividad de red, pero no garantiza que el servidor no procesara un efecto antes de que perdieras la respuesta. Reintenta solo si la semántica e idempotencia lo permiten.

Para tests, un RoundTripper falso permite respuestas controladas sin red. Un httptest.Server prueba también transporte local cuando el entorno permite sockets. El ejemplo tiene un transporte en memoria deliberado: demuestra ciclo de Do/Body/Status sin afirmar que probó TCP o TLS.

## Modelo mental visual

```text
Client.Do → Transport → response
     transport error OR status/body
     caller closes response body
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/03-clients/main.go):

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type memoryTransport struct{}

func (memoryTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Header: make(http.Header), Body: io.NopCloser(strings.NewReader("remote ok")), Request: r}, nil
}
func main() {
	client := &http.Client{Transport: memoryTransport{}, Timeout: time.Second}
	resp, err := client.Get("https://example.test/data")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1024))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode, string(body))
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/03-clients
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/03-clients
```

**Resultado e interpretación:** Imprime `200 remote ok`. No realiza una petición externa.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Fetch devuelve respuestas 4xx sin rechazar automáticamente por status; el mismo cuidado conceptual aplica. Go exige cerrar Body explícitamente.

## Errores comunes

No cerrar Body; default client sin límites en producción; retry de cualquier POST; transport fake como prueba de TLS.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa función `FetchText(ctx context.Context,client *http.Client,url string)(string,error)`, acepta solo status 200 y máximo 1024 bytes. Entrada:transport fake con body ok o 503. Salida:"ok",nil o error. Aceptación: Body cerrado y rechazo de cuerpo mayor usando lectura de límite+1.

[Enunciado y espacio de implementación](exercises/03-clients/guided/README.md) · [Pistas progresivas](exercises/03-clients/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Prueba tu cliente con un transporte que registre cierre y devuelva error o respuesta. Entrada:tres escenarios. Salida:assertions de estado y liberación. Restricción: sin Internet. Aceptación: diferencia claramente status y error de transporte.

[Enunciado y espacio](exercises/03-clients/challenge/README.md) · [Pistas opcionales](exercises/03-clients/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Por qué reutilizar Client?
2. ¿Cerrar implica que el efecto remoto no ocurrió?

[Respuestas razonadas](self-assessment/03-clients.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Llamar otra API exige controlar recursos y fallos de transporte separados de los de aplicación. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Client](https://pkg.go.dev/net/http#Client)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
