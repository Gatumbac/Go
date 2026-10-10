# Pistas · Operadores y control de flujo

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `Classify(score int) string`: 0..49 "retry", 50..79 "pass", 80..100 "strong", fuera de 0..100 "invalid".

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
condition? ── true ──> branch A
    |
   false
    v
 branch B → continuation
```

3. Evita este error de la lección: Usar asignación en lugar de comparación. Escribe una predicción para una frontera antes de programar.
