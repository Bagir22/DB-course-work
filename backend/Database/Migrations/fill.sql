

INSERT INTO "City" (name, country) VALUES
                                       ('New York', 'USA'),
                                       ('Los Angeles', 'USA'),
                                       ('London', 'UK'),
                                       ('Berlin', 'Germany'),
                                       ('Tokyo', 'Japan');

INSERT INTO "Airline" (name, IATA, phone, website) VALUES
                                                       ('American Airlines', 'AA', '800-433-7300', 'www.aa.com'),
                                                       ('Delta Airlines', 'DL', '800-221-1212', 'www.delta.com'),
                                                       ('British Airways', 'BA', '800-247-9297', 'www.britishairways.com'),
                                                       ('Lufthansa', 'LH', '800-645-3880', 'www.lufthansa.com'),
                                                       ('Japan Airlines', 'JL', '800-525-3663', 'www.jal.co.jp');

INSERT INTO "Airport" (name, cityId, IATA) VALUES
                                               ('John F. Kennedy International Airport', 1, 'JFK'),
                                               ('Los Angeles International Airport', 2, 'LAX'),
                                               ('Heathrow Airport', 3, 'LHR'),
                                               ('Berlin Tegel Airport', 4, 'TXL'),
                                               ('Tokyo Haneda Airport', 5, 'HND');

INSERT INTO "Aircraft" (model, airlineId, regNumber, rowsAmount, seatsInRowAmount) VALUES
                                                                                       ('Boeing 737', 1, 'N123AA', 30, 6),
                                                                                       ('Airbus A320', 2, 'N234DL', 32, 6),
                                                                                       ('Boeing 777', 3, 'G-BNLH', 40, 8),
                                                                                       ('Airbus A350', 4, 'D-ABCD', 40, 9),
                                                                                       ('Boeing 787', 5, 'JA123A', 35, 8);

INSERT INTO "Flight" (aircraftId, departueId, destinationId, departueDateTime, arrivalDateTime, price) VALUES
                                                                                                           (1, 1, 2, '2024-10-10 15:00:00', '2024-10-10 18:00:00', 299.99),
                                                                                                           (2, 2, 3, '2024-10-11 14:00:00', '2024-10-11 20:00:00', 499.99),
                                                                                                           (3, 3, 4, '2024-10-12 13:00:00', '2024-10-12 19:00:00', 350.00),
                                                                                                           (4, 4, 5, '2024-10-13 12:00:00', '2024-10-13 21:00:00', 450.50),
                                                                                                           (5, 1, 1, '2024-10-14 09:00:00', '2024-10-14 12:00:00', 199.99);

INSERT INTO "Passenger" (firstName, lastName, email, phone, dateOfBirth, passportSerie, passportNumber, password) VALUES
                                                                                                                      ('John', 'Doe', 'john.doe@example.com', '1234567890', '1990-01-01', 'AA1234', '987654321', 'password123'),
                                                                                                                      ('Jane', 'Smith', 'jane.smith@example.com', '0987654321', '1992-02-02', 'BB5678', '123456789', 'password456'),
                                                                                                                      ('Alice', 'Johnson', 'alice.johnson@example.com', '1231231234', '1988-03-03', 'CC9876', '456789123', 'password789'),
                                                                                                                      ('Bob', 'Brown', 'bob.brown@example.com', '3213214321', '1995-04-04', 'DD5432', '321654987', 'password321'),
                                                                                                                      ('Charlie', 'Davis', 'charlie.davis@example.com', '7897897890', '1985-05-05', 'EE1357', '987123456', 'password654');

INSERT INTO "FlightBooking" (flightId, passengerId, bookingStatus, bookingDateTime, row, seatInRow) VALUES
                                                                                                        (1, 1, 'booked', '2024-10-01 10:00:00', 'A', 1),
                                                                                                        (2, 2, 'paid', '2024-10-02 11:00:00', 'B', 2),
                                                                                                        (3, 3, 'canceled', '2024-10-03 12:00:00', 'C', 3),
                                                                                                        (4, 4, 'done', '2024-10-04 13:00:00', 'D', 4),
                                                                                                        (5, 5, 'booked', '2024-10-05 14:00:00', 'E', 5);
