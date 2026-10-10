# 02 · Servidores y routing

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [TCP y HTTP](01-protocols.md) · Siguiente: [Clientes y recursos](03-clients.md)

## Objetivos de aprendizaje

- Explicar servidores y routing y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [TCP y HTTP](01-protocols.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un handler necesita un contrato de rutas, métodos y respuestas que se pueda probar sin desplegar un servicio.

## Explicación conceptual

Un handler implementa ServeHTTP(ResponseWriter,*Request); HandlerFunc adapta una función. ServeMux enruta según patterns. La base Go 1.26 permite patterns con método y variables de path, disponibles desde Go 1.22, como `GET /greet/{name}` y Request.PathValue.

Configura headers antes de escribir status o body. El primer Write puede enviar implícitamente 200; luego intentar cambiar status no funciona como esperas. No escribas dos respuestas tras detectar un error: retorna. La semántica de HEAD y patterns específicos puede afectar rutas; prueba métodos válidos e inválidos.

Un servidor real usa http.Server con timeouts y dirección explícita. ListenAndServe normalmente bloquea; ErrServerClosed tras shutdown esperado no es un fallo de arranque. Para este ejemplo usamos Recorder: ejecuta el handler en memoria, sin aceptar conexiones.

Handlers pueden ejecutarse concurrentemente. El mux no vuelve seguro tu estado mutable. Valida entradas en frontera, limita cuerpos cuando corresponda y no expongas errores internos. Separa el cálculo de reglas de la escritura HTTP para simplificar tests.

## Modelo mental visual

```text
request → ServeMux method/path → Handler → ResponseWriter
                              validation → status/body
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/02-servers/main.go):

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "Hi "+r.PathValue("name"))
	})
	return mux
}
func main() {
	rec := httptest.NewRecorder()
	handler().ServeHTTP(rec, httptest.NewRequest("GET", "/greet/Uma", nil))
	fmt.Println(rec.Code, rec.Body.String())
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/02-servers
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/02-servers
```

**Resultado e interpretación:** Imprime `200 Hi Uma`. El Recorder permite comprobar la misma lógica sin red real.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

### Prueba completa del ejemplo

El test usa el mismo problema didáctico, no resuelve tus ejercicios. [Source del test](examples/02-servers/main_test.go):

```go
package main

import (
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	for _, tc := range []struct {
		method, path string
		want         int
	}{{"GET", "/greet/Uma", 200}, {"POST", "/greet/Uma", 405}, {"GET", "/absent", 404}} {
		t.Run(tc.method+tc.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler().ServeHTTP(rec, httptest.NewRequest(tc.method, tc.path, nil))
			if rec.Code != tc.want {
				t.Fatalf("got %d want %d", rec.Code, tc.want)
			}
		})
	}
}
```

### Verificación específica

Desde la raíz del repo:

```sh
go test -v ./modules/06-backend/examples/02-servers
```

## Comparación con otros lenguajes

Express Router usa middleware y patrones con otras reglas. Aprende las reglas de ServeMux en vez de trasladar paths literalmente.

## Errores comunes

Header después de Write; olvidar return tras error; global map sin lock; asumir un solo handler a la vez.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa handler GET /double/{n}, n entero -100..100. Entrada:/double/4. Salida:200 con body 8; inválido:400. Restricción: biblioteca estándar, sin servidor persistente. Aceptación: tests para método, rango y parseo.

[Enunciado y espacio de implementación](exercises/02-servers/guided/README.md) · [Pistas progresivas](exercises/02-servers/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Implementa GET /health con JSON {"status":"ok"}; otros métodos deben producir 405. Entrada:GET y POST. Salida:200/405. Aceptación: Content-Type correcto antes de body.

[Enunciado y espacio](exercises/02-servers/challenge/README.md) · [Pistas opcionales](exercises/02-servers/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Write sin WriteHeader qué hace?
2. ¿ServeMux protege datos de negocio?

[Respuestas razonadas](self-assessment/02-servers.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un handler necesita un contrato de rutas, métodos y respuestas que se pueda probar sin desplegar un servicio. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[ServeMux](https://pkg.go.dev/net/http#ServeMux)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
