# Pistas · Compilador, binario y runtime

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe `build-notes.md`: entrada, un binario construido; salida, qué pasaría al borrar el source sin borrar el binario y qué pasaría en otra arquitectura. No necesitas borrar nada. Aceptación: separar hipótesis de observación y mencionar enlace dinámico.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
source .go → type checking → object code → linker → executable
                                                       |
                                                       v
                                             main + Go runtime
```

3. Evita este error de la lección: Intentar ejecutar un binario de otro OS. Escribe una predicción para una frontera antes de programar.
