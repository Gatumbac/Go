# Pistas · Composición

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Modela `Document` que usa una interfaz `Formatter` con `Format(string) string`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
Report
 ├─ metadata (value)
 └─ formatter (collaborator)
explicit delegation → controlled public behavior
```

3. Evita este error de la lección: Tratar embedding como extends. Escribe una predicción para una frontera antes de programar.
