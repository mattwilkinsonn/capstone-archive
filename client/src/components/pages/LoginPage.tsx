import React from 'react'
import Avatar from '@material-ui/core/Avatar'
import Button from '@material-ui/core/Button'
import CssBaseline from '@material-ui/core/CssBaseline'
import TextField from '@material-ui/core/TextField'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined'
import Typography from '@material-ui/core/Typography'
import { makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import { Link } from 'react-router-dom'
import { Login, useLoginMutation } from '../../generated/graphql'
import { createClient } from '../../graphql/createClient'

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: '100%',
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}))

export default function LoginPage(): JSX.Element {
  const classes = useStyles()

  const rqClient = createClient()
  const { mutateAsync, data } = useLoginMutation(rqClient, {})

  const [user, setUsername] = React.useState('')
  const [pass, setPassword] = React.useState('')
  const [usernameValid, setUserFlag] = React.useState(false)
  const [passwordValid, setPassFlag] = React.useState(false)
  const [userErrorText, setUserText] = React.useState('')
  const [passErrorText, setPassText] = React.useState('')

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    if (event.target.id == 'username') {
      setUserFlag(false)
      setUsername(event.target.value)
      setUserText('')
    } else if (event.target.id == 'password') {
      setPassFlag(false)
      setPassword(event.target.value)
      setPassText('')
    }
  }

  const handleSubmit = async (input: Login): Promise<void> => {
    console.log(input)
    if (user == '') {
      setUserFlag(true)
      setUserText('Must enter username or email')
    }
    if (pass == '') {
      setPassFlag(true)
      setPassText('Must enter password')
    // } else if (user !== "admin@test.com") {
    //   setUserFlag(true)
    //   setUserText('Invalid Username')
    // } else if (user == "admin@test.com" && pass !== "hunter2") {
    //   setPassFlag(true)
    //   setPassText('Incorrect Password')
    } else {
      await mutateAsync({ input })
      // want to redirect to homepage or somewhere here
    }
  }

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <form className={classes.form}>
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            error={usernameValid}
            helperText={userErrorText}
            onChange={handleChange}
          />
          <TextField
            variant="outlined"
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            error={passwordValid}
            helperText={passErrorText}
            onChange={handleChange}
          />
          <Button
            type="button"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
            onClick={() =>
              handleSubmit({
                usernameOrEmail: user,
                password: pass,
              })
            }
            component={Link}
            to={{
              pathname: '/',
            }}
          >
            Sign In
          </Button>
        </form>
      </div>
    </Container>
  )
}
