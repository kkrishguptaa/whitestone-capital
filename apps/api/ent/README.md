# ent

`ent` is a ORM for Go. Most files in this directory are auto-generated by the `ent` tool. You can find the schema definitions in the `schema` subdirectory.

## To generate the ent code

Run the following command in the root of the project:

```bash
go generate ./ent
```

## To create a new schema

Run the following command in the root of the project:

```bash
go run -mod=mod entgo.io/ent/cmd/ent new <SchemaName>
```

## To update the ent code

Edit the schema files in the `ent/schema` directory and then run:

```bash
go generate ./ent
```
