# Pistas · Servicios y testing HTTP

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Prueba cancelación y concurrencia del fake de tu servicio.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request → handler adapter → lookup capability
                              real implementation / fake
assert status + body + error translation
```

3. Evita este error de la lección: Fake como integración DB. Escribe una predicción para una frontera antes de programar.
