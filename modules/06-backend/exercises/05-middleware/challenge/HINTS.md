# Pistas · Middleware y errores

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Crea dos middlewares que escriban marcas en un slice de traza durante una ejecución secuencial de test.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request → outer before → inner before → handler
response ← outer after ← inner after ← body/status
```

3. Evita este error de la lección: Orden no pensado. Escribe una predicción para una frontera antes de programar.
