# Pistas · Organización e idiomas

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Dibuja en architecture.md un grafo para una API de biblioteca con catálogo y almacenamiento.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
HTTP adapter → application → domain
                    ↓
               storage capability
(no reverse import from domain to HTTP)
```

3. Evita este error de la lección: Packages utils omnipresentes. Escribe una predicción para una frontera antes de programar.
