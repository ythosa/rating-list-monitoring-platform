import { Redirect, Route, Switch } from 'react-router-dom'
import SignInPage from '../pages/sign-in'
import SignUpPage from '../pages/sign-up'

export const useRoutes = (isAuthenticated: boolean) => {
    if (isAuthenticated) {
        return (
            <Switch>
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
