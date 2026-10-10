# Pistas · Worker pools y shutdown

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa un pool que cuente runes de un lote finito de strings con máximo dos workers.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
producer → bounded jobs → worker 1 ─┐
                          worker 2 ─┼→ results → consumer
coordinator: join workers → close results
```

3. Evita este error de la lección: Cerrar results antes del join. Escribe una predicción para una frontera antes de programar.
