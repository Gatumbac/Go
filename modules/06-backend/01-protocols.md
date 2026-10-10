# 01 · TCP y HTTP

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Worker pools y shutdown](../05-concurrency/09-workers.md) · Siguiente: [Servidores y routing](02-servers.md)

## Objetivos de aprendizaje

- Explicar tcp y http y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Worker pools y shutdown](../05-concurrency/09-workers.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Una API depende de protocolos con reglas de framing y fallos; una conexión no equivale a una petición exitosa.

## Explicación conceptual

TCP ofrece un stream ordenado de bytes entre extremos, no mensajes de aplicación. Una lectura puede devolver menos bytes de los que enviaste en una escritura. Define framing: longitud, delimitador o protocolo como HTTP. El orden del stream tampoco prueba que la aplicación remota haya aplicado una operación.

HTTP separa método, target, headers y body, con respuestas que incluyen status. Semántica de GET, POST o DELETE importa para caches y reintentos. Un status 500 es una respuesta HTTP válida: el cliente puede no tener error de transporte y aun así la operación falló.

Una conexión puede reutilizarse para varias operaciones según protocolo. Latencia incluye DNS, conexión, TLS, envío y espera de aplicación; cambiar de lenguaje no elimina esos costes. Timeouts deben cubrir fases relevantes. No supongas que todo HTTP moderno envía texto exactamente como un ejemplo HTTP/1.1: HTTP/2 y HTTP/3 tienen otro framing.

TLS aporta confidencialidad y autenticación del extremo según certificados; no reemplaza autorización de negocio. Define qué significa éxito de una petición y qué datos puedes reintentar sin duplicar efectos. El ejemplo parsea HTTP/1.1 de memoria para estudiar estructura sin sockets.

## Modelo mental visual

```text
client → transport bytes → HTTP framing → handler
client ← status + headers + body ← application decision
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/01-protocols/main.go):

```go
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	raw := "GET /health HTTP/1.1\r\nHost: example.test\r\n\r\n"
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(raw)))
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	fmt.Println(req.Method, req.URL.Path, req.Host)
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/01-protocols
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/01-protocols
```

**Resultado e interpretación:** Imprime `GET /health example.test`. panic solo expresa que el fixture constante sería un bug del ejemplo, no una política para entradas externas.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Node http también expone stream y framing; Express abstrae parte, pero no borra semántica de red y reintentos.

## Errores comunes

Una Write=una Read; 500 como error de transporte; TLS como autorización; retry de POST sin idempotencia.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe un programa que parsee un request HTTP/1.1 fijo POST /notes con body abc y Content-Length 3. Entrada:fixture en memoria. Salida:método,path y body. Restricción: cerrar Body. Aceptación: comprobar error antes de acceder.

[Enunciado y espacio de implementación](exercises/01-protocols/guided/README.md) · [Pistas progresivas](exercises/01-protocols/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe retry-notes.md para caída de conexión después de enviar una compra. Entrada:resultado remoto desconocido. Salida:por qué un retry puede duplicar efecto y una estrategia conceptual. Aceptación: distinguir timeout de confirmación de no ejecución.

[Enunciado y espacio](exercises/01-protocols/challenge/README.md) · [Pistas opcionales](exercises/01-protocols/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿TCP preserva mensajes de aplicación?
2. ¿Status 500 implica error de transporte en Client.Do?

[Respuestas razonadas](self-assessment/01-protocols.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Una API depende de protocolos con reglas de framing y fallos; una conexión no equivale a una petición exitosa. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Net](https://pkg.go.dev/net) · [HTTP](https://pkg.go.dev/net/http)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
