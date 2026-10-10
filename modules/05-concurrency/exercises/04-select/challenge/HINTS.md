# Pistas · Select

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Diseña un consumidor que espera datos o señal done y termina al cerrar cualquiera según contrato.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
wait
 ├─ data ready → process
 ├─ cancel ready → return
 └─ neither → block (unless default)
```

3. Evita este error de la lección: Default en busy loop. Escribe una predicción para una frontera antes de programar.
