# Pistas · Métodos

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Define `Balance` con Amount int y método `Deposit(n int)` para n>=0.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
value receiver → copy → local changes
pointer receiver → address copy → shared changes
```

3. Evita este error de la lección: Value receiver esperando mutación persistente. Escribe una predicción para una frontera antes de programar.
