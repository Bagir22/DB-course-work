import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useLocation } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

const Search = () => {
    const location = useLocation();
    const [flights, setFlights] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const queryParams = new URLSearchParams(location.search);
        const dep = queryParams.get('dep');
        const des = queryParams.get('des');
        const depDate = queryParams.get('depDate');

        const fetchFlights = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/search?dep=${dep}&des=${des}&depDate=${depDate}`);
                if (response.data.length === 0) {
                    setError('Flights not found');
                } else {
                    setFlights(response.data);
                }
            } catch (error) {
                console.error('Error: ', error);
                setError('Error fetching flights');
            } finally {
                setLoading(false);
            }
        };

        fetchFlights();
    }, [location.search]);

    if (loading) return <p className="text-center">Loading...</p>;
    if (error) return <p className="text-center text-danger">{error}</p>;

    const token = localStorage.getItem('token');

    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4">Available Flights</h1>
            {flights.length > 0 ? (
                <div className="d-flex flex-column align-items-center">
                    {flights.map(flight => (
                        <div key={flight.id} className="card mb-3" style={{width: '100%', maxWidth: '600px'}}>
                            <div className="card-body">
                                <h5 className="card-title">{flight.departure} to {flight.arrival}</h5>
                                <p className="card-text">
                                    Departure: {new Date(flight.departure_date).toLocaleString()}<br/>
                                    Price: {flight.price}
                                </p>
                                {token ? (
                                    <button className="btn btn-primary">Book Now</button>
                                ) : (
                                    <p className="text-muted">Please log in to book</p>
                                )}
                            </div>
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
