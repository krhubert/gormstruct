# gormstruct

This tool creates models and CRUD methods for PostgreSQL database.

# usage

```
gormstruct -schema public -conn "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable" -out=./model -pkg-name=model -skip-tables=db_migration,db_stats

```

print help:
```
gormstruct -h
```

# limitations

- does not support some arrays and user defined types
- does not support domains
