export default class CreatingAccountDataDTO {
    constructor(
        readonly username: string,
        readonly password: string,
        readonly firstName: string,
        readonly middleName: string,
        readonly lastName: string,
        readonly snils: string,
    ) {}
}
