import React, { useEffect, useState } from 'react';
import axios from 'axios';
import FlightForm from './Add';
import { useNavigate } from 'react-router-dom';


const FlightList = () => {
    const [flights, setFlights] = useState([]);
    const [limit, setLimit] = useState(10);
    const [offset, setOffset] = useState(0);
    const [totalFlights, setTotalFlights] = useState(0);
    const navigate = useNavigate();

    useEffect(() => {
        axios
            .get(`http://localhost:8080/admin/flights?limit=${limit}&offset=${offset}`)
            .then((response) => {
                setFlights(response.data || []);
            })
            .catch((error) => {
                console.error('Error fetching flights:', error);
                setFlights([]);
            });
    }, [limit, offset]);

    useEffect(() => {
        axios
            .get(`http://localhost:8080/admin/flightscount`)
            .then((response) => {
                setTotalFlights(response.data || 0);
                console.log(totalFlights)
            })
            .catch((error) => {
                console.error('Error fetching total flights count:', error);
                setTotalFlights(0);
            });
    }, []);

    const formatDate = (isoDate) => {
        const date = new Date(isoDate);
        return date.toISOString().split('T')[0] + " " + date.toISOString().split('T')[1].split('.')[0];
    };

    const redirectToAddPage = () => {
        navigate(`/admin/add/`);
    };

    const handleLimitChange = (newLimit) => {
        setLimit(newLimit);
        setOffset(0);
    };

    const handlePageChange = (direction) => {
        const newOffset = offset + direction * limit;
        if (newOffset >= 0 && newOffset < totalFlights) {
            setOffset(newOffset);
        }
    };

    const redirectToEditPage = (flightId) => {
        navigate(`/admin/edit-flight/${flightId}`);
    };

    const deleteFlight = (flightId) => {
        axios
            .delete(`http://localhost:8080/admin/flights/${flightId}`)
            .then(() => {
                setFlights(flights.filter((flight) => flight.id !== flightId));
            })
            .catch((error) => {
                console.error('Error deleting flight:', error);
            });
    };

    return (
        <div>
            <button className="btn btn-primary" onClick={() => redirectToAddPage()}>Add New Flight</button>

            <div className="mt-3">
                <span>Show:</span>
                <button className="btn btn-secondary m-1" onClick={() => handleLimitChange(2)}>2</button>
                <button className="btn btn-secondary m-1" onClick={() => handleLimitChange(5)}>5</button>
                <button className="btn btn-secondary m-1" onClick={() => handleLimitChange(10)}>10</button>
            </div>

            <table className="table mt-3">
                <thead>
                <tr>
                    <th>Flight ID</th>
                    <th>Aircraft</th>
                    <th>Airline</th>
                    <th>Departure Airport</th>
                    <th>Departure Datetime</th>
                    <th>Arrival Airport</th>
                    <th>Arrival Datetime</th>
                    <th>Count</th>
                    <th>Price</th>
                </tr>
                </thead>
                <tbody>
                {flights.map((flight) => {
                    const now = new Date();
                    const departureTime = new Date(flight.departure_datetime);
                    const isPastDeparture = now.getTime() >= departureTime;

                    return (
                        <tr key={flight.id}>
                            <td>{flight.id}</td>
                            <td>{flight.aircraft_name}</td>
                            <td>{flight.airline_name}</td>
                            <td>{flight.departure_airport}</td>
                            <td>{formatDate(flight.departure_datetime)}</td>
                            <td>{flight.destination_airport}</td>
                            <td>{formatDate(flight.arrival_datetime)}</td>
                            <td>{flight.booking_count}</td>
                            <td>{flight.price}</td>
                            <td>
                                {!isPastDeparture && (
                                    <button
                                        className="btn btn-warning m-1"
                                        onClick={() => redirectToEditPage(flight.id)}
                                    >
                                        Edit
                                    </button>
                                )}

                                {!isPastDeparture && (
                                    <button
                                        className="btn btn-danger m-1"
                                        onClick={() => deleteFlight(flight.id)}
                                        disabled={flight.booking_count !== 0}
                                    >
                                        Delete
                                    </button>
                                )}
                            </td>
                        </tr>
                    );
                })}
                </tbody>
            </table>

            <div className="pagination-controls mt-3">
                <button
                    className="btn btn-primary me-2"
                    onClick={() => handlePageChange(-1)}
                    disabled={offset === 0}
                >
                    Previous
                </button>
                <span>Page {Math.ceil(offset / limit) + 1} of {Math.ceil(totalFlights / limit)}</span>
                <button
                    className="btn btn-primary ms-2"
                    onClick={() => handlePageChange(1)}
                    disabled={offset + limit >= totalFlights}
                >
                    Next
                </button>
            </div>
        </div>
    );
};

export default FlightList;
