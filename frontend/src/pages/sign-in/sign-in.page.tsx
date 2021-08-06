import { useContext, useState } from 'react'
import { AuthContext } from '../../context/auth.context'
import { Button, Container, CssBaseline, Grid, Link, TextField, Typography } from '@material-ui/core'

import './sign-in.page.css'
import Authorization from '../../services/authorization'
import UserCredentials from '../../services/dto/user-credentials'

export const SignInPage = () => {
    const authContext = useContext(AuthContext)
    const authService = new Authorization()
    const [ form, setForm ] = useState<UserCredentials>({
        username: '', password: '',
    })

    const changeHandler = (event: { target: { name: any; value: any } }) => {
        setForm({ ...form, [event.target.name]: event.target.value })
    }

    const loginHandler = async () => {
        try {
            const { accessToken, refreshToken } = await authService.signIn(form)
            authContext.login(accessToken, refreshToken)
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <Container component="main" maxWidth="xs" className="sign-in-container">
            <CssBaseline/>
            <div>
                <Typography component="h1" variant="h5" className="sign-in-title">
                    Добро пожаловать
                </Typography>
                <form noValidate>
                    <TextField
                        variant="filled"
                        margin="normal"
                        required
                        fullWidth
                        id="username"
                        label="Username"
                        name="username"
                        autoFocus
                        className="sign-in-input"
                        value={form.username}
                        onChange={changeHandler}
                    />
                    <TextField
                        variant="filled"
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="Password"
                        type="password"
                        id="password"
                        autoComplete="current-password"
                        className="sign-in-input"
                        value={form.password}
                        onChange={changeHandler}
                    />
                    <Button
                        fullWidth
                        variant="contained"
                        color="primary"
                        className="sign-in-button"
                        onClick={loginHandler}
                    >
                        Войти
                    </Button>
                    <Grid container className="sign-up-link-wrapper">
                        <Grid item>
                            <Link href="/sign-up" variant="subtitle1" className="sign-up-link">
                                {'Впервые? Регистрация'}
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
        </Container>
    )
}
