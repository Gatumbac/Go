# Laboratorio opcional · PostgreSQL y database/sql

[Lección SQL](../../../modules/06-backend/07-sql.md) · [Inicio](../../../README.md)

## Objetivo y requisitos

Conectar el pool de database/sql a una base PostgreSQL de laboratorio mediante pgx/stdlib. El módulo aislado declara pgx v5.7.6; el módulo principal conserva cero dependencias externas. No necesitas este laboratorio para leer el curso ni validar ejemplos principales.

Necesitas una instancia de prueba que tú administres. No uses producción; el programa no crea tablas ni modifica datos. Introduce la URL de conexión en la variable DATABASE_URL de tu sesión usando tus datos; no la escribas en Markdown, go files ni Git. Un ejemplo de formato es `postgres://USER:PASSWORD@HOST:5432/DB?sslmode=require`; usa la política TLS adecuada a tu laboratorio, no desactives validación en sistemas reales.

## Ejecución manual

Desde la raíz:

```sh
cd docs/labs/postgres
go mod tidy
go build -o /tmp/go-postgres-lab .
go run .
```

`go mod tidy` descarga dependencias y genera go.sum dentro de este módulo; no se ejecuta como parte del recorrido principal. La salida esperada es `9`, con parámetros 4 y 5. El timeout cubre Ping y consulta. No se imprime la URL ni el detalle de conexión que podría revelar credenciales.

## Qué observar

Open crea el handle; Ping comprueba disponibilidad. SetMaxOpenConns limita a dos conexiones: no significa dos jobs del negocio. QueryRowContext usa placeholders de PostgreSQL y Scan lee un valor. Close libera el pool al terminar este programa finito.

## Tu práctica posterior

Diseña tú tabla y consulta para ciudades en una base de laboratorio. Prueba filas presentes/ausentes y un contexto cancelado. Si creas una transacción, usa operaciones de tx, rollback en fallos y commit en éxito. No se incluye solución de ese ejercicio.

## Evidencia y límites

Este laboratorio no fue conectado a una DB durante la generación. Compilación y conexión son evidencias distintas; registra ambas al ejecutarlo. El driver puede devolver errores específicos que el ejemplo deliberadamente traduce a mensajes seguros.

[pgx database/sql](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx-through-database-sql) · [Release utilizada](https://github.com/jackc/pgx/releases/tag/v5.7.6) · [Gestión de conexiones](https://go.dev/doc/database/manage-connections)
