import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
    return (
        <div>
            <h1>Welcome to the Home Page</h1>
            <p>
                Please <Link to="/login">Login</Link> or <Link to="/signup">Signup</Link>.
            </p>
        </div>
    );
};

export default Home;
