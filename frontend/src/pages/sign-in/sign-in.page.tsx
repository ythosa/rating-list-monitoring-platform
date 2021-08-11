import { useContext, useState } from 'react'
import { AuthContext } from '../../context/auth.context'
import { Button, Collapse, Container, CssBaseline, Grid, TextField, Typography } from '@material-ui/core'
import Alert from '@material-ui/lab/Alert'
import { Link } from 'react-router-dom'
import AuthorizationService from '../../services/authorization-service'
import UserCredentialsDTO from '../../services/dto/user-credentials.dto'

import './sign-in.page.css'

export const SignInPage = () => {
    const authContext = useContext(AuthContext)
    const authService = new AuthorizationService()
    const [ form, setForm ] = useState<UserCredentialsDTO>({
        username: '', password: '',
    })
    const [ error, setError ] = useState<string>('')

    const changeHandler = (event: { target: { name: any; value: any } }) => {
        setForm({ ...form, [event.target.name]: event.target.value })
    }

    const loginHandler = async () => {
        try {
            const { accessToken, refreshToken } = await authService.signIn(form)
            authContext.login(accessToken, refreshToken)
        } catch (e) {
            setError(e.message)
            console.log(e)
        }
    }

    const errorAlert = error ?
        <Alert className="alert" severity="error" variant="filled" onClose={() => {setError('')}}>
            {error}
        </Alert> : null

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
                    <Collapse in={!!error}>
                        {errorAlert}
                    </Collapse>
                    <Grid container className="sign-up-link-wrapper">
                        <Grid item>
                            <Link to="/sign-up" className="sign-up-link">
                                {'Впервые? Регистрация'}
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
        </Container>
    )
}
