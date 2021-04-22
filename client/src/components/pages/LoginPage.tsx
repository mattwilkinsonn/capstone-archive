import {
  Avatar,
  Button,
  Container,
  CssBaseline,
  makeStyles,
  TextField,
  Typography,
} from '@material-ui/core'
import React, { useEffect } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { useHistory } from 'react-router'
import { Login, useLoginMutation, useMeQuery } from '../../generated/graphql'
import { createClient } from '../../graphql/createClient'
import LockOutlinedIcon from '@material-ui/icons/LockOutlined'
import { useQueryClient } from 'react-query'

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

export const LoginPage: React.FC = () => {
  const classes = useStyles()
  const rqClient = createClient()
  const history = useHistory()
  const queryClient = useQueryClient()

  const { handleSubmit, control, setError } = useForm<Login>()

  const { data, isFetching } = useMeQuery(rqClient, {}, { staleTime: Infinity })

  const { mutateAsync: login } = useLoginMutation(rqClient, {
    onSuccess: () => {
      queryClient.invalidateQueries('Me')
    },
  })

  useEffect(() => {
    if (!isFetching && data?.me) {
      history.push('/')
    }
  }, [data, isFetching, history])

  const onSubmit = async (input: Login): Promise<void> => {
    const res = await login({ input })
    if (res.login.user) {
      history.push('/')
    } else if (
      res.login.error?.field === 'usernameOrEmail' ||
      res.login.error?.field === 'password'
    ) {
      setError(res.login.error.field, { message: res.login.error.message })
    } else {
      console.error('incorrect return from server')
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
        <form className={classes.form} onSubmit={handleSubmit(onSubmit)}>
          <Controller
            name="usernameOrEmail"
            control={control}
            rules={{ required: true }}
            render={({ field }) => (
              <TextField
                {...field}
                variant="outlined"
                fullWidth
                required
                margin="normal"
                label="Username/Email"
                autoComplete="username"
                style={{ margin: 8 }}
              />
            )}
          />
          <Controller
            name="password"
            control={control}
            rules={{ required: true }}
            render={({ field }) => (
              <TextField
                {...field}
                variant="outlined"
                fullWidth
                required
                margin="normal"
                label="Password"
                type="password"
                autoComplete="current-password"
                style={{ margin: 8 }}
              />
            )}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            color="primary"
            className={classes.submit}
          >
            Sign In
          </Button>
        </form>
      </div>
    </Container>
  )
}
