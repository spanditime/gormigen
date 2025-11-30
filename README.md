# gormigen
Gormigen is a tool for managing migrations, and generating boilerplate code for [gorm]() projects using [gormigrate](https://github.com/go-gormigrate/gormigrate).
It provides easy tool to create new migrations, make sure that the migrations are run in order of creation(similar to other migration tools) and generates boilerplate for funning database creation and migration with gorm.
## Usage
```
gormigrate [command] <options>
```
__Availabe commands are:__
- `init` - creates an example of code generation usage in a project
- `add <name>` - creates a new migration, versioned by todays date and number
- `generate` - regenerates boilerplate based on currently existing migrations and config file

## Getting started
Install or add it as a tool to your project
### Generate initial project structure example [WIP]
```
gormigen init
```
#### adding to existing project
Create `gormigen.yml` 
```yml
TODO: make an example
```
In your package where migration code will be generated - add a file with config
```go
package db_migration

//go:generate gormigen generate

TODO: make an example
```

