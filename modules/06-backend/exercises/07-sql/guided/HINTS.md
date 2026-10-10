# Pistas · Database/sql

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Escribe query-review.md: entrada, consulta de productos por precio mínimo y sort elegido por usuario. Salida:SQL con placeholder para precio y allowlist de sort. Restricción: no conectar DB. Aceptación: distinguir valor de identificador.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
sql.DB pool → borrowed connection → query → rows close → reusable
transaction: begin → operations on tx → commit OR rollback
```

3. Evita este error de la lección: Open=conectividad. Escribe una predicción para una frontera antes de programar.
