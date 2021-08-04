import { useCallback, useEffect, useState } from 'react'

const storageName = 'userData'

export const useAuth = () => {
    const [ accessToken, setAccessToken ] = useState(null)
    const [ refreshToken, setRefreshToken ] = useState(null)
    const [ ready, setReady ] = useState(false)

    const login = useCallback((accessJWT, refreshJWT) => {
        setAccessToken(accessJWT)
        setRefreshToken(refreshJWT)

        localStorage.setItem(storageName, JSON.stringify({
            accessToken: accessJWT,
            refreshToken: refreshJWT
        }))
    }, [])

    const logout = useCallback(() => {
        setAccessToken(null)
        setRefreshToken(null)

        localStorage.removeItem(storageName)
    }, [])

    useEffect(() => {
        const storedData = localStorage.getItem(storageName)
        if (!storedData) {
            return
        }

        const data = JSON.parse(storedData)
        if (data && data.accessToken && data.refreshToken) {
            login(data.accessToken, data.refreshToken)
        }
        setReady(true)
    }, [ login ])

    return { login, logout, accessToken, refreshToken, ready }
}
