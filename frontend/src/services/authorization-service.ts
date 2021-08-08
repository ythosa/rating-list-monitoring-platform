import API from './api'
import UserCredentialsDTO from './dto/user-credentials.dto'
import TokensDTO from './dto/tokens.dto'
import CreatingAccountDataDTO from './dto/creating-accout-data.dto'
import IdDTO from './dto/id.dto'

export default class AuthorizationService extends API {
    signIn = async (payload: UserCredentialsDTO): Promise<TokensDTO> => {
        const body = await this.getResource('/auth/sign-in', 'POST', { ...payload })

        return new TokensDTO(body.access_token, body.refresh_token)
    }

    signUp = async (payload: CreatingAccountDataDTO): Promise<IdDTO> => {
        const body = await this.getResource('/auth/sign-up', 'POST', {
            username: payload.username,
            password: payload.password,
            first_name: payload.firstName,
            middle_name: payload.middleName,
            last_name: payload.lastName,
            snils: payload.snils,
        })

        return new IdDTO(body.id)
    }
}
