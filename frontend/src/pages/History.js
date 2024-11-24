import React, { useState, useEffect } from 'react';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';

const FlightHistory = () => {
    const [history, setHistory] = useState([]);
    const [totalFlights, setTotalFlights] = useState(0);
    const [totalPrice, setTotalPrice] = useState(0);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    const [statusFilter, setStatusFilter] = useState('');
    const [cityFilter, setCityFilter] = useState('');
    const [dateFilter, setDateFilter] = useState('');

    const [sortField, setSortField] = useState('');
    const [sortDirection, setSortDirection] = useState('asc');

    const fetchFlightHistory = () => {
        const token = localStorage.getItem('token');
        setLoading(true);

        const params = {
            status: statusFilter,
            city: cityFilter,
            date: dateFilter,
        };

        axios
            .get('http://localhost:8080/api/history', {
                headers: { Authorization: `Bearer ${token}` },
                params,
            })
            .then((response) => {
                const flights = response.data || [];
                setHistory(flights);
                setTotalFlights(flights.length);

                const calculatedTotalPrice = flights.reduce((acc, flight) => {
                    const price = parseFloat(flight.price.replace('$', ''));
                    return !isNaN(price) ? acc + price : acc;
                }, 0);

                setTotalPrice(calculatedTotalPrice);
                setError('');
            })
            .catch((error) => {
                console.error('Error fetching flight history:', error);
                setError('Failed to load flight history.');
            })
            .finally(() => {
                setLoading(false);
            });
    };

    useEffect(() => {
        fetchFlightHistory();
    }, []);

    const handleFilterSubmit = (e) => {
        e.preventDefault();
        fetchFlightHistory();
    };

    const handleCancelFlight = (flightId) => {
        const token = localStorage.getItem('token');

        axios
            .post(`http://localhost:8080/api/cancel/${flightId}`, {}, {
                headers: { Authorization: `Bearer ${token}` },
            })
            .then(() => {
                setHistory((prevHistory) =>
                    prevHistory.map((flight) =>
                        flight.id === flightId ? { ...flight, status: 'canceled' } : flight
                    )
                );
            })
            .catch((error) => {
                console.error('Error canceling flight:', error);
                setError('Failed to cancel flight.');
            });
    };

    const handleSort = (field) => {
        const direction = sortField === field && sortDirection === 'asc' ? 'desc' : 'asc';
        setSortField(field);
        setSortDirection(direction);

        const sorted = [...history].sort((a, b) => {
            if (field === 'price') {
                const priceA = parseFloat(a.price.replace('$', '')) || 0;
                const priceB = parseFloat(b.price.replace('$', '')) || 0;
                return direction === 'asc' ? priceA - priceB : priceB - priceA;
            } else if (field === 'departure_date') {
                const dateA = new Date(a[field]);
                const dateB = new Date(b[field]);
                return direction === 'asc' ? dateA - dateB : dateB - dateA;
            } else {
                const valueA = a[field]?.toString().toLowerCase() || '';
                const valueB = b[field]?.toString().toLowerCase() || '';
                return direction === 'asc' ? valueA.localeCompare(valueB) : valueB.localeCompare(valueA);
            }
        });

        setHistory(sorted);
    };

    if (loading) {
        return <div className="text-center">Loading...</div>;
    }

    return (
        <div className="container mt-4">
            <h2 className="mb-4">Flight History</h2>
            <div className="mb-3">
                <strong>Total Flights:</strong> {totalFlights} <br />
                <strong>Total Price:</strong> ${totalPrice.toFixed(2)}
            </div>

            <form onSubmit={handleFilterSubmit} className="mb-3">
                <input
                    type="text"
                    placeholder="Status"
                    value={statusFilter}
                    onChange={(e) => setStatusFilter(e.target.value)}
                    className="form-control mb-2"
                />
                <input
                    type="text"
                    placeholder="City"
                    value={cityFilter}
                    onChange={(e) => setCityFilter(e.target.value)}
                    className="form-control mb-2"
                />
                <input
                    type="date"
                    value={dateFilter}
                    onChange={(e) => setDateFilter(e.target.value)}
                    className="form-control mb-2"
                />
                <button type="submit" className="btn btn-primary">
                    Apply Filters
                </button>
            </form>

            {error && <div className="text-danger">{error}</div>}

            <table className="table table-striped table-bordered">
                <thead className="thead-dark">
                <tr>
                    <th onClick={() => handleSort('departure')} style={{ cursor: 'pointer' }}>
                        Departure {sortField === 'departure' && (sortDirection === 'asc' ? '↑' : '↓')}
                    </th>
                    <th onClick={() => handleSort('arrival')} style={{ cursor: 'pointer' }}>
                        Arrival {sortField === 'arrival' && (sortDirection === 'asc' ? '↑' : '↓')}
                    </th>
                    <th onClick={() => handleSort('price')} style={{ cursor: 'pointer' }}>
                        Price {sortField === 'price' && (sortDirection === 'asc' ? '↑' : '↓')}
                    </th>
                    <th onClick={() => handleSort('departure_date')} style={{ cursor: 'pointer' }}>
                        Date {sortField === 'departure_date' && (sortDirection === 'asc' ? '↑' : '↓')}
                    </th>
                    <th>Status</th>
                    <th>Action</th>
                </tr>
                </thead>
                <tbody>
                {Array.isArray(history) && history.length > 0 ? (
                    history.map((flight) => (
                        <tr key={flight.id}>
                            <td>{flight.departure}</td>
                            <td>{flight.arrival}</td>
                            <td>{flight.price}</td>
                            <td>{new Date(flight.departure_date).toISOString().split('T')[0]}</td>
                            <td className={flight.status === 'done' || flight.status === 'canceled' ? 'text-muted' : ''}>
                                {flight.status}
                            </td>
                            <td>
                                {flight.status === 'booked' && (
                                    <button
                                        className="btn btn-danger btn-sm"
                                        onClick={() => handleCancelFlight(flight.id)}
                                    >
                                        Cancel
                                    </button>
                                )}
                            </td>
                        </tr>
                    ))
                ) : (
                    <tr>
                        <td colSpan="6" className="text-center">No flight history available.</td>
                    </tr>
                )}
                </tbody>
            </table>
        </div>
    );
};

export default FlightHistory;
