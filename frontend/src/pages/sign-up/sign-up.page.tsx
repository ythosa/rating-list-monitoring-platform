import { useContext, useState } from 'react'
import { AuthContext } from '../../context/auth.context'
import { useHttp } from '../../hooks/http.hook'
import { useHistory } from 'react-router-dom'

import { Button, Container, CssBaseline, Grid, Link, TextField, Typography } from '@material-ui/core'
import './sign-up.page.css'

export const SignUpPage = () => {
    const auth = useContext(AuthContext)
    const { loading, error, request, clearError } = useHttp()
    const [ form, setForm ] = useState({
        username: '', password: '', firstName: '', middleName: '', lastName: '', snils: ''
    })
    const history = useHistory()

    const changeHandler = (event: { target: { name: any; value: any } }) => {
        setForm({ ...form, [event.target.name]: event.target.value })
    }

    const registerHandler = async () => {
        try {
            const data = await request('http://localhost:8001/api/auth/sign-up', 'POST', {
                'username': form.username,
                'password': form.password,
                'first_name': form.firstName,
                'middle_name': form.middleName,
                'last_name': form.lastName,
                'snils': form.snils
            })
            console.log(data)
            history.push('/sign-in')
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <Container component="main" maxWidth="xs" className="sign-up-container">
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
                        id="firstName"
                        label="First name"
                        name="firstName"
                        className="sign-in-input"
                        value={form.firstName}
                        onChange={changeHandler}
                    />
                    <TextField
                        variant="filled"
                        margin="normal"
                        required
                        fullWidth
                        id="middleName"
                        label="Middle name"
                        name="middleName"
                        className="sign-in-input"
                        value={form.middleName}
                        onChange={changeHandler}
                    />
                    <TextField
                        variant="filled"
                        margin="normal"
                        required
                        fullWidth
                        id="lastName"
                        label="Last name"
                        name="lastName"
                        className="sign-in-input"
                        value={form.lastName}
                        onChange={changeHandler}
                    />
                    <TextField
                        variant="filled"
                        margin="normal"
                        required
                        fullWidth
                        id="snils"
                        label="Snils"
                        name="snils"
                        className="sign-in-input"
                        value={form.snils}
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
                        onClick={registerHandler}
                    >
                        Зарегистрироваться
                    </Button>
                    <Grid container>
                        <Grid item>
                            <Link href="/sign-in" variant="subtitle1" className="sign-up-link">
                                {'Уже есть аккаунт? Войти'}
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
        </Container>
    )
}