import React from 'react';
import { Link } from 'react-router-dom';

const HomePage = () => {
    return (
        <div>
            <h1>Welcome to the Polling App</h1>
            <Link to="/polls/create">
                <button>Create a Poll</button>
            </Link>
        </div>
    );
};

export default HomePage;