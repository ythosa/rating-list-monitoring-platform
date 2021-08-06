import UniversityDirections from '../../services/dto/university-directions'
import DirectionWithRatingDTO from '../../services/dto/direction-with-rating'

import './direction-with-rating.css'

export const DirectionWithRating = ({ info }: { info: UniversityDirections }) => {
    return (
        <div className="direction-with-rating-wrapper">
            <p className="direction-with-rating-wrapper-title">{info.fullName}</p>
            <p className="direction-with-rating-wrapper-text">
                Ваш балл с учетом ИД: <b>{info.directions[0].score}</b>
            </p>
            <p className="direction-with-rating-wrapper-text">Конкурсные группы:</p>
            {info.directions.map((d: DirectionWithRatingDTO) => (
                <ul className="direction-rating-ul direction-rating-text">
                    <li>{d.name}</li>
                    <p>Позиция в списке: <b>{d.position}</b></p>
                    <p>Число бюджетных мест: <b>{d.budgetPlaces}</b></p>
                    <p>Число П1 выше вас: <b>{d.priorityOneUpper}</b></p>
                    <p>Количество поданных согласий выше вас: <b>{d.submittedConsentUpper}</b></p>
                </ul>
            ))}
        </div>
    )
}
