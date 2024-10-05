import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password }),
            });

            if (response.ok) {
                const data = await response.json();
                localStorage.setItem('token', data.token);
                localStorage.setItem('email', email);
                navigate('/');
            } else {
                const errorData = await response.json();
                alert(errorData.error);
            }
        } catch (error) {
            console.error('Login error:', error);
            alert('An error occurred. Please try again.');
        }
    };

    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4" style={{fontFamily: "'Roboto', sans-serif"}}>
                Login
            </h1>
            <form onSubmit={handleLogin} className="mx-auto" style={{maxWidth: '400px'}}>
            <div className="mb-3">
                    <input
                        type="email"
                        className="form-control"
                        placeholder="Enter your email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <div className="mb-3">
                    <input
                        type="password"
                        className="form-control"
                        placeholder="Enter your password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        style={{fontFamily: "'Poppins', sans-serif"}}
                    />
                </div>
                <button type="submit" className="btn btn-primary w-100" style={{fontFamily: "'Poppins', sans-serif"}}>
                    Login
                </button>
            </form>
        </div>

    );
};

export default Login;
