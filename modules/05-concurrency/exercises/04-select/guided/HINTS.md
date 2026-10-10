# Pistas · Select

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `TryReceive(ch <-chan int) (int,bool)` como recepción no bloqueante para channel abierto.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
wait
 ├─ data ready → process
 ├─ cancel ready → return
 └─ neither → block (unless default)
```

3. Evita este error de la lección: Default en busy loop. Escribe una predicción para una frontera antes de programar.
