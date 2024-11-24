import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

const EditFlightPage = () => {
    const { id } = useParams();
    const [flight, setFlight] = useState(null);
    const [airlinesAndAircrafts, setAirlinesAndAircrafts] = useState([]);
    const [airlines, setAirlines] = useState([]);
    const [filteredAircrafts, setFilteredAircrafts] = useState([]);
    const currentDate = new Date().toISOString().slice(0, 16);


    useEffect(() => {
        axios.get(`http://localhost:8080/api/airlinesaircrafts`)
            .then((response) => {
                const data = response.data;

                const airlinesGrouped = data.reduce((acc, item) => {
                    if (!acc[item.airline_name]) {
                        acc[item.airline_name] = [];
                    }
                    acc[item.airline_name].push(item);
                    return acc;
                }, {});

                setAirlinesAndAircrafts(airlinesGrouped);

                setAirlines(Object.keys(airlinesGrouped));
            })
            .catch((error) => {
                console.error('Error loading airlines and aircrafts:', error);
            });

        axios.get(`http://localhost:8080/admin/flights/${id}`)
            .then((response) => {
                setFlight(response.data);
            })
            .catch((error) => {
                console.error('Error loading flight data:', error);
            });
    }, [id]);

    useEffect(() => {
        if (flight?.airline_name) {
            const filtered = airlinesAndAircrafts[flight.airline_name] || [];
            setFilteredAircrafts(filtered);
        } else {
            setFilteredAircrafts([]); // если авиакомпания не выбрана
        }
    }, [flight?.airline_name, airlinesAndAircrafts]);

    const handleSave = () => {
        axios.put(`http://localhost:8080/admin/flights/${id}`, {
            ...flight,
            aircraft_id: parseInt(flight.aircraft_id)
        })
            .then(() => {
                alert('Flight updated successfully!');
            })
            .catch((error) => {
                console.error('Error updating flight:', error);
            });
    };

    if (!flight || !airlines.length || !Object.keys(airlinesAndAircrafts).length) {
        return <div className="text-center mt-5">Loading...</div>;
    }

    return (
        <div className="container mt-5">
            <h1 className="mb-4 text-center">Edit Flight</h1>
            <form className="bg-light p-4 rounded shadow">
                <div className="mb-3">
                    <label className="form-label">Airline Name</label>
                    <select
                        className="form-select"
                        value={flight.airline_name || ''}
                        onChange={(e) => {
                            const selectedAirlineName = e.target.value;
                            setFlight({ ...flight, airline_name: selectedAirlineName, aircraft_id: '' });
                        }}
                    >
                        <option value="">Select an airline</option>
                        {airlines.map((airline, index) => (
                            <option key={index} value={airline}>
                                {airline}
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
                            setFlight({ ...flight, aircraft_id: selectedAircraftId });
                        }}
                        disabled={filteredAircrafts.length === 0}
                    >
                        <option value="">Select an aircraft</option>
                        {filteredAircrafts.length > 0 ? (
                            filteredAircrafts.map((aircraft) => (
                                <option key={aircraft.aircraft_id} value={aircraft.aircraft_id}>
                                    {aircraft.aircraft_name}
                                </option>
                            ))
                        ) : (
                            <option value="">No available aircraft</option>
                        )}
                    </select>
                </div>

                <div className="mb-3">
                    <label className="form-label">Departure Date</label>
                    <input
                        type="datetime-local"
                        className="form-control"
                        value={new Date(flight.departure_datetime).toISOString().slice(0, 16)}
                        min={currentDate}
                        onChange={(e) => setFlight({ ...flight, departure_datetime: e.target.value })}
                    />
                </div>
                <div className="mb-3">
                    <label className="form-label">Arrival Date</label>
                    <input
                        type="datetime-local"
                        className="form-control"
                        value={new Date(flight.arrival_datetime).toISOString().slice(0, 16)}
                        min={currentDate}
                        onChange={(e) => setFlight({ ...flight, arrival_datetime: e.target.value })}
                    />
                </div>
                <div className="mb-3">
                    <label className="form-label">Price</label>
                    <input
                        type="number"
                        className="form-control"
                        value={flight.price || ''}
                        onChange={(e) => setFlight({ ...flight, price: parseInt(e.target.value) })}
                    />
                </div>
                <button type="button" className="btn btn-primary w-100" onClick={handleSave}>
                    Save
                </button>
            </form>
        </div>
    );
};

export default EditFlightPage;
