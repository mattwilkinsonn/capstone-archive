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
import { NewCapstone, useCreateCapstoneMutation } from '../../generated/graphql'

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
      value: 'S21',
      label: 'Spring 2021',
    },
    {
      value: 'F20',
      label: 'Fall 2020',
    },
  ]
  const rqClient = createClient()

  const { mutateAsync, data } = useCreateCapstoneMutation(rqClient, {})

  // what you want to call with the form when it submits
  const handleSubmit = async (input: NewCapstone): Promise<void> => {
    await mutateAsync({ input })
    // probably want to send them back to the homepage or something
    // could also get the id from the return and then send them to the capstone they just created
  }

  let semester

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    semester = event.target.value
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
                  <form noValidate autoComplete="off">
                    <TextField
                      required
                      fullWidth
                      style={{ margin: 8 }}
                      label="Title"
                    />
                    <TextField
                      required
                      fullWidth
                      style={{ margin: 8 }}
                      label="Description"
                      multiline
                    />
                    <TextField
                      required
                      fullWidth
                      select
                      label="Select"
                      value={semester}
                      onChange={handleChange}
                      style={{ margin: 8 }}
                      helperText="Please select your semester"
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
                      label="Author 1"
                    />
                    <TextField style={{ margin: 8 }} label="Author 2" />
                    <TextField style={{ margin: 8 }} label="Author 3" />
                    <TextField style={{ margin: 8 }} label="Author 4" />
                    <TextField
                      fullWidth
                      style={{ margin: 8 }}
                      label="URL (optional)"
                    />
                    <Button
                      color="primary"
                      className={classes.formButtons}
                      size="small"
                      variant="contained"
                      component={Link}
                      to={{
                        pathname: '/',
                      }}
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
