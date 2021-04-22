import {
  Container,
  Grid,
  Card,
  CardContent,
  Typography,
  TextField,
  Select,
  Button,
  MenuItem,
  makeStyles,
  fade,
  InputLabel,
} from '@material-ui/core'
import React, { useEffect } from 'react'
import { Controller, useForm } from 'react-hook-form'
import { Link, useHistory } from 'react-router-dom'
import {
  useMeQuery,
  useCreateCapstoneMutation,
  NewCapstone,
} from '../../generated/graphql'
import { createClient } from '../../graphql/createClient'
import NavBar from '../NavBar'

const useStyles = makeStyles((theme) => ({
  heroContent: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(8, 0, 6),
  },
  cardGrid: {
    paddingTop: theme.spacing(8),
    paddingBottom: theme.spacing(8),
  },
  card: {
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
  },
  cardContent: {
    flexGrow: 1,
  },
  footer: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(6),
  },
  formButtons: {
    margin: theme.spacing(2, 0, 0, 1),
  },
}))

export const CreateCapstone: React.FC = (): JSX.Element => {
  const classes = useStyles()
  const rqClient = createClient()
  const history = useHistory()

  const { data, isFetching } = useMeQuery(rqClient, {}, { staleTime: Infinity })

  useEffect(() => {
    if (!isFetching && data?.me?.role != 'ADMIN') {
      history.push('/')
    }
  }, [data, isFetching, history])

  const { mutateAsync } = useCreateCapstoneMutation(rqClient, {})
  const { handleSubmit, control } = useForm<NewCapstone>()

  const onSubmit = async (input: NewCapstone): Promise<void> => {
    const res = await mutateAsync({ input })
    history.push('/view/' + res?.createCapstone?.capstone?.slug)
  }

  return (
    <>
      <NavBar />
      <main>
        <Container className={classes.cardGrid} maxWidth="md">
          <Grid container spacing={4}>
            <Grid item xs={12}>
              <Card className={classes.card}>
                <CardContent className={classes.cardContent}>
                  <Typography gutterBottom variant="h5" component="h2">
                    Add a New Project
                  </Typography>
                  <form onSubmit={handleSubmit(onSubmit)}>
                    <Controller
                      name="title"
                      control={control}
                      rules={{ required: true, minLength: 5 }}
                      render={({ field }) => (
                        <TextField
                          {...field}
                          fullWidth
                          required
                          label="Title"
                          style={{ margin: 8 }}
                        />
                      )}
                    />
                    <Controller
                      name="description"
                      control={control}
                      rules={{ required: true, minLength: 5 }}
                      render={({ field }) => (
                        <TextField
                          {...field}
                          fullWidth
                          label="Description"
                          required
                          style={{ margin: 8 }}
                          multiline
                        />
                      )}
                    />
                    <InputLabel style={{ margin: 8 }} id="semester-label">
                      Semester
                    </InputLabel>
                    <Controller
                      name="semester"
                      control={control}
                      rules={{ required: true }}
                      defaultValue="Spring 2021"
                      render={({ field }) => (
                        <Select
                          {...field}
                          required
                          labelId="semester-label"
                          label="Semester"
                          style={{ margin: 8, width: '20%' }}
                          multiline
                        >
                          <MenuItem value={'Fall 2019'}>Fall 2019</MenuItem>
                          <MenuItem value={'Spring 2020'}>Spring 2020</MenuItem>
                          <MenuItem value={'Fall 2020'}>Fall 2020</MenuItem>
                          <MenuItem value={'Spring 2021'}>Spring 2021</MenuItem>
                        </Select>
                      )}
                    />
                    <Controller
                      name="author"
                      control={control}
                      rules={{ required: true, minLength: 3 }}
                      render={({ field }) => (
                        <TextField
                          {...field}
                          label="Author(s)"
                          fullWidth
                          required
                          style={{ margin: 8 }}
                        />
                      )}
                    />
                    <Button
                      type="submit"
                      color="primary"
                      className={classes.formButtons}
                      size="small"
                      variant="contained"
                    >
                      Add
                    </Button>
                    <Button
                      style={{
                        backgroundColor: '#e57373',
                      }}
                      className={classes.formButtons}
                      size="small"
                      variant="contained"
                      component={Link}
                      to={{
                        pathname: '/',
                      }}
                    >
                      Cancel
                    </Button>
                  </form>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Container>
      </main>
    </>
  )
}
