# Pistas · Maps

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `SortedKeys(values map[string]int) []string`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
map value A ─┐
             ├─> association { key: value }
map value B ─┘
ordered output = collect keys + sort + lookup
```

3. Evita este error de la lección: Orden estable asumido. Escribe una predicción para una frontera antes de programar.
