# Pistas · Composición

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Define `Profile` con campo nombrado Address (City string) y Name string. Implementa `Location() string` delegando explícitamente al campo.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
Report
 ├─ metadata (value)
 └─ formatter (collaborator)
explicit delegation → controlled public behavior
```

3. Evita este error de la lección: Tratar embedding como extends. Escribe una predicción para una frontera antes de programar.
