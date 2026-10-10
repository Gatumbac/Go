# Pistas · Clientes y recursos

[Enunciado](README.md)

Abre una por vez; no contienen implementación.

1. Relee el contrato: Implementa función `FetchText(ctx context.Context,client *http.Client,url string)(string,error)`, acepta solo status 200 y máximo 1024 bytes.

2. Usa este modelo como guía de razonamiento, no de implementación:

```text
Client.Do → Transport → response
     transport error OR status/body
     caller closes response body
```

3. Evita este error de la lección: No cerrar Body. Escribe una predicción para una frontera antes de programar.
