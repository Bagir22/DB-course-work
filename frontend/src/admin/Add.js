import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

const AddFlightPage = () => {
    const [flight, setFlight] = useState({
        airline_id: '',
        aircraft_id: '',
        departure_datetime: '',
        arrival_datetime: '',
        price: '',
        departure_id: '',
        arrival_id: '',
    });
    const [airlinesAndAircrafts, setAirlinesAndAircrafts] = useState([]);
    const [airlines, setAirlines] = useState([]);
    const [filteredAircrafts, setFilteredAircrafts] = useState([]);
    const [airports, setAirports] = useState([]);
    const navigate = useNavigate();
    const currentDate = new Date().toISOString().slice(0, 16);

    useEffect(() => {
        const token = localStorage.getItem('token');
        axios.get(`http://localhost:8080/auth/airlinesaircrafts`, {
            headers: { Authorization: `${token}` }
            })
            .then((response) => {
                const data = response.data;

                const airlinesGrouped = data.reduce((acc, item) => {
                    if (!acc[item.airline_id]) {
                        acc[item.airline_id] = [];
                    }
                    acc[item.airline_id].push(item);
                    return acc;
                }, {});

                setAirlinesAndAircrafts(airlinesGrouped);

                const uniqueAirlines = Object.keys(airlinesGrouped).map((key) => ({
                    airline_id: parseInt(key),
                    airline_name: airlinesGrouped[key][0].airline_name,
                }));
                setAirlines(uniqueAirlines);
            })
            .catch((error) => {
                console.error('Error loading airlines and aircrafts:', error);
            });

        axios.get('http://localhost:8080/auth/airports', {
            headers: { Authorization: `${token}` }
            })
            .then((response) => {
                setAirports(response.data);
            })
            .catch((error) => {
                console.error('Error loading airports:', error);
            });
    }, []);

    useEffect(() => {
        if (flight.airline_id) {
            const filtered = airlinesAndAircrafts[flight.airline_id] || [];
            setFilteredAircrafts(filtered);
        } else {
            setFilteredAircrafts([]);
        }
    }, [flight.airline_id, airlinesAndAircrafts]);

    const handleSave = () => {
        const token = localStorage.getItem('token');
        const formattedFlight = {
            ...flight,
            departure_datetime: new Date(flight.departure_datetime).toISOString(),
            arrival_datetime: new Date(flight.arrival_datetime).toISOString(),
        };

        console.log(flight);

        axios.post('http://localhost:8080/admin/flights', formattedFlight, {
            headers: { Authorization: `${token}` }
        })
            .then(() => {
                alert('Flight added successfully!');
                navigate('/admin');
            })
            .catch((error) => {
                console.error('Error adding flight:', error);
            });
    };

    return (
        <div className="container mt-5">
            <h1 className="mb-4 text-center">Add Flight</h1>
            <form className="bg-light p-4 rounded shadow">
                <div className="mb-3">
                    <label className="form-label">Airline Name</label>
                    <select
                        className="form-select"
                        value={flight.airline_id || ''}
                        onChange={(e) => {
                            const selectedAirlineId = parseInt(e.target.value);
                            setFlight({
                                ...flight,
                                airline_id: selectedAirlineId,
                                aircraft_id: '',
                            });
                        }}
                    >
                        <option value="">Select an airline</option>
                        {airlines.map((airline) => (
                            <option key={airline.airline_id} value={airline.airline_id}>
                                {airline.airline_name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="mb-3">
                    <label className="form-label">Aircraft Name</label>
                    <select
                        className="form-select"
                        value={flight.aircraft_id || ''}
                        onChange={(e) => {
                            const selectedAircraftId = parseInt(e.target.value);
                            setFlight({...flight, aircraft_id: selectedAircraftId});
                        }}
                        disabled={filteredAircrafts.length === 0}
                    >
                        <option value="">Select an aircraft</option>
                        {filteredAircrafts.map((aircraft) => (
                            <option key={aircraft.aircraft_id} value={aircraft.aircraft_id}>
                                {aircraft.aircraft_name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="mb-3">
                    <label className="form-label">Departure Airport</label>
                    <select
                        className="form-select"
                        value={flight.departure_id || ''}
                        onChange={(e) => setFlight({...flight, departure_id: parseInt(e.target.value)})}
                    >
                        <option value="">Select a departure airport</option>
                        {airports.map((airport) => (
                            <option key={airport.id} value={airport.id}>
                                {airport.name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="mb-3">
                    <label className="form-label">Arrival Airport</label>
                    <select
                        className="form-select"
                        value={flight.arrival_id || ''}
                        onChange={(e) => setFlight({...flight, arrival_id: parseInt(e.target.value)})}
                    >
                        <option value="">Select an arrival airport</option>
                        {airports.map((airport) => (
                            <option key={airport.id} value={airport.id}>
                                {airport.name}
                            </option>
                        ))}
                    </select>
                </div>

                <div className="mb-3">
                    <label className="form-label">Departure Date</label>
                    <input
                        type="datetime-local"
                        className="form-control"
                        value={flight.departure_datetime}
                        min={currentDate}
                        onChange={(e) => setFlight({...flight, departure_datetime: e.target.value})}
                    />
                </div>

                <div className="mb-3">
                    <label className="form-label">Arrival Date</label>
                    <input
                        type="datetime-local"
                        className="form-control"
                        value={flight.arrival_datetime}
                        min={currentDate}
                        onChange={(e) => setFlight({...flight, arrival_datetime: e.target.value})}
                    />
                </div>

                <div className="mb-3">
                    <label className="form-label">Price</label>
                    <input
                        type="number"
                        className="form-control"
                        value={flight.price || ''}
                        onChange={(e) => setFlight({...flight, price: parseInt(e.target.value)})}
                    />
                </div>

                <button
                    type="button"
                    className="btn btn-primary w-100"
                    onClick={handleSave}
                    disabled={
                        !flight.airline_id ||
                        !flight.aircraft_id ||
                        !flight.departure_datetime ||
                        !flight.arrival_datetime ||
                        !flight.price ||
                        !flight.departure_id ||
                        !flight.arrival_id ||
                        new Date(flight.arrival_datetime) <= new Date(flight.departure_datetime)
                    }
                >
                    Add Flight
                </button>
            </form>
        </div>
    );
};

export default AddFlightPage;
