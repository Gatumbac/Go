# Pistas · Genéricos básicos

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `Contains[T comparable](values []T, target T) bool`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
generic algorithm + constraint
       ↓ choose T
checked operations for T → concrete call
```

3. Evita este error de la lección: Any como permiso para sumar. Escribe una predicción para una frontera antes de programar.
