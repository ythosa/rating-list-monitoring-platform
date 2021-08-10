import API from './api'

export default class UserService extends API {
    private readonly accessToken: string

    constructor(accessToken: string) {
        super()
        this.accessToken = accessToken
    }

    getUsername = async (): Promise<string> => {
        const data = await this.getResource('/user/get_username', 'GET', null, {
            ...this.getAuthorizationHeader(this.accessToken),
        })

        return data.username
    }
}
