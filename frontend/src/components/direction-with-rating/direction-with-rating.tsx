// @ts-ignore
export const DirectionWithRating = ({ info }) => {

    console.log(info)

    return (
        <div className="direction-with-rating-wrapper">
            <p>{info.university_name}</p>
            {info.directions.map((d: any) => (
                <div>
                    <p>{d.name}</p>
                    <p>Место: {d.position}</p>
                    <p>Балл: {d.score}</p>
                    <p>Число бюджетных мест: {d.budget_places}</p>
                    <p>П1 выше вас: {d.priority_one_upper}</p>
                    <p>Количество согласий выше вас: {d.submitted_consent_upper}</p>
                    <br/>
                </div>
            ))}
        </div>
    )
}
