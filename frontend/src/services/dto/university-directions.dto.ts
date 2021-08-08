import DirectionWithRatingDTO from './direction-with-rating.dto'

export default class UniversityDirectionsDTO {
    constructor(
        readonly id: number,
        readonly name: string,
        readonly fullName: string,
        readonly directions: DirectionWithRatingDTO[],
    ) {}
}
