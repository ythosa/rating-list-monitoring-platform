import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { AuthContext } from '../context/auth.context';
import { useAuth } from '../hooks/auth.hook';
import './app.css';
import { useRoutes } from './routes';

const App = () => {
    const { login, logout, token, userId, ready } = useAuth()
    const isAuthenticated = !!token
    const routes = useRoutes(isAuthenticated)

    return (
        <AuthContext.Provider value={{
            login, logout, token, userId, isAuthenticated
        }}>
            <Router>
                <div className="container">
                    { isAuthenticated ? <div>site</div> : null }
                    { routes }
                </div>
            </Router>
        </AuthContext.Provider>
    )
};

export default App;
