import Header from '../../components/header'
import { useContext, useEffect, useState } from 'react'
import { AuthContext } from '../../context/auth.context'
import DirectionWithRating from '../../components/direction-with-rating'

import './main.page.css'
import Direction from '../../services/direction'
import UniversityDirections from '../../services/dto/university-directions'

export const MainPage = () => {
    const authContext = useContext(AuthContext)
    const [ universities, setUniversities ] = useState<UniversityDirections[]>()
    const [ loading, setLoading ] = useState<boolean>(true)

    useEffect(() => {
        const directionsService = new Direction(authContext.accessToken!!)
        directionsService.getForUserWithRating().then((data) => {
            setUniversities(data)
            setLoading(false)
        })
    }, [ authContext.accessToken ])

    const loadingBanner = loading ? <span>Loading...</span> : null
    const content = !loading ? universities?.map(u => <DirectionWithRating info={u}/>) : null

    return (
        <div className="main-page-wrapper">
            <Header/>
            <div className="directions-with-rating-container">
                {loadingBanner}
                {content}
            </div>
        </div>
    )
}
