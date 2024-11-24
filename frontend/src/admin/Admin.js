import React, {useEffect} from 'react';
import FlightList from './FlightList';
import { useNavigate } from 'react-router-dom';

const Admin = () => {
    const navigate = useNavigate();

    useEffect(() => {
        const isAdmin = localStorage.getItem('userIsAdmin');

        if (!isAdmin) {
            navigate('/login');
        }
    }, [navigate]);

    return (
        <div className="container">
            <h1 className="mt-3">Flights:</h1>
            <FlightList />
        </div>
    );
};

export default Admin;