# Hito 07 · Pruebas y medición

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Demostrar contratos y límites de la versión que escribiste.

## Requisitos y restricciones

Escribe unit tests, integración HTTP en memoria, tests concurrentes y benchmarks. Benchmarks separan registrar, consultar y procesar. Guarda hardware/toolchain/configuración.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

go test, go vet y -race sin fallos. Tabla de escenarios de todos los hitos, con determinismo por coordinación. Benchmarks con resultados usados y memoria por operación.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

No inventes cobertura de producción. Coverage alto no prueba shutdown ni carreras. Documenta retención de historial en memoria y una política acotada futura.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
