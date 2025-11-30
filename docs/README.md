# Gormigen documentation

## Project structure and generated code
- [migration package](./migration.md)

## CLI TOOL
```
gormigen [command] <options>
```
### Commands
- `init`  
        Adds an example structure, with one migration
        yml config file, and generates the code
- `generate`   
        generates the code based on config.yml file
- `add <name>`  
        adds an empty migration package - that you need to implement `Up` and `down` functions
