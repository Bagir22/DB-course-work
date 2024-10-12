import React, { useState, useEffect } from 'react';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css'; // Убедитесь, что Bootstrap импортирован

const FlightHistory = () => {
    const [history, setHistory] = useState([]);

    useEffect(() => {
        const token = localStorage.getItem('token');
        axios
            .get('http://localhost:8080/api/history', {
                headers: { Authorization: `${token}` },
            })
            .then((response) => {
                setHistory(response.data);
            })
            .catch((error) => {
                console.error('Error fetching flight history:', error);
            });
    }, []);

    return (
        <div className="container mt-4">
            <h2 className="mb-4">Flight History</h2>
            <table className="table table-striped table-bordered">
                <thead className="thead-dark">
                <tr>
                    <th>Flight ID</th>
                    <th>Departure</th>
                    <th>Arrival</th>
                    <th>Price</th>
                    <th>Status</th>
                </tr>
                </thead>
                <tbody>
                {history.map((flight, index) => (
                    <tr key={flight.flightId || index}>
                        <td>{flight.flightId}</td>
                        <td>{flight.departure}</td>
                        <td>{flight.arrival}</td>
                        <td>{flight.price}</td>
                        <td>{flight.status}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default FlightHistory;
