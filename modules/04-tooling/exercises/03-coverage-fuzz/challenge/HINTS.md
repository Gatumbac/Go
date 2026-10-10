# Pistas · Coverage y fuzzing

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea un fuzz test de una transformación reversible escrita por ti, con strings acotados a 1024 bytes.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
seeds → mutated inputs → property check
                          ↓ violation
                    minimal regression input
```

3. Evita este error de la lección: Coverage como garantía. Escribe una predicción para una frontera antes de programar.
