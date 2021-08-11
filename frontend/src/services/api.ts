export default class API {
    private baseURL: string = 'http://localhost:8001/api'

    protected getResource = async (
        url: string, method: string = 'GET', payload?: any, headers?: { [key: string]: string },
    ) => {
        const res = await fetch(
            `${this.baseURL}${url}`,
            {
                method,
                body: payload ? JSON.stringify(payload) : undefined,
                headers: { accept: 'application/json', ...headers },
            },
        )
        
        if (!res.ok) {
            const error = await res.json()
            throw Error(error.message)
        }

        return await res.json()
    }

    protected getAuthorizationHeader(accessToken: string): { Authorization: string } {
        return {
            Authorization: `Bearer ${accessToken}`
        }
    }
}
