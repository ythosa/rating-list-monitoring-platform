import {Redirect, Route, Switch} from 'react-router-dom'
import {SignInPage} from '../pages/sign-in.page'

export const useRoutes = (isAuthenticated: boolean) => {
    if (isAuthenticated) {
        return (
            <Switch>
            </Switch>
        )
    }

    return (
        <Switch>
            <Route path="/" exact>
                <SignInPage/>
            </Route>
            <Redirect to="/"/>
        </Switch>
    )
}
