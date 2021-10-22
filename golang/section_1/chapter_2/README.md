# Arrays
1. Arrays have fixed size.
2. To add more elements in an array we need to create a bigger array, copy the old elements into it and then add the new array elements.
3. If the array is big then the above operation might give performance issues.  

Creating arrays 
```
var arr = [5]int{1, 2, 3, 4, 5}
```
Checking type of `arr`
```
[5]int
```
# Slices
1. An array has a fix size and so elements can't be added to it. To solve this problem we create a slice which has a capacity of double the length of elemnts with which it is initialized.
2. The slice is just a reference to the pointer of array that is created in the background when slice is initialized.
3. Observe the program `slice-pointer.go`. In this when we change the element in slice `b` then the element in slice `a` also changes. This is because they both point to the same array.
4. Similar to this, the program `slice.go` states that the slice `slice` has a capacity similar to the underlying array `arr` starting from its lower bound which is pointed to the index 1 of array `arr`.

# 2-D slices or 2-D array
The two dimensional array and slices can be created the same way. The examples can be found in the 2-d-array-slices directory.

# Maps
Maps are data structure that helps in storing key-value pairs.


# Database
Here we are trying to perform CRUD operations on mysql database using `database/sql` library with `mysql` driver. To run the application we need to use docker container for mysql application.

1. Run the following command to start the docker container.
```
docker-compose up -d
```
2. Go inside the docker container by running the command 
```
docker container exec -it database_db_1 bash
```
3. Enter the root password and create the table named `customer`.
```
root@723fe8986bc8:/# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.26 MySQL Community Server - GPL

Copyright (c) 2000, 2021, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> use crm;
Database changed
mysql> CREATE TABLE customer (id int NOT NULL AUTO_INCREMENT, name varchar(255), ssn varchar(255), PRIMARY KEY (id));
Query OK, 0 rows affected (0.03 sec)

mysql> show tables;
+---------------+
| Tables_in_crm |
+---------------+
| customer      |
+---------------+
1 row in set (0.00 sec)

# This is to insert the values with the ID
mysql> INSERT INTO customer values(1,"Vedant Pareek", 39048732);
Query OK, 1 row affected (0.01 sec)

# As ID is auto-incremented, so we can use the below syntax enter details without passing the ID
mysql> INSERT INTO customer (name, ssn) values("Harry Potter", "294587230");
Query OK, 1 row affected (0.11 sec)

mysql> INSERT INTO customer values(3,"Hermionie Granger", "294387230");
Query OK, 1 row affected (0.52 sec)

mysql> select * from customer
    -> ;
+------+-------------------+-----------+
| id   | name              | ssn       |
+------+-------------------+-----------+
|    1 | Vedant Pareek     | 39048732  |
|    2 | Harry Potter      | 294587230 |
|    3 | Hermionie Granger | 294387230 |
+------+-------------------+-----------+
3 rows in set (0.00 sec)
```
4. To retrieve the contents of the database
```
go run get_customer_list.go
```
**NOTE: `defer` statements are called after `panic` statements are executed.**

# Webforms
In this we will learn about webforms using `net/http` package from golang. In this we will perform the CRUD operations which are CREATE, READ, UPDATE, DELETE.
1. To run the basic http server, this will display the `main.html` over port 8000
```
go run webforms.go
```
2. To perform the operation through UI over `crm` database, we will go inside the `webforms` directory and then execute `go run .`
3. Post this hit the browser with the URL `http://127.0.0.1:8000` to perform the required operations.