import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Header from './components/Header/Header';
import 'bootstrap/dist/css/bootstrap.min.css'
import SearchResults from "./components/SearchtResult/SearchResult";
import Booking from "./pages/Booking";
import EditProfile from "./pages/EditProfile";
import History from "./pages/History";
import Admin from "./admin/Admin";
import EditFlight from "./admin/EditFlight";
import Add from "./admin/Add";

function App() {
    return (
        <Router>
            <Header />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/search" element={<SearchResults />} />
                <Route path="/book/:flightId" element={<Booking />} />
                <Route path="/edit-profile" element={<EditProfile />} />
                <Route path="/history" element={<History />} />
                <Route path="/admin" element={<Admin />} />
                <Route path="/admin/edit-flight/:id" element={<EditFlight />} />
                <Route path="/admin/add" element={<Add />} />
            </Routes>
        </Router>
    );
}

export default App;
