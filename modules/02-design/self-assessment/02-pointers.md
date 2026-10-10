# Autoevaluación · Punteros y valores

[Volver a la lección](../02-pointers.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Reasignar p cambia el pointer del caller?

No: p es una copia local. Para alterar el objeto se accede a *p.

## 2. ¿Devolver &local produce dangling pointer?

No en Go ordinario: la vida útil se preserva aunque el objeto deje el scope léxico.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
