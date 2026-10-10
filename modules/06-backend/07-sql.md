# 07 · Database/sql

[Índice del módulo](README.md) · [Roadmap](../../ROADMAP.md) · Anterior: [Context y timeouts](06-timeouts.md) · Siguiente: [Servicios y testing HTTP](08-service-tests.md)

## Objetivos de aprendizaje

- Explicar database/sql y el problema que resuelve.
- Predecir el comportamiento del ejemplo, incluyendo sus límites.
- Aplicar el contrato del ejercicio y justificar una decisión sin copiar el ejemplo.

## Conocimientos previos

Estudia [Context y timeouts](06-timeouts.md) antes. El recorrido es acumulativo: usa el [glosario](../../GLOSSARY.md) si un término de módulos anteriores no está claro. Para comprobar el entorno consulta [setup](../../docs/setup.md).

## Motivación

El acceso a datos requiere controlar conexiones, transacciones y errores sin mezclar SQL con transporte HTTP.

## Explicación conceptual

database/sql es una interfaz de acceso, no incluye un driver concreto de PostgreSQL. *sql.DB es un pool concurrente, no una conexión individual. Open puede preparar el handle sin verificar conectividad; PingContext comprueba conexión dentro de un plazo. Configura límites de conexiones según capacidad del servidor, no según número de requests.

QueryContext devuelve Rows que debes cerrar y recorrer con Next, Scan y comprobar Rows.Err al final. QueryRowContext difiere: Scan devuelve sql.ErrNoRows si no existe fila. Usa placeholders del driver para valores; concatenar entrada en SQL crea riesgo de inyección. Identificadores y direcciones de orden requieren allowlists, no placeholders de valores.

Una transacción usa BeginTx, operaciones sobre tx y Commit. Un defer Rollback cubre salidas tempranas; tras Commit su resultado esperado ya no equivale a que se deshizo. No mezcles db.Exec con tx.Exec esperando la misma atomicidad. Timeouts y cancelación dependen también del driver y DB.

No mantengas una transacción abierta mientras esperas una llamada HTTP lenta si puedes evitarlo. Define consistencia e idempotencia. El ejemplo principal demuestra clasificación y una función de consulta compilable sin conectar; el laboratorio aislado ofrece ejecución PostgreSQL real con credenciales por entorno.

### Checklist de vida útil de recursos

| Operación | Responsable de cierre | Fallo que aún debes comprobar |
| --- | --- | --- |
| `sql.Open` | Dueño de DB al apagar servicio. | Conectividad mediante PingContext. |
| `QueryContext` | Caller de Rows.Close. | Rows.Err después de Next. |
| `QueryRowContext` | Scan gestiona resultado de esa fila. | ErrNoRows y errores de Scan. |
| `BeginTx` | Caller mediante Commit/Rollback. | Resultado de Commit, no solo de statements. |

No cierres el pool DB después de cada request: el servicio lo reutiliza. Tampoco mantengas Rows sin consumir/cerrar porque monopoliza conexiones y puede hacer que otros requests esperen aunque el servidor tenga CPU disponible.

## Modelo mental visual

```text
sql.DB pool → borrowed connection → query → rows close → reusable
transaction: begin → operations on tx → commit OR rollback
```

El diagrama representa relaciones y decisiones observables, no direcciones reales ni un ABI garantizado. Predice un cambio y dibuja qué parte del modelo cambia.

## Sintaxis y ejemplos

[Programa completo](examples/07-sql/main.go):

```go
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func cityName(ctx context.Context, db *sql.DB, id int) (string, error) {
	var name string
	err := db.QueryRowContext(ctx, "SELECT name FROM cities WHERE id = $1", id).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("city %d: %w", id, err)
	}
	return name, nil
}
func main() {
	err := fmt.Errorf("lookup: %w", sql.ErrNoRows)
	fmt.Println(errors.Is(err, sql.ErrNoRows))
}
```

Desde la raíz:

```sh
go run ./modules/06-backend/examples/07-sql
go build -o /tmp/go-from-zero-example ./modules/06-backend/examples/07-sql
```

**Resultado e interpretación:** Imprime `true`. cityName se compila pero no se ejecuta sin driver y DB; no afirmar conexión validada por este ejemplo.

Antes de ejecutarlo, anota tu predicción. Después señala qué instrucción explica cada resultado y qué propiedad no demuestra esta ejecución.

## Comparación con otros lenguajes

Pools de Node/PostgreSQL tienen responsabilidades semejantes. Un ORM no elimina cierre, transacciones ni presupuestos de conexiones.

## Errores comunes

Open=conectividad; Query sin Close; olvidar Rows.Err; SQL concatenado; operaciones fuera de tx.

Si observas un fallo, conserva el mensaje exacto y una hipótesis. Compilar y ejecutar son verificaciones diferentes; no arregles por ensayo ciego.

## Ejercicio guiado

Escribe query-review.md: entrada, consulta de productos por precio mínimo y sort elegido por usuario. Salida:SQL con placeholder para precio y allowlist de sort. Restricción: no conectar DB. Aceptación: distinguir valor de identificador.

[Enunciado y espacio de implementación](exercises/07-sql/guided/README.md) · [Pistas progresivas](exercises/07-sql/guided/HINTS.md). Escribe tú la implementación; no hay solución preparada.

## Desafío independiente

Haz el laboratorio opcional y después implementa una consulta tuya `FindCity(ctx,id)` con ErrNoRows distinguible. Entrada:id presente y ausente. Salida:nombre o error clasificable. Aceptación: timeout, parametrización y recursos; no guardar credenciales.

[Enunciado y espacio](exercises/07-sql/challenge/README.md) · [Pistas opcionales](exercises/07-sql/challenge/HINTS.md). Resuelve primero sin abrir pistas.

## Preguntas de comprensión

1. ¿DB representa una conexión?
2. ¿Qué comprueba Rows.Err?

[Respuestas razonadas](self-assessment/07-sql.md): consulta después de responder y explica en tus palabras cualquier diferencia. No contienen soluciones de ejercicios.

## Resumen y referencias

El acceso a datos requiere controlar conexiones, transacciones y errores sin mezclar SQL con transporte HTTP. El criterio de comprensión es explicar el mecanismo, predecir el ejemplo y justificar sus límites.

[Database guide](https://go.dev/doc/database/) · [SQL injection](https://go.dev/doc/database/sql-injection)

[Compatibilidad y fuentes](../../docs/sources.md) · [Registrar evidencia](../../PROGRESS.md)
