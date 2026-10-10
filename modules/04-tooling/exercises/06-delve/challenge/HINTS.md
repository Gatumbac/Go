# Pistas · Debugging con Delve

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Depura una función tuya que modifica un slice compartido.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
reproduce → hypothesis → breakpoint before mutation
                         ↓ inspect before/after
                    explain cause → verify repair
```

3. Evita este error de la lección: Cambiar código sin hipótesis. Escribe una predicción para una frontera antes de programar.
