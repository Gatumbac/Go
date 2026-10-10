# Autoevaluación · JSON

[Volver a la lección](../04-json.md)

Responde antes de leer. Las respuestas explican conceptos; no implementan ejercicios.

## 1. ¿Un struct válido sintácticamente garantiza dominio?

No, zero values y rangos deben validarse según contrato.

## 2. ¿Por qué otro Decode para EOF?

El primer valor puede ser válido aunque haya datos adicionales; un solo documento exige consumir y verificar fin.

## Verifica tu razonamiento

Si acertaste por memoria, predice un caso distinto y explica por qué. Registra lectura y comprensión por separado en [PROGRESS](../../../PROGRESS.md).
