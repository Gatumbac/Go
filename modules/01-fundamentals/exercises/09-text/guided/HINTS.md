# Pistas · Strings, bytes y runes

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `CountRunes(text string) int` para UTF-8 válido.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
"ñA" → bytes: C3 B1 41 (len=3)
       runes: U+00F1 U+0041 (count=2)
       offsets from range: 0, 2
```

3. Evita este error de la lección: Llamar caracteres a len. Escribe una predicción para una frontera antes de programar.
