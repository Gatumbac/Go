# Hito 03 · Cola y worker pool

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Limitar trabajo activo y pendiente, con ownership de cierre definido.

## Requisitos y restricciones

Integra queue capacity y workers configurables. Admisión debe ser no bloqueante ante cola llena. Espera trabajadores y evita cerrar canales con productores aún activos.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

Con dependencia bloqueable y workers=2, el máximo activo observado es 2. Llenar cola rechaza otro job; liberar tareas permite continuar. Cada aceptado se procesa sin pérdidas ni duplicación normal.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

Tests coordinados con channels, no Sleep. Asegura snapshots independientes y ejecuta -race.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
