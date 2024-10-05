import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Header from './components/Header/Header';
import 'bootstrap/dist/css/bootstrap.min.css'
import SearchResults from "./components/FlightResult/flightResult";

function App() {
    return (
        <Router>
            <Header />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/search" element={<SearchResults />} />
            </Routes>
        </Router>
    );
}

export default App;
