import React, { useState, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../css/Booking.css';

const Booking = () => {
    const location = useLocation();
    const navigate = useNavigate();
    const { flight } = location.state;
    const [user, setUser] = useState(null);
    const [selectedRow, setSelectedRow] = useState('');
    const [selectedSeat, setSelectedSeat] = useState('');
    const [seats, setSeats] = useState([]);
    const [error, setError] = useState('');
    const [isBooked, setIsBooked] = useState(false);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            axios
                .get('http://localhost:8080/auth/user', {
                    headers: { Authorization: `${token}` }
                })
                .then(response => {
                    setUser(response.data);
                })
                .catch(err => {
                    console.error('Error fetching user data:', err);
                    setError('Error fetching user data.');
                });
        }
    }, []);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (flight && flight.id && token) {
            axios
                .get(`http://localhost:8080/auth/flights/${flight.id}/seats`, {
                    headers: { Authorization: `${token}` }
                })
                .then(response => {
                    const seats = response.data;
                    setSeats(seats);
                })
                .catch(err => {
                    console.error('Error fetching available seats:', err);
                });

            axios
                .get(`http://localhost:8080/auth/flights/${flight.id}/isBooked`, {
                    headers: { Authorization: `${token}` }
                })
                .then(response => {
                    setIsBooked(response.data.isBooked);
                })
                .catch(err => {
                    console.error('Error checking booking status:', err);
                });
        }
    }, [flight.id]);

    const handleSeatSelection = (row, seat) => {
        setSelectedRow(row);
        setSelectedSeat(seat);
    };

    const rowLetters = [...new Set(seats.map(seat => seat.row))];

    const groupedSeats = rowLetters.map(row => {
        return {
            row: row,
            seats: seats.filter(seat => seat.row === row).sort((a, b) => a.seats - b.seats)
        };
    });

    const handleBooking = () => {
        if (!selectedRow || !selectedSeat) {
            alert('Please select a row and seat.');
            return;
        }

        const token = localStorage.getItem('token');
        const bookingData = {
            flightId: flight.id,
            status: "booked",
            row: selectedRow,
            seat: selectedSeat,
        };

        console.log(bookingData);

        axios
            .post('http://localhost:8080/auth/book', bookingData, {
                headers: { Authorization: `${token}` }
            })
            .then(response => {
                alert('Booking successful!');
                navigate('/');
            })
            .catch(err => {
                console.error('Error booking flight:', err);
                alert('Error booking flight.');
            });
    };

    const handleEditClick = () => {
        navigate('/edit-profile');
    };

    const [showDetails, setShowDetails] = useState(true);

    const toggleDetails = () => {
        setShowDetails(prevState => !prevState);
    };

    if (error) return <p className="text-center text-danger">{error}</p>;

    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4">Book Your Flight</h1>
            {user && (
                <div className="mb-4 bg-light border border-3 rounded-2 position-relative p-3">
                    <div className="d-flex justify-content-between align-items-start">
                        <h5 className="pt-2 pl-2">User Information</h5>

                        <div>
                            <button
                                className="btn btn-outline-warning me-2"
                                style={{marginTop: '5px'}}
                                onClick={handleEditClick}
                            >
                                Edit
                            </button>

                            <button
                                className="btn btn-outline-warning"
                                style={{marginTop: '5px'}}
                                onClick={toggleDetails}
                            >
                                {showDetails ? 'Hide' : 'Show'} Details
                            </button>
                        </div>
                    </div>

                    {showDetails && (
                        <div>
                            <p><strong>Name:</strong> {user.firstName} {user.lastName}</p>
                            <p><strong>Email:</strong> {user.email}</p>
                            <p><strong>Phone:</strong> {user.phone}</p>
                            <p><strong>Date of Birth:</strong> {new Date(user.dateOfBirth).toLocaleDateString()}</p>
                            <p><strong>Passport Serie:</strong> {user.passportSerie}</p>
                            <p><strong>Passport Number:</strong> {user.passportNumber}</p>
                        </div>
                    )}
                </div>
            )}

            <div className="mb-4 bg-light border border-3 rounded-2 position-relative p-3">
                <h5>Flight Information</h5>
                <p>
                    <strong>Departure:</strong> {flight.departure_city} ({flight.departure}) {new Date(flight.departure_date).toLocaleString()}
                </p>
                <p>
                    <strong>Arrival:</strong> {flight.arrival_city} ({flight.arrival}) {new Date(flight.arrival_date).toLocaleString()}
                </p>
                <p><strong>Price:</strong> {flight.price}</p>
                {isBooked && <p className="text-danger">This flight is already booked.</p>} {/* Show message if booked */}
            </div>

            <div className="mb-4 bg-light border border-3 rounded-2 position-relative p-3">
                <h5>Select Your Seat</h5>
                <div className="seats-wrapper"
                     style={{maxWidth: '100%', maxHeight: '300px', overflowX: 'auto', whiteSpace: 'nowrap'}}>
                    {groupedSeats.map((group, groupIndex) => (
                        <div key={groupIndex} className="d-flex align-items-start mb-1" style={{flexWrap: 'nowrap'}}>
                            <div style={{width: '30px', textAlign: 'center', marginRight: '10px'}}>
                                {group.row}
                            </div>
                            <div className="d-flex" style={{gap: '2px', whiteSpace: 'nowrap'}}>
                                {group.seats.map((seat, seatIndex) => (
                                    <button
                                        key={seatIndex}
                                        className={`seat-btn btn m-1
                                                ${seat.status === 'available' ? 'available' : 'unavailable'}
                                                ${selectedRow === seat.row && selectedSeat === seat.seat ? 'active' : ''}`}
                                        onClick={() => handleSeatSelection(seat.row, seat.seat)}
                                        disabled={seat.status !== 'available' || isBooked}
                                    >
                                        {seat.seat}
                                    </button>
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            </div>

            <button
                className="btn btn-success"
                onClick={handleBooking}
                disabled={isBooked}
            >
                Confirm Booking
            </button>
        </div>
    );
};

export default Booking;
