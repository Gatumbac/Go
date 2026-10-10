# Pistas · Servidores y routing

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa handler GET /double/{n}, n entero -100..100.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
request → ServeMux method/path → Handler → ResponseWriter
                              validation → status/body
```

3. Evita este error de la lección: Header después de Write. Escribe una predicción para una frontera antes de programar.
