import API from './api'
import UserCredentials from './dto/user-credentials'
import Tokens from './dto/tokens'
import CreatingAccountData from './dto/creating-accout-data'
import ID from './dto/id'

export default class Authorization extends API {
    signIn = async (payload: UserCredentials): Promise<Tokens> => {
        const body = await this.getResource('/auth/sign-in', 'POST', { ...payload })

        return new Tokens(body.access_token, body.refresh_token)
    }

    signUp = async (payload: CreatingAccountData): Promise<ID> => {
        const body = await this.getResource('/auth/sign-up', 'POST', {
            username: payload.username,
            password: payload.password,
            first_name: payload.firstName,
            middle_name: payload.middleName,
            last_name: payload.lastName,
            snils: payload.snils,
        })

        return new ID(body.id)
    }
}
