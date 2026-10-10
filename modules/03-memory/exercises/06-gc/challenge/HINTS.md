# Pistas · Garbage collection

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea un experimento que lea MemStats antes y después de crear una colección temporal.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
roots → reachable graph → keep
         no path to object → eligible for reclamation
resource Close ≠ immediate heap reclamation
```

3. Evita este error de la lección: GC=Close. Escribe una predicción para una frontera antes de programar.
