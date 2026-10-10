# Autoevaluación · Mutex y RWMutex

[Volver a la lección](../06-mutex.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Lock por cada acceso basta para stock?

No si comprobar y decrementar se separan: la operación lógica debe estar en una sección crítica.

## 2. ¿RWMutex siempre mejora lecturas?

No. Depende de carga, duración y coste de coordinación; mide.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
