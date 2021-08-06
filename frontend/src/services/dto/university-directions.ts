import DirectionWithRatingDTO from './direction-with-rating'

export default class UniversityDirections {
    constructor(
        readonly id: number,
        readonly name: string,
        readonly fullName: string,
        readonly directions: DirectionWithRatingDTO[],
    ) {}
}
