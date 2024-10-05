create schema courseWork;

create table if not exists "City"
(
    id serial  primary key,
    name    varchar(50) not null,
    country varchar(50)     not null
);

create table if not exists "Airline"
(
    id serial  primary key,
    name    varchar(50) not null,
    IATA varchar(3)    unique not null,
    phone varchar(20) not null,
    website varchar(50) not null
);

create table if not exists "Airport"
(
    id serial  primary key,
    name    varchar(50) not null,
    cityId integer constraint Airport_city_id_fk
        references "City"
        on update restrict on delete restrict,
    IATA varchar(3)    unique not null
);

create table if not exists "Aircraft"
(
    id serial  primary key,
    model    varchar(20) not null,
    airlineId integer constraint Aircraft_city_id_fk
        references "Airline"
        on update restrict on delete restrict,
    regNumber    varchar(20) not null,
    rowsAmount smallint,
    seatsInRowAmount smallint
);

create table if not exists "Flight"
(
    id serial  primary key,
    aircraftId integer constraint Flight_aircraft_id_fk
        references "Aircraft"
        on update restrict on delete restrict,
    departueId integer constraint Flight_departue_id_fk
        references "Airport"
        on update restrict on delete restrict,
    destinationId integer constraint Flight_destination_id_fk
        references "Airport"
        on update restrict on delete restrict,
    departueDateTime timestamp not null,
    arrivalDateTime timestamp not null,
    price money
);

create table if not exists "Passenger"
(
    id serial  primary key,
    firstName    varchar(50) not null,
    lastName    varchar(50) not null,
    email    varchar(50) not null,
    phone    varchar(20) not null,
    dateOfBirth date not null,
    passportSerie varchar(10) not null,
    passportNumber varchar(10) not null,
    password varchar(50) not null
);

create type status as enum ('booked', 'paid', 'canceled', 'done');
create type "row" as enum ('A', 'B', 'C', 'D', 'E', 'F');

create table if not exists "FlightBooking"
(
    id serial  primary key,
    flightId integer constraint FlightBooking_flight_id_fk
        references "Flight"
        on update restrict on delete restrict,
    passengerId integer constraint FlightBooking_passenger_id_fk
        references "Passenger"
        on update restrict on delete restrict,
    bookingStatus   status,
    bookingDateTime timestamp,
    row "row",
    seatInRow smallint
);