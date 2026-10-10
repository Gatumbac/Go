# Pistas · Escape analysis

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Construye un benchmark de esas funciones después del módulo de herramientas; por ahora escribe el protocolo de medición.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
reference flow → compiler proof?
 yes: local storage possible
 no: longer-lived storage may be needed
measure → determine practical cost
```

3. Evita este error de la lección: Escapes como fallo de calidad. Escribe una predicción para una frontera antes de programar.
