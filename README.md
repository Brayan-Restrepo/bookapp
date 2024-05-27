# CRUD Book -> Author

**Run the project**
```bash
go run server.go --env=dev
```
**Primeros pasos**
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

## CRUD - User

**Consultar Users**
```graphql
{
    users {
        id
        name
        books {
            id
            title
            author {
                name
            }
            createdAt
        }
    }
}
```
**Create User**
```graphql
mutation {
    createUser(name: "Juan") {
        id
        name
    }
}
```
**Update User**
```graphql
mutation {
    updateUser(id: 6, name: "Anis") {
        id
        name
    }
}
```
**Delete User**
```graphql
mutation {
    deleteUser(id: 5) {
        id
        name
    }
}
```

## CRUD - Book

**Get Book**
```graphql
{
    book(id:1) {
        id
        title
        author {
            name
        }
        createdAt
    }
}
```
**Get Books**
```graphql
{
  books {
    id
    title
    author {
      name
    }
    createdAt
  }
}
```
**Create Book**
```graphql
mutation {
  createBook(title: "The Go Programming Language", authorId: 2) {
    id
    title
    author {
      name
    }
  }
}
```
**Update Book**
```graphql
mutation {
  updateBook(
    id: 2
    title: "Concurrency in Go: Tools and Techniques for Developers"
    authorId: 1
  ) {
    id
    title
    author {
      name
    }
  }
}
```
**Delete Book**
```graphql
mutation {
  deleteBook(id: 4) {
    id
    title
  }
}
```

