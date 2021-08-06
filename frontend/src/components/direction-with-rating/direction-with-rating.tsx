import "./direction-with-rating.css"

// @ts-ignore
export const DirectionWithRating = ({ info }) => {

    console.log(info)

    return (
        <div className="direction-with-rating-wrapper">
            <p className="direction-with-rating-wrapper-title">{info.university_full_name}</p>
            <p className="direction-with-rating-wrapper-text">Ваш балл с учетом ИД: <b>{info.directions[0].score}</b></p>
            <p className="direction-with-rating-wrapper-text">Конкурсные группы:</p>
            {info.directions.map((d: any) => (
                <ul className="direction-rating-ul direction-rating-text">
                    <li>{d.name}</li>
                    <p>Позиция в списке: <b>{d.position}</b></p>
                    <p>Число бюджетных мест: <b>{d.budget_places}</b></p>
                    <p>Число П1 выше вас: <b>{d.priority_one_upper}</b></p>
                    <p>Количество поданных согласий выше вас: <b>{d.submitted_consent_upper}</b></p>
                </ul>
            ))}
        </div>
    )
}
