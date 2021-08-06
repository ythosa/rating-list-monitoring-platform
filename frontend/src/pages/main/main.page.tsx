import Header from '../../components/header'
import { useAsync } from 'react-async'
import { useContext } from 'react'
import { AuthContext } from '../../context/auth.context'
import DirectionWithRating from '../../components/direction-with-rating'

import './main.page.css'

// @ts-ignore
const getDirectionsWithRating = async ({accessToken}, { signal }) => {
    const res = await fetch('http://localhost:8001/api/direction/get_for_user_with_rating', {
        headers: { accept: 'application/json', Authorization: `rlmp ${accessToken}` },
        method: 'GET',
    })
    if (!res.ok) throw new Error(res.statusText)

    return res.json()
}

export const MainPage = () => {
    const auth = useContext(AuthContext)
    // @ts-ignore
    const { data, error, isPending } = useAsync({ promiseFn: getDirectionsWithRating, accessToken: auth.accessToken})

    console.log(data)
    console.log(error)

    if (isPending) return (
        <div className="main-page-wrapper">
            <Header/>
            Loading...
        </div>
    )

    if (error) auth.logout()

    const directionsWithRating = (data as [any]).map((d : any) => <DirectionWithRating info={d}/>)

    return (
        <div className="main-page-wrapper">
            <Header/>
            <div className="directions-with-rating-container">
                {directionsWithRating}
            </div>
        </div>
    )
}
