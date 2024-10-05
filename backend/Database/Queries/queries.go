package Queries

const InsertUserQuery = `insert into "Passenger" (firstname, lastname, email, phone, dateofbirth, passportserie, passportnumber, password) values ($1, $2, $3, $4, $5, $6, $7, $8);`

const CheckUserExistQuery = `select email, password from "Passenger" where email = $1`
