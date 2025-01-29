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
        passportNumber: '',
        image: '',
    });

    const [message, setMessage] = useState(null);
    const [errors, setErrors] = useState({});
    const [image, setImage] = useState(null);
    const [isChanged, setIsChanged] = useState(false);

    const formatDate = (isoDate) => {
        const date = new Date(isoDate);
        return date.toISOString().split('T')[0];
    };

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            axios
                .get('http://localhost:8080/auth/user', {
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
                        passportNumber: userData.passportNumber || '',
                        image: userData.image || '',
                    });
                })
                .catch((err) => {
                    console.error('Error fetching user data:', err);
                });
        }
    }, []);

    const handleChange = (e) => {
        const { name, value } = e.target;
        if (profile.hasOwnProperty(name)) {
            setProfile({
                ...profile,
                [name]: value,
            });
            validateField(name, value);
        } else {
            console.warn(`Unexpected form field: ${name}`);
        }

        if (user[name] !== value) {
            setIsChanged(true);
        } else {
            const isStillChanged = Object.keys(profile).some((key) => profile[key] !== user[key]);
            setIsChanged(isStillChanged);
        }

        validateField(name, value);
    };

    const validateField = (name, value) => {
        let error = '';
        if (name === 'email' && !/^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/.test(value)) {
            error = 'Please enter a valid email.';
        } else if (name === 'dateOfBirth' && new Date(value) > new Date()) {
            error = 'Date of birth cannot be in the future.';
        } else if ((name === 'phone' || name === 'passportNumber') && !/^\d+$/.test(value)) {
            error = 'Only numbers are allowed.';
        }
        setErrors((prevErrors) => ({ ...prevErrors, [name]: error }));
    };

    const handleFileChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            const imageUrl = URL.createObjectURL(file);

            setProfile((prevProfile) => {
                const updatedProfile = { ...prevProfile, image: imageUrl };
                checkChanges(updatedProfile);
                return updatedProfile;
            });
            setImage(file);
        }
    };

    const checkChanges = (updatedProfile) => {
        if (!user) return;

        const isChanged = Object.keys(updatedProfile).some((key) => {
            return updatedProfile[key] !== user[key];
        });

        setIsChanged(isChanged);
    };

    const handleSave = () => {
        const formData = new FormData();

        if (isChanged) {
            Object.entries(profile).forEach(([key, value]) => {
                formData.append(key, value);
            });
        }

        if (image) {
            formData.append('image', image);
        } else {
            formData.append('image', profile.image);
        }

        formData.forEach((value, key) => {
            console.log(`${key}: ${value}`);
        });

        const token = localStorage.getItem('token');
        axios
            .put('http://localhost:8080/auth/user', formData, {
                headers: {
                    Authorization: `Bearer ${token}`,
                    'Content-Type': 'multipart/form-data',
                }
            })
            .then((response) => {
                const newToken = response.data.token;
                if (newToken) {
                    localStorage.setItem('token', newToken);
                }
                setMessage({ type: 'success', text: 'Profile updated successfully!' });
                if (image) {
                    setProfile((prevState) => ({
                        ...prevState,
                        image: image.name,
                    }));
                    localStorage.setItem('userImage', image.name);
                    console.log(image.name)
                }
                setIsChanged(false);
            })
            .catch((err) => {
                setMessage({ type: 'error', text: 'Error updating profile.' });
                console.log(err);
            });
    };

    if (!user) {
        return <div>Loading...</div>;
    }

    return (
        <div className="container mt-5">
            <h1 className="mb-4">Edit Profile</h1>

            {message && (
                <div className={`alert ${message.type === 'error' ? 'alert-danger' : 'alert-success'}`}>
                    {message.text}
                </div>
            )}

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
                            className={`form-control ${errors.email ? 'is-invalid' : ''}`}
                            name="email"
                            value={profile.email}
                            onChange={handleChange}
                        />
                        {errors.email && <div className="invalid-feedback">{errors.email}</div>}
                    </div>
                    <div className="col-md-6">
                        <label className="form-label">Phone</label>
                        <input
                            type="tel"
                            className={`form-control ${errors.phone ? 'is-invalid' : ''}`}
                            name="phone"
                            value={profile.phone}
                            onChange={handleChange}
                        />
                        {errors.phone && <div className="invalid-feedback">{errors.phone}</div>}
                    </div>
                </div>

                <div className="row mb-3">
                    <div className="col-md-6">
                        <label className="form-label">Date of Birth</label>
                        <input
                            type="date"
                            className={`form-control ${errors.dateOfBirth ? 'is-invalid' : ''}`}
                            name="dateOfBirth"
                            value={profile.dateOfBirth}
                            onChange={handleChange}
                        />
                        {errors.dateOfBirth && <div className="invalid-feedback">{errors.dateOfBirth}</div>}
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
                            className={`form-control ${errors.passportNumber ? 'is-invalid' : ''}`}
                            name="passportNumber"
                            value={profile.passportNumber}
                            onChange={handleChange}
                        />
                        {errors.passportNumber && <div className="invalid-feedback">{errors.passportNumber}</div>}
                    </div>
                </div>

                <div className="mb-3">
                    <label className="form-label">Profile Image</label>
                    <input
                        type="file"
                        className="form-control"
                        onChange={handleFileChange}
                    />
                    {image && <small className="text-muted">Selected: {image.name}</small>}
                </div>

                {profile.image && (
                    <div className="mb-3">
                        <img
                            src={profile.image.startsWith('blob:')
                                ? profile.image
                                : `http://localhost:8080/uploads/${profile.image}`}
                            alt="Profile"
                            className="img-thumbnail"
                            style={{maxWidth: '200px'}}
                        />
                    </div>
                )}

                <button type="button" className="btn btn-success" onClick={handleSave} disabled={!isChanged} >
                    Save
                </button>
            </form>
        </div>
    );
};

export default EditProfile;
