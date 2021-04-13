import React from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { fade, makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import NavBar from '../NavBar'
import Divider from '@material-ui/core/Divider'
import Button from '@material-ui/core/Button'
import { Link } from 'react-router-dom'
import { createClient } from '../../graphql/createClient'
import { useCapstoneQuery } from '../../generated/graphql'

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
  divide: {
    marginBottom: theme.spacing(1),
  },
  info: {
    margin: theme.spacing(1, 0, 1, 0),
  },
  backButton: {
    margin: theme.spacing(0, 0, 0, 2),
  },
}))

export default function Moreviewpage(props: {
  match: {
    params: {
      id: string
    }
  }
}): JSX.Element {
  const classes = useStyles()

  const rqClient = createClient()
  const id = parseInt(props.match.params.id)
  if (isNaN(id) || id < 0) {
    throw new Error('page id invalid')
  }
  const { data } = useCapstoneQuery(rqClient, { id }, { staleTime: 300000 })

  console.log(data?.capstone)

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
                    {data?.capstone?.title}
                  </Typography>
                  <Divider
                    variant="middle"
                    className={classes.divide}
                  ></Divider>
                  <Typography className={classes.info}>
                    <strong>Semester: </strong>
                    {data?.capstone?.semester}
                  </Typography>
                  <Typography className={classes.info}>
                    <strong>Authors: </strong>
                    {data?.capstone?.author}
                  </Typography>
                  <Typography className={classes.info}>
                    <strong>Project Description: </strong>
                    {data?.capstone?.description}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
            <Button
              className={classes.backButton}
              size="small"
              variant="contained"
              component={Link}
              to={{
                pathname: '/',
              }}
            >
              Back
            </Button>
          </Grid>
        </Container>
      </main>
    </React.Fragment>
  )
}
