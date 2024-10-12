import React, { useState, useEffect } from 'react';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';

const EditProfile = () => {
    const [user, setUser] = useState(null);
    const [profile, setProfile] = useState({
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        dateOfBirth: '',
        passportSerie: '',
        passportNumber: ''
    });

    const formatDate = (isoDate) => {
        const date = new Date(isoDate);
        return date.toISOString().split('T')[0];
    };

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            axios
                .get('http://localhost:8080/api/user', {
                    headers: { Authorization: `Bearer ${token}` }
                })
                .then((response) => {
                    const userData = response.data;
                    setUser(userData);
                    setProfile({
                        firstName: userData.firstName || '',
                        lastName: userData.lastName || '',
                        email: userData.email || '',
                        phone: userData.phone || '',
                        dateOfBirth: formatDate(userData.dateOfBirth) || '',
                        passportSerie: userData.passportSerie || '',
                        passportNumber: userData.passportNumber || ''
                    });
                })
                .catch((err) => {
                    console.error('Error fetching user data:', err);
                });
        }
    }, []);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setProfile({
            ...profile,
            [name]: value
        });
    };

    const handleSave = () => {
        const token = localStorage.getItem('token');
        axios
            .put('http://localhost:8080/api/user', profile, {
                headers: { Authorization: `Bearer ${token}` }
            })
            .then((response) => {
                console.log('Profile updated successfully:', response.data);
            })
            .catch((err) => {
                console.error('Error updating profile:', err);
            });
    };

    if (!user) {
        return <div>Loading...</div>;
    }

    return (
        <div className="container mt-5">
            <h1 className="mb-4">Edit Profile</h1>
            <form>
                <div className="row mb-3">
                    <div className="col-md-6">
                        <label className="form-label">First Name</label>
                        <input
                            type="text"
                            className="form-control"
                            name="firstName"
                            value={profile.firstName}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="col-md-6">
                        <label className="form-label">Last Name</label>
                        <input
                            type="text"
                            className="form-control"
                            name="lastName"
                            value={profile.lastName}
                            onChange={handleChange}
                        />
                    </div>
                </div>

                <div className="row mb-3">
                    <div className="col-md-6">
                        <label className="form-label">Email</label>
                        <input
                            type="email"
                            className="form-control"
                            name="email"
                            value={profile.email}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="col-md-6">
                        <label className="form-label">Phone</label>
                        <input
                            type="tel"
                            className="form-control"
                            name="phone"
                            value={profile.phone}
                            onChange={handleChange}
                        />
                    </div>
                </div>

                <div className="row mb-3">
                    <div className="col-md-6">
                        <label className="form-label">Date of Birth</label>
                        <input
                            type="date"
                            className="form-control"
                            name="dateOfBirth"
                            value={profile.dateOfBirth}
                            onChange={handleChange}
                        />
                    </div>
                </div>

                <div className="row mb-3">
                    <div className="col-md-6">
                        <label className="form-label">Passport Series</label>
                        <input
                            type="text"
                            className="form-control"
                            name="passportSerie"
                            value={profile.passportSerie}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="col-md-6">
                        <label className="form-label">Passport Number</label>
                        <input
                            type="text"
                            className="form-control"
                            name="passportNumber"
                            value={profile.passportNumber}
                            onChange={handleChange}
                        />
                    </div>
                </div>

                <button type="button" className="btn btn-success" onClick={handleSave}>
                    Save
                </button>
            </form>
        </div>
    );
};

export default EditProfile;
