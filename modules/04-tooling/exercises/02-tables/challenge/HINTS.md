# Pistas · Table-driven tests y subtests

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe tests de `Category(age int) string`: menores de 18 "minor", resto "adult", dominio 0..130.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
table rows → subtest names → independent assertions
empty / boundary / normal / invalid
```

3. Evita este error de la lección: Muchos casos redundantes sin fronteras. Escribe una predicción para una frontera antes de programar.
