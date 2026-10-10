# Hito 01 · Contrato y diseño

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Definir tipos, operaciones, ownership y máquina de estados antes de programar.

## Requisitos y restricciones

Escribe design.md con transiciones legales, punto de admisión, regla de saturación, snapshots e interfaces mínimas para ejecutar tareas. No añadas HTTP todavía.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

Entrada: job válido, inválido y cancelación tardía. Salida: tabla de transición y errores. Cada estado terminal es estable; ninguna regla depende del orden de logs.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

Explica qué protege la consistencia entre registro, cola y cancelación. Distingue una interfaz necesaria de una capa ornamental.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
