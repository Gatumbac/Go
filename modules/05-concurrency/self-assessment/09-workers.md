# Autoevaluación · Worker pools y shutdown

[Volver a la lección](../09-workers.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Por qué cerrar results después de Wait?

Porque ningún worker podrá enviar de nuevo; el cierre no compite con envíos.

## 2. ¿Qué dos límites necesitas?

Trabajo activo y trabajo pendiente; solo limitar workers no limita una cola ilimitada.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
