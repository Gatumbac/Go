# Pistas · Funciones y retornos

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa `DivideExact(a, b int) (int, bool)` para enteros no negativos: devuelve cociente y true solo si b>0 y división exacta; en otro caso 0,false.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
input → small function → value, status
caller → decide how to print / retry / reject
```

3. Evita este error de la lección: Ignorar retornos que señalan fallo. Escribe una predicción para una frontera antes de programar.
