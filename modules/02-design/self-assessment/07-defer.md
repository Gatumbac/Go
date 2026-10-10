# Autoevaluación · Defer, panic y recover

[Volver a la lección](../07-defer.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Defer se ejecuta al salir de un bloque?

No, al retornar la función que lo registró.

## 2. ¿Recover en main captura panic de otro worker?

No; la recuperación pertenece a la goroutine donde ocurre el desenrollado.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
