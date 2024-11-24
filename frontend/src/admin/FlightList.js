import React, { useEffect, useState } from 'react';
import axios from 'axios';
import FlightForm from './Add';
import { useNavigate } from 'react-router-dom';

const FlightList = () => {
    const [flights, setFlights] = useState([]);
    const [selectedFlight, setSelectedFlight] = useState(null);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        axios.get('http://localhost:8080/admin/flights')
            .then((response) => {
                setFlights(response.data);
            })
            .catch((error) => {
                console.error('Error fetching flights:', error);
            });
    }, []);

    const redirectToEditPage = (flightId) => {
        navigate(`/admin/edit-flight/${flightId}`);
    };

    const redirectToAddPage = () => {
        navigate(`/admin/add/`);
    };

    const formatDate = (isoDate) => {
        const date = new Date(isoDate);
        return date.toISOString().split('T')[0] + " " + date.toISOString().split('T')[1].split('.')[0];
    };

    const deleteFlight = (flightId) => {
        axios.delete(`http://localhost:8080/admin/flights/${flightId}`)
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
                    <th>Price</th>
                    <th>Actions</th>
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
                            <td>{flight.price}</td>
                            <td>
                                {!isPastDeparture && (
                                    <>
                                        <button
                                            className="btn btn-warning"
                                            onClick={() => redirectToEditPage(flight.id)}
                                        >
                                            Edit
                                        </button>
                                        <button className="btn btn-danger"
                                                onClick={() => deleteFlight(flight.id)}>Delete
                                        </button>
                                    </>
                                )}
                            </td>
                        </tr>
                    );
                })}
                </tbody>
            </table>
        </div>
    );
};

export default FlightList;
