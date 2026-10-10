# 05 · Middleware y errores

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [JSON](04-json.md) · Siguiente: [Context y timeouts](06-timeouts.md)

## Objetivos de aprendizaje

- Explicar middleware y errores y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [JSON](04-json.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

Varias rutas comparten observación o políticas; no necesitas repetirlas ni ocultar control de flujo.

## Explicación conceptual

Un middleware recibe un handler y devuelve otro. Puede hacer trabajo antes, después o detener la cadena. El orden de composición determina qué ve cada capa: logging externo observa respuestas de autenticación interna; el orden inverso puede omitirlas. Dibuja el recorrido.

ResponseWriter tiene capacidades opcionales como Flusher o Hijacker; envolverlo solo con Write puede perderlas. Para ejemplos pequeños basta un wrapper limitado, pero documenta ese alcance. No copies un patrón simplificado a streaming sin comprobar interfaces y ResponseController.

No registres tokens, passwords o bodies arbitrarios. Usa IDs y duración con contexto útil. Un error interno se traduce a respuesta segura, manteniendo diagnóstico privado. Recover puede proteger una frontera HTTP deliberada, pero no arregla estado corrupto ni captura panics de otras goroutines.

El middleware del ejemplo agrega un header antes de delegar, sin estado global ni wrappers complicados. Los handlers siguen siendo comprobables en memoria. Un middleware no debe llamar next dos veces salvo un protocolo muy explícito: duplicaría efectos.

## Modelo mental visual

```text
request → outer before → inner before → handler
response ← outer after ← inner after ← body/status
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/05-middleware/main.go):

```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func stamp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Header().Set("X-Course", "Go"); next.ServeHTTP(w, r) })
}
func main() {
	h := stamp(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(204) }))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	fmt.Println(rec.Code, rec.Header().Get("X-Course"))
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/05-middleware
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/05-middleware
```

**Resultado e interpretación:** Imprime `204 Go`. No body acompaña a 204.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Express middleware usa next como continuación; Go suele componer handlers directamente y retornar para cortar.

## Errores comunes

Orden no pensado; logs de secretos; next llamado dos veces; wrapper pierde capacidades de streaming.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Implementa middleware que requiera header X-Client no vacío y responda 400 sin llamar next si falta. Entrada:dos requests con/sin header. Salida:status y contador de next 1. Aceptación: sin autenticación ficticia por ese header.

[Enunciado y espacio de implementación](exercises/05-middleware/guided/README.md) · [Pistas progresivas](exercises/05-middleware/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Crea dos middlewares que escriban marcas en un slice de traza durante una ejecución secuencial de test. Entrada:una petición. Salida:orden before/after documentado. Restricción:no compartir slice entre requests. Aceptación: deducir composición inversa.

[Enunciado y espacio](exercises/05-middleware/challenge/README.md) · [Pistas opcionales](exercises/05-middleware/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿Por qué importa el orden?
2. ¿Un wrapper mínimo conserva streaming?

[Respuestas razonadas](self-assessment/05-middleware.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

Varias rutas comparten observación o políticas; no necesitas repetirlas ni ocultar control de flujo. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Handler](https://pkg.go.dev/net/http#Handler) · [ResponseController](https://pkg.go.dev/net/http#ResponseController)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
