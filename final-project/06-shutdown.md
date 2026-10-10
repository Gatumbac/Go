# Hito 06 · Graceful shutdown

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Terminar con política observable y espera acotada.

## Requisitos y restricciones

Cierra admisión primero. Dentro del presupuesto drena aceptados, incluidos retries permitidos; al vencer solicita cancelación de restantes y espera goroutines cooperativas. Coordina cierre HTTP y workers como vidas distintas.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

Shutdown sin trabajo termina; uno con trabajo corto lo completa; presupuesto vencido marca restantes canceled y no reintenta. Registro posterior devuelve 503 o no hay listener. No quedan workers vivos en tests.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

No prometas detener trabajo que ignore ctx. Usa dependencia controlada para probar join y cleanup; distingue señales de confirmación.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
