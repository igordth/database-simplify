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

### Config

* `Name` - name of database
* `User` - user for connection
* `Password` - password of user
* `Host` - host of database
* `Schema` - scheme of database, if not set default postgres schema usage - public
* `MaxOpenConn` - sets the maximum number of open connections to the database
* `MaxIdleConn` - maximum number of connections in the idle pool
* `Options` - other options for [connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING)

### Creating connection

```go
cnn, df, err := pggorm.NewConnection(cfg, log)
if err != {
	// todo
	return
}
defer df()
```

### Base connection methods

 * `DB()`
 * `Gorm(ctx context.Context)`

### Connection methods for transaction

* `TrxBegin(ctx context.Context, opts ...*sql.TxOptions)` - creates a context through which the transaction will proceed
* `TrxRollback(ctx *context.Context)` - roll back a transaction by context
* `TrxCommit(ctx *context.Context)` - commit transaction by context

