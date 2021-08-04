import React, { useContext } from 'react'
import { AuthContext } from '../../context/auth.context'
import { useFetch } from 'react-async'

import './header.css'
import { Avatar, createStyles, Link, makeStyles, Theme } from '@material-ui/core'
import profileIcon from './profile.png'

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        large: {
            width: theme.spacing(6),
            height: theme.spacing(6)
        }
    })
)

export const Header = () => {
    const classes = useStyles()
    const preventDefault = (event: React.SyntheticEvent) => event.preventDefault()
    const auth = useContext(AuthContext)

    const { data, error } = useFetch('http://localhost:8001/api/user/get_username', {
        headers: { accept: 'application/json', Authorization: `rlmp ${auth.accessToken}` },
        method: 'GET'
    })

    if (error) auth.logout()

    const username = (data as { username: string })?.username

    return (
        <div className="header-wrapper">
            <div className="header-username-profile">
                <Avatar alt={username} src={profileIcon} className={classes.large}/>
                <span className="header-username">{username}</span>
            </div>
            <Link href="#" onClick={preventDefault} className="header-nav-link">ВУЗЫ</Link>
            <Link href="#" onClick={preventDefault} className="header-nav-link">Программы</Link>
            <Link href="#" onClick={auth.logout} className="header-nav-link">Выйти</Link>
        </div>
    )
}
