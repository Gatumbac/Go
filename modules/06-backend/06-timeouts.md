# 06 · Context y timeouts

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Middleware y errores](05-middleware.md) · Siguiente: [Database/sql](07-sql.md)

## Objetivos de aprendizaje

- Explicar context y timeouts y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Middleware y errores](05-middleware.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Un servidor necesita límites de espera y un camino de cierre, sin confundir cancelación local con rollback remoto.

## Explicación conceptual

El contexto de Request se cancela bajo condiciones documentadas, como desconexión o retorno de ServeHTTP. Deriva de él para trabajo dependiente del request. Si quieres aceptar un job que vive después, necesitas un ciclo de vida del servicio definido, no conservar el contexto del request por accidente.

ReadHeaderTimeout limita espera de headers; ReadTimeout puede cubrir lectura completa; WriteTimeout y IdleTimeout cumplen otros roles. No son intercambiables ni apropiados con los mismos valores para toda API, especialmente streaming. Define un presupuesto según fases y dependencias.

Server.Shutdown deja de aceptar conexiones y espera algunas actividades dentro del contexto dado; no cancela por sí mismo cada trabajo de negocio ni espera conexiones hijacked. Tras el plazo puede requerirse una política adicional, pero no la ejecutes indiscriminadamente sin entender estados.

Un timeout envía una señal de que ya no esperas; si una tarea ignora ctx sigue viva. En este ejemplo la dependencia simulada escucha cancelación antes de producir resultado. En tests usa ctx ya cancelado para un escenario determinista, no un Sleep diminuto que falla según carga.

## Modelo mental visual

```text
request ctx → child deadline → dependency
service shutdown → stop admission → bounded wait → cancel → join
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/06-timeouts/main.go):

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func slow(w http.ResponseWriter, r *http.Request) {
	select {
	case <-r.Context().Done():
		http.Error(w, "canceled", 408)
	default:
		w.WriteHeader(204)
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
	rec := httptest.NewRecorder()
	slow(rec, req)
	fmt.Println(rec.Code)
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/06-timeouts
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/06-timeouts
```

**Resultado e interpretación:** Imprime `408` en esta demostración. No es una política universal de traducción de errores; un cliente desconectado puede no recibir la respuesta.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

AbortController cancela operaciones compatibles; un timeout en un proxy también puede dejar el servicio remoto trabajando.

## Errores comunes

Request ctx para jobs duraderos; Shutdown=cancelar todos los workers; timeouts idénticos sin presupuesto; sleeps para provocar deadline.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa handler que llame una dependencia `Lookup(ctx)` inyectada; un error de deadline se traduce según contrato a 504 y cancelación a 408 en tests locales. Entrada:dependency fake. Salida:status. Aceptación: propagar el ctx, no Background.

[Enunciado y espacio de implementación](exercises/06-timeouts/guided/README.md) · [Pistas progresivas](exercises/06-timeouts/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Escribe shutdown.md para una API con requests y tareas aceptadas. Entrada:plazo de cinco segundos. Salida:orden de señales, joins y estados al vencer plazo. Restricción:no implementación final. Aceptación: distinguir shutdown HTTP y workers.

[Enunciado y espacio](exercises/06-timeouts/challenge/README.md) · [Pistas opcionales](exercises/06-timeouts/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Shutdown cancela todos los jobs?
2. ¿Por qué no usar request ctx para un job aceptado?

[Respuestas razonadas](self-assessment/06-timeouts.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Un servidor necesita límites de espera y un camino de cierre, sin confundir cancelación local con rollback remoto. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Server](https://pkg.go.dev/net/http#Server) · [Request context](https://pkg.go.dev/net/http#Request.Context)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
