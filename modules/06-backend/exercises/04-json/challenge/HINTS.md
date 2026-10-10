# Pistas · JSON

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea handler POST /settings con body limitado a 1024 bytes y campo enabled requerido usando *bool.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
limited bytes → decode DTO → validate domain → use
                unknown/trailing input → reject by contract
```

3. Evita este error de la lección: Campos minúsculos. Escribe una predicción para una frontera antes de programar.
