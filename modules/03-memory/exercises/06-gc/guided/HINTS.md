# Pistas · Garbage collection

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe retention.md: entrada, un cache map sin eliminación que conserva respuestas grandes. Salida: grafo de alcanzabilidad y una política acotada. Restricción: no afirmar que runtime.GC arregla retención. Aceptación: distinguir objeto vivo de basura.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
roots → reachable graph → keep
         no path to object → eligible for reclamation
resource Close ≠ immediate heap reclamation
```

3. Evita este error de la lección: GC=Close. Escribe una predicción para una frontera antes de programar.
