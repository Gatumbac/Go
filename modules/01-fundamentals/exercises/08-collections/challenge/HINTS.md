# Pistas · Arrays, slices y maps

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `Frequencies(values []string) map[string]int`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
array: [10 20 30 40]
              ↑        ↑
slice: pointer + len + cap
map: key ─────> value (no iteration order contract)
```

3. Evita este error de la lección: Olvidar asignar append. Escribe una predicción para una frontera antes de programar.
