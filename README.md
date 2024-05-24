# CRUD Book -> Author

El archivo graph/schema.graphqls desempeña un papel central en el correcto funcionamiento de nuestro servicio. Es aquí donde definimos la estructura de datos y las operaciones que nuestro servidor GraphQL admitirá. Es crucial asegurarse de que este archivo esté configurado correctamente para garantizar que el proceso de generate resolvers funcione sin problemas.

Después de realizar ajustes en el archivo graph/schema.graphqls, es crucial asegurarse de que el archivo schemas.resolvers.go sea eliminado. Este archivo es generado automáticamente por el proceso de generación de resolutores (generate resolvers), y contiene implementaciones de resolutores que pueden quedar obsoletas si se realizan cambios en el esquema GraphQL.

Para eso debe de ejecutar cualquiera de estos 2 comandos:

```bash
go mod tidy
go generate 
go run github.com/99designs/gqlgen generate
```

