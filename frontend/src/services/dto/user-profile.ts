export default class UserProfileDTO {
    constructor(
        readonly username: string,
        readonly firstName: string,
        readonly middleName: string,
        readonly lastName: string,
        readonly snils: string,
    ) {}
}
