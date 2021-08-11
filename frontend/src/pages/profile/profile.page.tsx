import React, { useContext, useEffect, useState } from 'react'
import Header from '../../components/header'
import Loader from '../../components/loader'
import { AuthContext } from '../../context/auth.context'
import UserProfileDTO from '../../services/dto/user-profile'
import UserService from '../../services/user-service'
import './profile.page.css'

export const ProfilePage = () => {
    const authContext = useContext(AuthContext)
    const [ profileData, setProfileData ] = useState<UserProfileDTO | null>(null)
    const [ loading, setLoading ] = useState<boolean>(true)

    useEffect(() => {
        const userService = new UserService(authContext.accessToken!!)
        userService.getProfile().then(data => {
            setProfileData(data)
            setLoading(false)
        }).catch(e => {
            console.log(e)
            authContext.logout()
        })
    }, [ authContext ])

    const loadingBanner = loading ? <Loader/> : null
    const content = !loading ? (
        <React.Fragment>
            <div className="profile-page-title">
                {profileData?.username} profile
            </div>
            <div className="profile-page-credentials">
                <div className="profile-page-credential">
                    <b>Фамилия:</b> {profileData?.lastName}
                </div>
                <div className="profile-page-credential">
                    <b>Имя:</b> {profileData?.firstName}
                </div>
                <div className="profile-page-credential">
                    <b>Отчество:</b> {profileData?.middleName}
                </div>
                <div className="profile-page-credential">
                    <b>СНИЛС:</b> {formatSnils(profileData!!.snils)}
                </div>
            </div>
        </React.Fragment>
    ) : null

    console.log(profileData)

    return (
        <React.Fragment>
            <Header/>
            <div className="profile-page-wrapper">
                {loadingBanner}
                {content}
            </div>
        </React.Fragment>
    )
}

function formatSnils(snils: string): string {
    return `${snils.slice(0, 3)}-${snils.slice(3, 6)}-${snils.slice(6, 9)} ${snils.slice(9, 11)}`
}
