package Queries

const InsertUserQuery = `insert into "Passenger" (firstname, lastname, email, phone, dateofbirth, passportserie, passportnumber, password) values ($1, $2, $3, $4, $5, $6, $7, $8);`

const CheckUserExistQuery = `select email, password from "Passenger" where email = $1`

const GetFlights = `select f.id, depairport.name as departure, desairport.name as arrival, f.departuedatetime, f.price from "Flight" f join "Airport" depairport on f.departueid = depairport.id join "Airport" desairport on f.destinationid = desairport.id join "City" depcitytable on depairport.cityid = depcitytable.id join "City" descitytable on desairport.cityid = descitytable.id where depcitytable.name = $1 and descitytable.name = $2 and f.departuedatetime::date = $3
`
