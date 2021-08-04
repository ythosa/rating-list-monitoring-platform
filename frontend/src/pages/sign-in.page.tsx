import {useContext, useState} from "react"
import {AuthContext} from "../context/auth.context"
import {useHttp} from "../hooks/http.hook"
import {Button, Container, CssBaseline, Grid, Link, TextField, Typography} from "@material-ui/core"

import "./sign-in.page.css"

export const SignInPage = () => {
    const auth = useContext(AuthContext)
    const {loading, error, request, clearError} = useHttp()
    const [form, setForm] = useState({
        username: '', password: ''
    })

    const changeHandler = (event: { target: { name: any; value: any } }) => {
        setForm({...form, [event.target.name]: event.target.value})
    }

    const loginHandler = async () => {
        try {
            const data = await request('http://localhost:8001/api/auth/sign-in', 'POST', {...form})
            console.log(data)
            auth.login(data.access_token, data.refresh_token)
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
                    <Grid container>
                        <Grid item>
                            <Link href="#" variant="body2" className="sign-up-link">
                                {"Впервые? Регистрация"}
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
        </Container>
    )
}
