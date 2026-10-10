# Autoevaluación · WaitGroups

[Volver a la lección](../05-waitgroups.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿WaitGroup evita carreras durante el trabajo?

No. Espera terminación, pero las escrituras concurrentes aún requieren coordinación o independencia.

## 2. ¿Por qué Add antes de go?

Para que Wait vea el trabajo registrado antes de que pueda concluir que no hay tareas pendientes.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
