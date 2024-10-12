import React, { useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Home = () => {
    const [departure, setDeparture] = useState('');
    const [arrival, setArrival] = useState('');
    const [date, setDate] = useState('');
    const navigate = useNavigate();
    const currentDate = new Date().toISOString().split('T')[0];

    const handleSearch = async (e) => {
        e.preventDefault();
        navigate(`/search?dep=${departure}&des=${arrival}&depDate=${date}`);
    };


    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4">Search Flights</h1>
            <form
                onSubmit={handleSearch}
                className="d-flex justify-content-center align-items-center"
                style={{ maxWidth: '600px', margin: '0 auto' }}
            >
                <div className="mb-0 flex-grow-1 me-2">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Departue"
                        value={departure}
                        onChange={(e) => setDeparture(e.target.value)}
                        required
                        style={{ height: '50px' }}
                    />
                </div>
                <div className="mb-0 flex-grow-1 me-2">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Destination"
                        value={arrival}
                        onChange={(e) => setArrival(e.target.value)}
                        required
                        style={{ height: '50px' }}
                    />
                </div>
                <div className="mb-0 flex-grow-1 me-2">
                    <input
                        type="date"
                        className="form-control"
                        value={date}
                        onChange={(e) => setDate(e.target.value)}
                        required
                        style={{ height: '50px' }}
                        min={currentDate}
                    />
                </div>
                <button type="submit" className="btn btn-warning" style={{ height: '50px', marginBottom: 0 }}>
                    Search
                </button>
            </form>
        </div>
    );
};

export default Home;
