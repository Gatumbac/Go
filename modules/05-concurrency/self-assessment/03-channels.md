# Autoevaluación · Channels

[Volver a la lección](../03-channels.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Por qué bloquea un channel sin buffer?

La comunicación necesita contraparte; no almacena un elemento pendiente como buffer.

## 2. ¿Close significa que la recepción falla de inmediato?

No, primero se drenan valores pendientes; luego se recibe zero value y ok=false.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
