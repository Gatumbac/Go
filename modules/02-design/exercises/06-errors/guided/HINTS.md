# Pistas · Errores y wrapping

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `ParseAge(text string) (int, error)` usando strconv: dominio 0..130.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
failure → context wrapping → caller classification
               preserves cause        ↓
                                  retry / reject
```

3. Evita este error de la lección: Comparar mensajes. Escribe una predicción para una frontera antes de programar.
