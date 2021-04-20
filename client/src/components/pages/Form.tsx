import React from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { fade, makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import NavBar from '../NavBar'
import { TextField } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { Link } from 'react-router-dom'
import MenuItem from '@material-ui/core/MenuItem'
import { createClient } from '../../graphql/createClient'
import {
  NewCapstone,
  useCreateCapstoneMutation,
  useMeQuery,
} from '../../generated/graphql'

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

export default function Form(): JSX.Element {
  const classes = useStyles()

  const semesters = [
    {
      value: 'Spring 2021',
      label: 'Spring 2021',
    },
    {
      value: 'Fall 2020',
      label: 'Fall 2020',
    },
  ]

  const [tValid, setTFlag] = React.useState(false)
  const [dValid, setDFlag] = React.useState(false)
  const [sValid, setSFlag] = React.useState(false)
  const [aValid, setAFlag] = React.useState(false)

  const rqClient = createClient()

  const { data, isFetching } = useMeQuery(rqClient, {}, { staleTime: 360000 })

  console.log(data)

  const { mutateAsync } = useCreateCapstoneMutation(rqClient, {})

  // what you want to call with the form when it submits
  const handleSubmit = async (input: NewCapstone): Promise<void> => {
    console.log(input)
    if (t == '') {
      setTFlag(true)
    }
    if (d == '') {
      setDFlag(true)
    }
    if (s == '') {
      setSFlag(true)
    }
    if (a == '') {
      setAFlag(true)
    } else {
      await mutateAsync({ input })
    }
    // probably want to send them back to the homepage or something
    // could also get the id from the return and then send them to the capstone they just created
  }

  const [s, setSemester] = React.useState('')
  const [t, setTitle] = React.useState('')
  const [d, setDescription] = React.useState('')
  const [a, setAuthor] = React.useState('')

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    //fix this
    if (
      event.target.value == 'Fall 2020' ||
      event.target.value == 'Spring 2021'
    ) {
      setSFlag(false)
      setSemester(event.target.value)
    } else if (event.target.id == 'title') {
      setTFlag(false)
      setTitle(event.target.value)
    } else if (event.target.id == 'desc') {
      setDFlag(false)
      setDescription(event.target.value)
    } else if (event.target.id == 'auth1') {
      setAFlag(false)
      setAuthor(event.target.value)
    }
  }

  return (
    <React.Fragment>
      <NavBar></NavBar>
      <main>
        <Container className={classes.cardGrid} maxWidth="md">
          <Grid container spacing={4}>
            <Grid item xs={12}>
              <Card className={classes.card}>
                <CardContent className={classes.cardContent}>
                  <Typography gutterBottom variant="h5" component="h2">
                    Add New Project
                  </Typography>
                  <form autoComplete="off" noValidate>
                    <TextField
                      required
                      fullWidth
                      id="title"
                      value={t}
                      onChange={handleChange}
                      style={{ margin: 8 }}
                      label="Title"
                      error={tValid}
                    />
                    <TextField
                      required
                      fullWidth
                      id="desc"
                      value={d}
                      onChange={handleChange}
                      style={{ margin: 8 }}
                      label="Description"
                      multiline
                      error={dValid}
                    />
                    <TextField
                      required
                      select
                      label="Select"
                      value={s}
                      id="sem"
                      onChange={handleChange}
                      style={{ margin: 8, width: '20%' }}
                      error={sValid}
                    >
                      {semesters.map((option) => (
                        <MenuItem key={option.value} value={option.value}>
                          {option.label}
                        </MenuItem>
                      ))}
                    </TextField>
                    <TextField
                      required
                      style={{ margin: 8 }}
                      id="auth1"
                      label="Author(s)"
                      value={a}
                      onChange={handleChange}
                      error={aValid}
                      fullWidth
                    />
                    <Button
                      type="submit"
                      color="primary"
                      className={classes.formButtons}
                      size="small"
                      variant="contained"
                      onClick={() =>
                        handleSubmit({
                          title: t,
                          description: d,
                          author: a,
                          semester: s,
                        })
                      }
                      // component={Link}
                      // to={{
                      //   pathname: '/',
                      // }}
                    >
                      Add
                    </Button>
                    <Button
                      color="secondary"
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
    </React.Fragment>
  )
}
