database-simplify
============

Package for simplify working with databases.

# Install

Install with

```shell
go get github.com/igordth/database-simplify
```

Or import

```go
import "github.com/igordth/database-simplify"
```

# Database list

* [PostgreSQL](https://www.postgresql.org/)

# Basic ORM package

* [gorm](https://gorm.io/)

## pggorm

Package for work with PostgreSQL using gorm.

### Creating logger for connection with `zap.Logger`

```go
log := NewLog(log, cfg)
```

where 
 * log is `*zap.Logger` [hear](https://pkg.go.dev/go.uber.org/zap)
 * cfg is config from `gorm.io/gorm/logger` [hear](https://gorm.io/docs/logger.html)

If log is nil returned `logger.Discard` - print any log to io.Discard

_PS_: If you don\`t want use logger put `logger.Discard` to `NewConnection` or nil to `NewLog` - don\`t use `zap.NewNope` because it is more resource-intensive

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/connection

### Config

* `Name` - name of database
* `User` - user for connection
* `Password` - password of user
* `Host` - host of database
* `Schema` - scheme of database, if not set default postgres schema usage - public
* `MaxOpenConn` - sets the maximum number of open connections to the database
* `MaxIdleConn` - maximum number of connections in the idle pool
* `Options` - other options for [connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/connection

### Creating connection

```go
cnn, df, err := pggorm.NewConnection(cfg, log)
if err != {
	// todo
	return
}
defer df()
```

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/connection

### Base connection methods

 * `DB()`
 * `Gorm(ctx context.Context)`

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/connection

### Connection methods for transaction

* `TrxBegin(ctx context.Context, opts ...*sql.TxOptions)` - creates a context through which the transaction will proceed
* `TrxRollback(ctx *context.Context)` - roll back a transaction by context
* `TrxCommit(ctx *context.Context)` - commit transaction by context

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/transaction

### Usage

Append default methods to gateways [db->schema->table-gateway].

* `Count` - get count of records. [gorm doc](https://gorm.io/docs/advanced_query.html#Count)
* `CreateModel` - create record by model. [gorm doc](https://gorm.io/docs/create.html#Create-Record)
* `CreateMap` - create record by map. [gorm doc](https://gorm.io/docs/create.html#Create-Record)
* `Delete` - deletes value matching given conditions. [gorm doc](https://gorm.io/docs/delete.html)
* `Find` - retrieving object(s) with conditions. [gorm doc](https://gorm.io/docs/query.html)
* `Save` - updates value in database. If value doesn't contain a matching primary key, value is inserted. [gorm doc](https://gorm.io/docs/update.html#Save-All-Fields)
* `Update` - updates column with value using callbacks. [gorm doc](https://gorm.io/docs/update.html#Update-single-column)
* `Updates` - updates attributes using callbacks. values must be a struct or map. [gorm doc](https://gorm.io/docs/update.html#Updates-multiple-columns)

Add needed methods to your gateway like:

```go
type Model struct {
    ID     uint   `gorm:"primary_key;column:id"`
    Name   string `gorm:"column:name"`
}

func (Model) TableName() string {
    return "some_table_name"
}

type MySomeTableName struct {
	pggorm.Connect
	usage.FindCompare[Model]
}

func New(cnn pggorm.Connect) *MySomeTableName {
	return &MySomeTableName{
		cnn,
		usage.NewFindCompare[Model](cnn),
	}
}

// todo my extra methods of MySomeTableName
```

Link:
 * [Grom Model](https://gorm.io/docs/models.html)

And use like:

```go
gate := New(cnn)
result, err := gate.Find.One.Execute(ctx) // result is *Model of firs row from table some_table_name
if err != nil {
    return panic(err)
}
fmt.Printf("%+v", result)
```

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/usage

#### With

Conditions for `usage`. To create more complex queries.

* `Preload(query string, args ...any)` - preload associations with given conditions. [gorm doc](https://gorm.io/docs/preload.html#Preload)
* `Order(values ...string)` - specify order when retrieving records from database. [gorm doc](https://gorm.io/docs/query.html#Order)
* `Where(query any, args ...any)` - add where conditions. [gorm doc](https://gorm.io/docs/query.html#Conditions)
* `Limit(limit int, offset int)` - specify the [limit] number of records to be retrieved and skip [offset] before starting. [gorm doc](https://gorm.io/docs/query.html#Limit-amp-Offset)
* `Joins(query string, args ...any)` - specify Joins conditions. [gorm doc1](https://gorm.io/docs/query.html#Joins) [gorm doc2](https://gorm.io/docs/query.html#Joins-Preloading)
* `GroupBy(query string)` - specify the group method on the find. [gorm doc](https://gorm.io/docs/query.html#Group-By-amp-Having)
* `Having(query string, args ...any)` - specify HAVING conditions for GROUP BY. [gorm doc](https://gorm.io/docs/query.html#Group-By-amp-Having)
* `Distinct(values ...string)` - specify distinct fields that you want querying. [gorm doc](https://gorm.io/docs/query.html#Distinct)
* `Select(query any, args ...any)` - specify fields that you want when querying, creating, updating. [gorm doc](https://gorm.io/docs/query.html#Selecting-Specific-Fields)
* `Omit(columns ...string)` - specify fields that you want to ignore when creating, updating and querying. [gorm doc](https://gorm.io/docs/associations.html#Select-x2F-Omit-Association-fields)

Use like this:

```go
result, err := gate.Find.Many.
    With(
        with.Where("name like ?", "Bi%"),
        with.Order("name desc"),
    ).
    With(with.Limit(5, 0)).
    Execute(ctx)
if err != nil {
    return panic(err)
}
fmt.Printf("%+v", result)
```

_Example in:_ https://github.com/igordth/database-simplify/tree/master/examples/pggorm/usage