export default class DirectionWithRatingDTO {
    constructor(
        readonly id: number,
        readonly name: string,
        readonly position: number,
        readonly budgetPlaces: number,
        readonly score: number,
        readonly submittedConsentUpper: number,
        readonly priorityOneUpper: number,
    ) {}
}
