import API from './api'
import UniversityDirectionsDTO from './dto/university-directions.dto'
import DirectionWithRatingDTO from './dto/direction-with-rating.dto'

export default class DirectionService extends API {
    private readonly accessToken: string

    constructor(accessToken: string) {
        super()
        this.accessToken = accessToken
    }

    getForUserWithRating = async () => {
        const data = await this.getResource(
            '/direction/get_for_user_with_rating', 'GET', null, {
                ...this.getAuthorizationHeader(this.accessToken),
            },
        )

        return data.map((u: any) => new UniversityDirectionsDTO(
            u.university_id, u.university_name, u.university_full_name,
            u.directions.map((d: any) => new DirectionWithRatingDTO(
                d.id, d.name, d.position, d.budget_places, d.score, d.submitted_consent_upper, d.priority_one_upper,
            )),
        ))
    }
}
