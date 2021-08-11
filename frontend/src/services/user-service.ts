import API from './api'

import UserProfileDTO from './dto/user-profile'

export default class UserService extends API {
    private readonly accessToken: string

    constructor(accessToken: string) {
        super()
        this.accessToken = accessToken
    }

    getProfile = async (): Promise<UserProfileDTO> => {
        const data = await this.getResource('/user/get_profile', 'GET', null, {
            ...this.getAuthorizationHeader(this.accessToken),
        })

        return new UserProfileDTO(data.username, data.first_name, data.middle_name, data.last_name, data.snils)
    }

    getUsername = async (): Promise<string> => {
        const data = await this.getResource('/user/get_username', 'GET', null, {
            ...this.getAuthorizationHeader(this.accessToken),
        })

        return data.username
    }
}
