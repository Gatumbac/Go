# Pistas · Escape analysis

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea dos funciones con el mismo resultado lógico, una devuelve struct pequeño y otra pointer.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
reference flow → compiler proof?
 yes: local storage possible
 no: longer-lived storage may be needed
measure → determine practical cost
```

3. Evita este error de la lección: Escapes como fallo de calidad. Escribe una predicción para una frontera antes de programar.
