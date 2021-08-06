import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { AuthContext } from '../context/auth.context'
import { useAuth } from '../hooks/auth.hook'
import './app.css'
import { useRoutes } from './routes'

const App = () => {
    const { login, logout, accessToken, refreshToken } = useAuth()
    const isAuthenticated = !!accessToken
    const routes = useRoutes(isAuthenticated)

    return (
        <AuthContext.Provider value={{
            login, logout, accessToken, refreshToken, isAuthenticated,
        }}>
            <Router>
                <div className="container">{routes}</div>
            </Router>
        </AuthContext.Provider>
    )
}

export default App
