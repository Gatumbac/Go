# Pistas · TCP y HTTP

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe un programa que parsee un request HTTP/1.1 fijo POST /notes con body abc y Content-Length 3.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
client → transport bytes → HTTP framing → handler
client ← status + headers + body ← application decision
```

3. Evita este error de la lección: Una Write=una Read. Escribe una predicción para una frontera antes de programar.
