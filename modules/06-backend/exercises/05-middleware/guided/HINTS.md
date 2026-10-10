# Pistas · Middleware y errores

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa middleware que requiera header X-Client no vacío y responda 400 sin llamar next si falta.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request → outer before → inner before → handler
response ← outer after ← inner after ← body/status
```

3. Evita este error de la lección: Orden no pensado. Escribe una predicción para una frontera antes de programar.
