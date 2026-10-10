# Pistas · Errores y wrapping

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Define un sentinel `ErrEmpty` y función `First(values []int) (int,error)`.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
failure → context wrapping → caller classification
               preserves cause        ↓
                                  retry / reject
```

3. Evita este error de la lección: Comparar mensajes. Escribe una predicción para una frontera antes de programar.
