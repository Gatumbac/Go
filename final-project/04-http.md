# Hito 04 · API HTTP

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Exponer los contratos sin mezclar parsing con procesamiento.

## Requisitos y restricciones

Implementa las tres rutas del README con net/http, limitación de body y traducción de errores. No abras listeners en unit tests: usa httptest.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

GET conocido/desconocido:200/404; POST inválido:400; body excesivo:413; cola llena:503; registro:202 con ID consultable; métodos incorrectos:405.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

Verifica Content-Type y body, no solo status. No retornes errores internos ni mantengas request ctx para la vida del job.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
