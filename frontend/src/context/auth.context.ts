import {createContext} from 'react'

function login(jwtToken: string, id: number) {
    return
}

function logout() {
    return
}

export const AuthContext = createContext({
    token: null,
    userId: null,
    login,
    logout,
    isAuthenticated: false
})
