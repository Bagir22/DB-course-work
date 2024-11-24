import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import './Header.css';
import reactLogo from '../../logo.svg';

const Header = () => {
    const navigate = useNavigate();
    const token = localStorage.getItem('token');
    const email = localStorage.getItem('email');
    const [showDropdown, setShowDropdown] = useState(false);
    const [userProfileImage, setUserProfileImage] = useState(`http://localhost:8080/uploads/${localStorage.getItem('userImage')}` || 'http://localhost:8080/uploads/defaultImage.jpeg');

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    return (
        <header className="header d-flex align-items-center justify-content-between p-3">
            <img
                src={reactLogo}
                alt="React Logo"
                style={{width: '50px', height: '50px', cursor: 'pointer'}}
                onClick={() => navigate('/')}
            />
            {token ? (
                <div className="d-flex align-items-center">
                    <div
                        className="position-relative d-flex align-items-center"
                        onMouseEnter={() => setShowDropdown(true)}
                        onMouseLeave={() => setShowDropdown(false)}
                        style={{cursor: 'pointer'}}
                    >
                        <img
                            src={userProfileImage}
                            alt="Profile"
                            style={{
                                width: '40px',
                                height: '40px',
                                borderRadius: '50%',
                                objectFit: 'cover',
                                marginRight: '10px'
                            }}
                        />
                        <pre className="mb-0 me-3">{email}</pre>
                        {showDropdown && (
                            <div className="dropdown-menu show p-2"
                                 style={{position: 'absolute', top: '100%', right: '0'}}>
                                <button
                                    className="dropdown-item btn btn-link text-decoration-none"
                                    onClick={() => navigate('/edit-profile')}
                                >
                                    Edit Profile
                                </button>
                                <button
                                    className="dropdown-item btn btn-link text-decoration-none"
                                    onClick={() => navigate('/history')}
                                >
                                    History
                                </button>
                            </div>
                        )}
                    </div>
                    <button onClick={handleLogout} className="btn btn-warning ms-3">Logout</button>
                </div>
            ) : (
                <nav>
                    <Link to="/login" className="btn btn-warning me-2">Login</Link>
                    <Link to="/signup" className="btn btn-warning">Signup</Link>
                </nav>
            )}
        </header>
    );
};

export default Header;
