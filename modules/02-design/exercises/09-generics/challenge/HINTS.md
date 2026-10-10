# Pistas · Genéricos básicos

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Define un tipo `Queue[T any]` solo para operaciones secuenciales Push y Pop con bool de presencia.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
generic algorithm + constraint
       ↓ choose T
checked operations for T → concrete call
```

3. Evita este error de la lección: Any como permiso para sumar. Escribe una predicción para una frontera antes de programar.
