import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Signup = () => {
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [phone, setPhone] = useState('');
    const [dateOfBirth, setDateOfBirth] = useState('');
    const [passportSerie, setPassportSerie] = useState('');
    const [passportNumber, setPassportNumber] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSignup = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch('http://localhost:8080/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    firstName,
                    lastName,
                    email,
                    phone,
                    dateOfBirth,
                    passportSerie,
                    passportNumber,
                    password,
                }),
            });

            if (response.ok) {
                const data = await response.json();
                alert("Success");
                navigate('/login');
            } else {
                const errorData = await response.json();
                alert(errorData.error);
            }
        } catch (error) {
            console.error('Signup error:', error);
            alert('An error occurred. Please try again.');
        }
    };

    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4" style={{fontFamily: "'Roboto', sans-serif"}}>
                Signup
            </h1>
            <form onSubmit={handleSignup} className="mx-auto" style={{maxWidth: '400px'}}>
                <div className="mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="First Name"
                        value={firstName}
                        onChange={(e) => setFirstName(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Last Name"
                        value={lastName}
                        onChange={(e) => setLastName(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="email"
                        className="form-control"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="tel"
                        className="form-control"
                        placeholder="Phone"
                        value={phone}
                        onChange={(e) => setPhone(e.target.value)}
                        required
                        pattern="[0-9]{10}"
                        title="Введите номер телефона, состоящий из 10 цифр."
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="date"
                        className="form-control"
                        placeholder="Date of Birth"
                        value={dateOfBirth}
                        onChange={(e) => setDateOfBirth(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Passport Series"
                        value={passportSerie}
                        onChange={(e) => setPassportSerie(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Passport Number"
                        value={passportNumber}
                        onChange={(e) => setPassportNumber(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="password"
                        className="form-control"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <button type="submit" className="btn btn-primary w-100" style={{fontFamily: "'Poppins', sans-serif"}}>
                    Signup
                </button>
            </form>
        </div>
    );
};

export default Signup;
