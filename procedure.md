


# Using Gopherjs 

https://github.com/tdewolff/parse/issues/43
To make this compatible:

minify -
```sh
 git checkout tags/v2.3.6
cd ../parse
git checkout tags/v2.3.4
```

# Certs
1. Generate certs
   a. go run $GOROOT/src/crypto/tls/generate_cert.go -host localhost (run in the application root)
   b. Add http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})

* to clean up old images

```sh 
docker system prune --all --force --volumes
```
 ##  mySQL
1. set up docker mysql	
```docker pull mysql```

*   Start server 	
```docker run -p3306:3306 --name gopherfacemysql -e MYSQL_ROOT_PASSWORD=gopherface -d mysql``

## Start client	
```docker run -it --link gopherfacemysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"gopherface"'```
	
	
DDL statements 	Database Definition Languge
```sql
# Create user	

CREATE USER 'gopherface'@'%' IDENTIFIED BY 'gopherface';
# Create DB (backtics, not single quotes)	
CREATE DATABASE IF NOT  EXISTS `gopherfacedb` DEFAULT CHARACTER SET 'utf8' COLLATE `utf8_unicode_ci`;

# GRANT PRIVILEGES	

GRANT All PRIVILEGES ON gopherfacedb.* TO 'gopherface'@'%';

# FLUSH PRIVILEBES SO CHANGES TAKE AFFECT	
FLUSH PRIVILEGES;
```	
	
## DML 	Database Manipulation Language
	
	
docker run -p3306:3306 --name songsmysql -e MYSQL_ROOT_PASSWORD=songslist  -d mysql	
docker exec -it 3a33c2649b6f /bin/bash	
* In db config directory

docker cp gopherfacedb.sql heuristic_jackson:/tmp/

* In the mysql client  exec session
source /tmp/gopherfacedb.sql

```mysql 
use gopherfacedb;

SELECT USER();

describe views

select distinct userid from cells

describe cells;
desc user;
+---------------+---------------------+------+-----+-------------------+-----------------------------------------------+
| Field         | Type                | Null | Key | Default           | Extra                                         |
+---------------+---------------------+------+-----+-------------------+-----------------------------------------------+
| id            | tinyint(1) unsigned | NO   | PRI | NULL              | auto_increment                                |
| username      | varchar(18)         | NO   | UNI | NULL              |                                               |
| uuid          | varchar(64)         | NO   |     | NULL              |                                               |
| first_name    | varchar(64)         | NO   |     | NULL              |                                               |
| last_name     | varchar(64)         | NO   |     | NULL              |                                               |
| password_hash | char(64)            | NO   |     | NULL              |                                               |
| email         | varchar(255)        | NO   |     | NULL              |                                               |
| created_ts    | timestamp           | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED                             |
| updated_ts    | timestamp           | NO   |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED on update CURRENT_TIMESTAMP |
+---------------+---------------------+------+-----+-------------------+-----------------------------------------------+
9 rows in set (0.01 sec) 

select * from user where username = "JayneJacobs";
```

Go Docs
```sh
go doc http.Request.FormValue
godoc --http :6060
```
