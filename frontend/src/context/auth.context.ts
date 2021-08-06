import { createContext } from 'react'

function login(accessToken: string, refreshToken: string) {
    return
}

function logout() {
    return
}

export class AuthContextType {
    constructor(
        public accessToken: string | null,
        public refreshToken: string | null,
        public login: (at: string, rt: string) => void,
        public logout: () => void,
        public isAuthenticated: boolean,
    ) {}
}

export const AuthContext = createContext<AuthContextType>({
    accessToken: '',
    refreshToken: '',
    login,
    logout,
    isAuthenticated: false,
})
