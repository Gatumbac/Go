# Pistas · Punteros y valores

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `Reset(value *int)` para un pointer no nil.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
caller p ──┐
           ├──> shared integer
callee q ──┘
q = other changes q; *q = value changes pointee
```

3. Evita este error de la lección: Dereferenciar nil. Escribe una predicción para una frontera antes de programar.
