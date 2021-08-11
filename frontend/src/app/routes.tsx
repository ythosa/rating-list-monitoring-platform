import { Redirect, Route, Switch } from 'react-router-dom'
import SignInPage from '../pages/sign-in'
import SignUpPage from '../pages/sign-up'
import MainPage from '../pages/main'
import ProfilePage from '../pages/profile'

export const useRoutes = (isAuthenticated: boolean) => {
    if (isAuthenticated) {
        return (
            <Switch>
                <Route path="/results" exact>
                    <MainPage/>
                </Route>
                <Route path="/profile" exact>
                    <ProfilePage />
                </Route>
                <Redirect to="/results"/>
            </Switch>
        )
    }

    return (
        <Switch>
            <Route path="/sign-in" exact>
                <SignInPage/>
            </Route>
            <Route path="/sign-up" exact>
                <SignUpPage/>
            </Route>
            <Redirect to="/sign-in"/>
        </Switch>
    )
}
