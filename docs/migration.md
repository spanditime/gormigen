# Migration package
Migration packages are created with 
```
gormigen add <name>
```
That creates directory with name `v{date}{number}-{snake_case_name}`  
with single file - `{snake_case_name}.go` - with this structure
```go
package v{date}{number}

import (
	"errors"
	"gorm.io/gorm"
)

//-------------------------------------------------------------
// {name}
//-------------------------------------------------------------

func Up(db *gorm.DB) error {
	// write your up migration code here
	return errors.New("not implemented")
}

func Down(db *gorm.DB) error {
	// write your down migration code here
	return errors.New("not implemented")
}
```
If the date of last existing migration is in the future from today - `gormigen add` will produce an error - to try and ensure that migrations are executed in the order of creation

## Migration order
Migration order is decided by its version - sorted by date and number.

If you want to add a migration in between some other(already existing) transactions for some reason - you can add new transaction and just rename it
```
v202511280001-user_email_index
v202511280002-change_user_fk_from_username_to_id
v202511280003-remove_username -> v202511280004-remove_username 
v202511290001-my_new_migration -> v202511290003-my_new_migration 
```
> Altho its not desired, you should rarely encounter such problems.
If the need to do so comes from merging code and the two features have collision(my_new_migration added a table which referes to username instead of id which was removed by other developer) - then it probably is a problem with your migrations and that they arent incremental enough.  

