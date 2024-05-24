# CRUD Book -> Author

```bash
mkdir bookapp
cd gqlgen-todos
go mod init github.com/Brayan-Restrepo/bookapp
```

Se crea el archivo tools.go

```bash
//go:build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
)
```
```bash
go mod tidy
go get -d github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init

```


El archivo graph/schema.graphqls desempeña un papel central en el correcto funcionamiento de nuestro servicio. Es aquí donde definimos la estructura de datos y las operaciones que nuestro servidor GraphQL admitirá. Es crucial asegurarse de que este archivo esté configurado correctamente para garantizar que el proceso de generate resolvers funcione sin problemas.

Después de realizar ajustes en el archivo graph/schema.graphqls, es crucial asegurarse de que el archivo schemas.resolvers.go sea eliminado. Este archivo es generado automáticamente por el proceso de generación de resolutores (generate resolvers), y contiene implementaciones de resolutores que pueden quedar obsoletas si se realizan cambios en el esquema GraphQL.

Para eso debe de ejecutar cualquiera de estos 2 comandos:

```bash
go mod tidy
go generate 
go run github.com/99designs/gqlgen generate
```

