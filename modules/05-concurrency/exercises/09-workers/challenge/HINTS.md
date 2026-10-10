# Pistas · Worker pools y shutdown

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Extiende tu pool con ctx y prueba consumidor que cancela antes de consumir todo.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
producer → bounded jobs → worker 1 ─┐
                          worker 2 ─┼→ results → consumer
coordinator: join workers → close results
```

3. Evita este error de la lección: Cerrar results antes del join. Escribe una predicción para una frontera antes de programar.
