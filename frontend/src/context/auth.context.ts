import { createContext } from 'react'

function login(jwtToken: string, id: number) {
    return
}

function logout() {
    return
}

export const AuthContext = createContext({
    accessToken: null,
    refreshToken: null,
    login,
    logout,
    isAuthenticated: false
})
