[!Important]
create the test DB before running tests:
```shell
sqlite3 test.db "create table if not exists todos (id integer not null primary key, description text, completed boolean, created_at int);"
```


[!Note] 
A more idiomatic approach would be to keep test files in the same package as 
the functions they are testing, but this is more similar to an integration
test than a unit test so it should be better to have it separate

to run the test run this command from the project root (where go.mod resides)
```shell
go test ./test
```
