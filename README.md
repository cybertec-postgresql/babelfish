# babelfish
Babelfish samples applications in Golang

This repo provides you with examples of applications speaking in MsSQL dialect to the Babelfish server in front of PostgreSQL:

`App -> Babelfish -> PostgrSQL`

## Folder Structure:
- `babelfishdb`: helper package, provides `Open` and `Get` methods
- `CreateTable`: shows how to use MsSQL spefific DDL construction through Babelfish against underlying PostgreSQL
- `SelectVersion`: shows how to get server version and current database name


## Quick Start
0. Make sure you have PostgreSQL and Babelfish installed and running.
1. Clone the repository 
```sh
$ gh repo clone cybertec-postgresql/babelfish
$ cd babelfish
```
2. Check credentials used by default in `babelfishdb/babelfishdb.go` or change connection string in the code of samples

3. Choose the sample and change directory
```sh
$ cd CreateTable
```

4. Run the sample: 
```sh
$ go run main.go
2021/11/02 16:07:39 Connected!

Employee #3:
Tom from Germany
```
