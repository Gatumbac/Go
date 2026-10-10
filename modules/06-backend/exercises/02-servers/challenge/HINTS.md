# Pistas · Servidores y routing

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa GET /health con JSON {"status":"ok"}; otros métodos deben producir 405.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request → ServeMux method/path → Handler → ResponseWriter
                              validation → status/body
```

3. Evita este error de la lección: Header después de Write. Escribe una predicción para una frontera antes de programar.
