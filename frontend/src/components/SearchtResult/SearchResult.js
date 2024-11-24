import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useLocation, useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

const Search = () => {
    const location = useLocation();
    const [flights, setFlights] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const queryParams = new URLSearchParams(location.search);
        const dep = queryParams.get('dep');
        const des = queryParams.get('des');
        const depDate = queryParams.get('depDate');

        const fetchFlights = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/search?dep=${dep}&des=${des}&depDate=${depDate}`);
                if (response.data.length === 0) {
                    console.error('No flights found');
                } else {
                    setFlights(response.data);
                }
            } catch (error) {
                console.error('Error: ', error);
            } finally {
                setLoading(false);
            }
        };

        fetchFlights();
    }, [location.search]);

    const token = localStorage.getItem('token');

    const handleBooking = (flight) => {
        if (token) {
            if (flight.isBooked) {
                alert('You have already booked this flight.');
            } else {
                navigate(`/book/${flight.id}`, { state: { flight } });
            }
        } else {
            alert('Please log in to book a flight');
            navigate('/login');
        }
    };

    if (loading) return <p className="text-center">Loading...</p>;
    if (error) return <p className="text-center text-danger">{error}</p>;

    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4">Available Flights</h1>
            {flights.length > 0 ? (
                <div className="d-flex flex-column align-items-center">
                    {flights.map(flight => (
                        <div key={flight.id} className="list-group-item mb-3 p-4 border rounded">
                            <h5>{flight.departure_city} ({flight.departure}) → {flight.arrival_city} ({flight.arrival})</h5>
                            <p><strong>Departure:</strong> {new Date(flight.departure_date).toLocaleString()}</p>
                            <p><strong>Arrival:</strong> {new Date(flight.arrival_date).toLocaleString()}</p>
                            <p><strong>Price:</strong> {flight.price}</p>
                            <p><strong>Available Seats:</strong> {flight.available_seats}</p>
                            {flight.isBooked ? (
                                <button className="btn btn-secondary" disabled>
                                    Already Booked
                                </button>
                            ) : (
                                token ? (
                                    <button className="btn btn-info" onClick={() => handleBooking(flight)}>
                                        Book Now
                                    </button>
                                ) : (
                                    <p className="text-muted">Please log in to book</p>
                                )
                            )}
                        </div>
                    ))}
                </div>
            ) : (
                <p className="text-center">No flights found</p>
            )}
        </div>
    );
};

export default Search;
