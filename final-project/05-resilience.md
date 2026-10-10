# Hito 05 · Resiliencia y cancelación

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Hacer finitos los intentos y terminar cooperativamente.

## Requisitos y restricciones

Clasifica errores como transitorios o permanentes mediante un contrato explícito. Reintenta solo transitorios, con max attempts incluyendo el inicial. Backoff no debe ocupar un worker esperando con un Sleep no cancelable; define cómo programar y limitar retries.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

Operación controlada falla transitoriamente dos veces y luego éxito → attempts 3. Error permanente → 1. Vencer timeout puede reintentar si se clasificó transitorio; cancelación del usuario → canceled sin nuevos intentos. Cancelar durante backoff termina la espera.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

Documenta garantía cooperativa y evita superponer intentos de un mismo job. Serializa decisión de cancelación frente a éxito y prueba ambas carreras.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
