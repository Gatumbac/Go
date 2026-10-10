# Autoevaluación · Context y timeouts

[Volver a la lección](../06-timeouts.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Shutdown cancela todos los jobs?

No. Gestiona el servidor HTTP según su contrato; el ciclo de workers requiere coordinación propia.

## 2. ¿Por qué no usar request ctx para un job aceptado?

Puede cancelarse al terminar la respuesta; el job necesita el scope de vida definido por el servicio.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
