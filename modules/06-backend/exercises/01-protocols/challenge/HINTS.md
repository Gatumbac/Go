# Pistas · TCP y HTTP

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe retry-notes.md para caída de conexión después de enviar una compra.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
client → transport bytes → HTTP framing → handler
client ← status + headers + body ← application decision
```

3. Evita este error de la lección: Una Write=una Read. Escribe una predicción para una frontera antes de programar.
