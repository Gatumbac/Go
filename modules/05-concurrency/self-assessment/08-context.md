# Autoevaluación · Context y cancelación

[Volver a la lección](../08-context.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Cancel mata una goroutine?

No; cierra una señal que el trabajo debe observar cooperativamente.

## 2. ¿Por qué esperar después de cancel?

Porque la señal no demuestra que la goroutine liberó recursos y retornó.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
