# Pistas · Channels

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa productor de números 1..4 y consumidor que sume.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
producer ── send ──> [buffer] ── receive ──> consumer
close: no more sends; drain → zero,false
```

3. Evita este error de la lección: Close desde receptor sin coordinación. Escribe una predicción para una frontera antes de programar.
