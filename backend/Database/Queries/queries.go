package Queries

const InsertUserQuery = `insert into "Passenger" 
    (firstname, lastname, email, phone, dateofbirth, passportserie, passportnumber, password) 
	values ($1, $2, $3, $4, $5, $6, $7, $8);`

const CheckUserExistQuery = `select email, password from "Passenger" where email = $1`

const GetFlights = `select
			f.id,
			a1.name AS departure_airport,
			c1.name AS departure_city,
			a2.name AS arrival_airport,
			c2.name AS arrival_city,
			f.departuedatetime AS departure_date,
			f.arrivaldatetime AS arrival_date,
			f.price,
			(ac.rowsamount * ac.seatsinrowamount) - coalesce(sum(case when fb.bookingstatus in ('booked', 'paid') then 1 else 0 end), 0) as available_seats
		from
			"Flight" f
		join "Airport" a1 on f.departueid = a1.id
		join "City" c1 on a1.cityid = c1.id
		join "Airport" a2 on f.destinationid = a2.id
		join "City" c2 on a2.cityid = c2.id
		join "Aircraft" ac on f.aircraftid = ac.id
		left join "FlightBooking" fb ON f.id = fb.flightid
		where
			c1.name = $1
			and c2.name = $2
			and f.departuedatetime::date = $3
			and f.departuedatetime > NOW()
		group by
			f.id, a1.name, c1.name, a2.name, c2.name, f.departuedatetime, f.arrivaldatetime, f.price, ac.rowsamount, ac.seatsinrowamount;`

const GetUserLongByEmailQuery = `select id, firstname, lastname, email, phone, dateofbirth, 
	passportserie, passportnumber, password from "Passenger" where email = $1`

const GetSeatsForFlightQuery = `SELECT
    chr(64 + row_series.row) AS row,
    seat_series.seat_in_row AS seat,
    CASE
        WHEN fb.id IS NOT NULL AND fb.bookingStatus != 'canceled' THEN 'unavailable'
        ELSE 'available'
    END AS status
FROM "Flight" f
JOIN "Aircraft" ac
    ON ac.id = f.aircraftId
CROSS JOIN generate_series(1, ac.rowsamount) AS row_series(row)
CROSS JOIN generate_series(1, ac.seatsinrowamount) AS seat_series(seat_in_row)
LEFT JOIN "FlightBooking" fb
    ON fb.flightId = f.id
    AND fb.row::text = chr(64 + row_series.row)
    AND fb.seatInRow = seat_series.seat_in_row 
WHERE f.id = $1
ORDER BY row_series.row, seat_series.seat_in_row;`

const InsertBooking = `insert into "FlightBooking" (flightId, passengerId, bookingStatus, bookingDateTime, row, seatInRow) values ($1, $2, $3, NOW(), $4, $5);`

const GetUserIdByEmail = `select id from "Passenger" where email = $1`

const UpdateUserById = `UPDATE "Passenger"
	SET 
		firstName = $1,
		lastName = $2,
		email = $3,
		phone = $4,
		dateOfBirth = $5,
		passportSerie = $6,
		passportNumber = $7
	WHERE 
		id = $8;`

const GetHistory = `SELECT
	fb.id, 
    da.name AS departure_airport,
    dc.name AS departure_city,
    aa.name AS arrival_airport,
    ac.name AS arrival_city,
    f.departueDateTime AS departure_date,
    f.arrivalDateTime AS arrival_date,
    f.price::TEXT AS price,
    fb.bookingstatus AS status,
    fb.row AS row,
    fb.seatInRow AS seat
FROM
    "FlightBooking" fb
JOIN
    "Flight" f ON fb.flightId = f.id
JOIN
    "Airport" da ON f.departueId = da.id
JOIN
    "Airport" aa ON f.destinationId = aa.id
JOIN "City" dc on da.cityid = dc.id
JOIN "City" ac on aa.cityid = ac.id
JOIN "Passenger" p on fb.passengerid = p.id
WHERE
    p.email = $1;`
