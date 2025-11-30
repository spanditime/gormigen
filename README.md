# gormigen
Gormigen is a tool for managing migrations, and generating boilerplate code for [Gorm](https://gorm.io/) projects using [gormigrate](https://github.com/go-gormigrate/gormigrate).  
It provides easy tool to create new migrations, make sure that the migrations are run in order of creation(similar to other migration tools) and generates boilerplate for running database creation and migration with gorm via gormigrate.
## Usage
```
gormigrate [command] <options>
```
__Availabe commands are:__
- `init` - creates an example of code generation usage in a project [*WIP*]
- `add <name>` - creates a new migration, versioned by todays date and number
- `generate` - regenerates boilerplate based on currently existing migrations and config file

## Getting started
Add a tool to your project
```
    go get -tool github.com/spanditime/gormigen
```
### Generate initial project structure example 
> Note: this is currently *WIP* - refer to next topic or examples
```
gormigen init
```
#### adding to existing project
Create `gormigen.yml` in root of your repo, where `go.mod` located
```yml
migrations:
  path: path/to/db/migrations

# there is a way to enable initial database schema
# init_migration:
#   package: github.com/spanditime/gormigen_example/path/to/db/init

output:
  filename: internal/db/migrations.generated.go
  package_name: db_migrations
```
In your package where migration code will be generated - add a file with config
```go
package db_migrations // same as output.package_name in gormigen.yml in same directory

import "github.com/go-gormigrate/gormigrate/v2"

//go:generate gormigen generate

var MigrationConfig = &gormigrate.Options{
	TableName: "migrations",
}
```
## Feature plans
- [ ] `init` command
- [ ] usage example 
- [x] documentation
- [ ] add a way to not apply migration and roll it back if it was applied
- [ ] add fix-index command - to fix collisions in migration versions automatically

