import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './Header.css';
import reactLogo from '../../logo.svg';

const Header = () => {
    const navigate = useNavigate();
    const token = localStorage.getItem('token');
    const email = localStorage.getItem('email');

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    return (
        <header className="header">
            <img
                src={reactLogo}
                alt="React Logo"
                style={{ width: '50px', height: '50px' }}
            />
            {token ? (
                <div className="d-flex align-items-center">
                    <pre className="mb-0 me-3">{email}</pre>
                    <button onClick={handleLogout}
                            className="btn btn-warning">Logout</button>
                </div>
            ) : (
                <nav>
                    <Link to="/login" className="btn btn-warning" >Login</Link>
                    <Link to="/signup" className="btn btn-warning" >Signup</Link>
                </nav>
            )}
        </header>
    );
};

export default Header;
