# babelfish
Babelfish samples applications in Golang

This repo provides you with examples of applications speaking in MsSQL dialect to the Babelfish server in front of PostgreSQL:

`App -> Babelfish -> PostgrSQL`

## Quick Start
0. Make sure you have PostgreSQL and Babelfish installed and running.
1. Clone the repository 
```sh
$ gh repo clone cybertec-postgresql/babelfish
$ cd babelfish
```
2. Check credentials used by default in `main.go` or change connection string in the code of samples

3. Run the sample: 
```sh
$  go run main.go
2021/11/03 17:26:18 Connected!

MSSQL version:
Babelfish for PostgreSQL with SQL Server Compatibility - 12.0.2000.8
Nov  2 2021 10:59:02
Copyright (c) Amazon Web Services
PostgreSQL 13.4 Babelfish for PostgreSQL on x86_64-pc-linux-gnu

MSSQL dbname:
master

Employee #3:
Ants from Estonia
```
