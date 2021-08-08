import Header from '../../components/header'
import { useContext, useEffect, useState } from 'react'
import { AuthContext } from '../../context/auth.context'
import DirectionWithRating from '../../components/direction-with-rating'
import DirectionService from '../../services/direction-service'
import UniversityDirectionsDTO from '../../services/dto/university-directions.dto'

import './main.page.css'
import Loader from '../../components/loader'

export const MainPage = () => {
    const authContext = useContext(AuthContext)
    const [ universities, setUniversities ] = useState<UniversityDirectionsDTO[]>()
    const [ loading, setLoading ] = useState<boolean>(true)

    useEffect(() => {
        const directionsService = new DirectionService(authContext.accessToken!!)
        directionsService.getForUserWithRating().then((data) => {
            setUniversities(data)
            setLoading(false)
        }).catch((e) => {
            console.log(e)
            authContext.logout()
        })
    }, [ authContext ])

    const loadingBanner = loading ? <Loader/> : null
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
