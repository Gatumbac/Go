# Hito 02 · Ejecución secuencial

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Obtener una operación correcta antes de introducir interleavings.

## Requisitos y restricciones

Implementa ejecución de una tarea y almacenamiento local secuencial. Un intento pasa por running y produce result o fallo. Inicialmente no hay workers ni API.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

Texto "ñA" → result 2; vacío → 0; UTF-8 inválido rechazado antes de ejecutar. Tests de lógica puros y ausencia de estados terminales inconsistentes.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

¿Qué cambia cuando aparecen dos callers? Identifica dónde existe mutación antes de agregar locks.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
