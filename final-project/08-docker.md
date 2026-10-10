# Hito 08 · Docker opcional

[Contrato y recorrido](README.md) · [Progreso](../PROGRESS.md)

## Objetivo

Empaquetar tu aplicación sin convertirlo en requisito de aprendizaje.

## Requisitos y restricciones

Escribe tú un Dockerfile multi-stage que compile tu main y ejecute con usuario no root; define configuración y señal de terminación. Decide base según necesidades de TLS/cgo; no copies un binario incompatible.

Tú escribes código y tests. No hay implementación preparada; biblioteca estándar en el recorrido principal. Conserva contratos de hitos anteriores y evita agregar funcionalidades fuera del objetivo.

## Entradas, salidas y criterios de aceptación

Construcción local y request de prueba; SIGTERM inicia política del hito 6; credenciales ausentes de imagen; puerto/configuración documentados. No se requiere push ni despliegue.

Cada escenario debe tener evidencia reproducible. Un resultado esperado sin prueba o razonamiento se mantiene pendiente.

## Revisión conceptual

Un contenedor no aporta persistencia, idempotencia ni seguridad completa. Distingue reinicio del proceso de recuperación de cola.

Antes de pasar al siguiente hito, explica tu modelo, una decisión de ownership y un fallo posible. Registra por separado implementación, verificación y comprensión. No hay fechas límite.
