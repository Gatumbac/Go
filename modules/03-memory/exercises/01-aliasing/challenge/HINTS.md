# Pistas · Valores, referencias y aliasing

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Diseña `Snapshot` de un struct con []string y explica su contrato.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
slice A ─┐
         ├──> [shared elements]
slice B ─┘
copy(elements) → [different elements]
```

3. Evita este error de la lección: Copiar solo el descriptor. Escribe una predicción para una frontera antes de programar.
