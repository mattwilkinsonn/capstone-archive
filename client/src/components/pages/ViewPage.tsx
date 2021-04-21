import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Container from '@material-ui/core/Container'
import Divider from '@material-ui/core/Divider'
import Grid from '@material-ui/core/Grid'
import { fade, makeStyles } from '@material-ui/core/styles'
import Typography from '@material-ui/core/Typography'
import React from 'react'
import { Link } from 'react-router-dom'
import { useCapstoneBySlugQuery } from '../../generated/graphql'
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
      slug: string
    }
  }
}): JSX.Element {
  const classes = useStyles()

  const rqClient = createClient()
  const slug = props.match.params.slug

  const { data } = useCapstoneBySlugQuery(
    rqClient,
    { slug },
    { staleTime: Infinity }
  )

  // console.log(data?.capstoneBySlug)

  const capstone = data?.capstoneBySlug

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
                    {capstone?.title}
                  </Typography>
                  <Divider
                    variant="middle"
                    className={classes.divide}
                  ></Divider>
                  <Typography className={classes.info}>
                    <strong>Semester: </strong>
                    {capstone?.semester}
                  </Typography>
                  <Typography className={classes.info}>
                    <strong>Author(s): </strong>
                    {capstone?.author}
                  </Typography>
                  <Typography className={classes.info}>
                    <strong>Project Description: </strong>
                    {capstone?.description}
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
