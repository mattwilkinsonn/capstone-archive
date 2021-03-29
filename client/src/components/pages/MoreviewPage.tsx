import React from 'react'
import Card from '@material-ui/core/Card'
import CardContent from '@material-ui/core/CardContent'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { fade, makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import NavBar from '../NavBar'
import Divider from '@material-ui/core/Divider'
import CardMedia from '@material-ui/core/CardMedia'
import Button from '@material-ui/core/Button'
import { Link } from 'react-router-dom'
import CardActions from '@material-ui/core/CardActions'

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
  cardMedia: {
    paddingTop: '56.25%', // 16:9
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
}))

export default function Moreviewpage(props: {
  location: {
    state: {
      name: string
      description: string
      image: string
      semester: string
    }
  }
}): JSX.Element {
  const classes = useStyles()

  const result = props.location.state.name
  const desc = props.location.state.description
  const img = props.location.state.image
  const sem = props.location.state.semester

  return (
    <React.Fragment>
      <NavBar></NavBar>
      <main>
        <Container className={classes.cardGrid} maxWidth="md">
          <Grid container spacing={4}>
            <Grid item xs={12}>
              <Card className={classes.card}>
                <CardMedia
                  className={classes.cardMedia}
                  image={img}
                  title="Image title"
                />
                <CardContent className={classes.cardContent}>
                  <Typography gutterBottom variant="h5" component="h2">
                    {result}
                  </Typography>
                  <Divider
                    variant="middle"
                    className={classes.divide}
                  ></Divider>
                  <Typography className={classes.info}>
                    <strong>Semester:</strong> {sem}
                  </Typography>
                  <Typography className={classes.info}>
                    <strong>Authors:</strong>{' '}
                  </Typography>
                  <Typography className={classes.info}>
                    <strong>Project Description:</strong> {desc}
                  </Typography>
                  <Typography>
                    <strong>Project URL:</strong>{' '}
                  </Typography>
                </CardContent>
                <CardActions>
                  <Button
                    size="small"
                    color="primary"
                    component={Link}
                    to={{
                      pathname: '/',
                    }}
                  >
                    Back
                  </Button>
                </CardActions>
              </Card>
            </Grid>
          </Grid>
        </Container>
      </main>
      {/* Footer */}
      <footer className={classes.footer}>
        <Typography variant="h6" align="center" gutterBottom>
          American University CS Capstone Project
        </Typography>
        <Typography
          variant="subtitle1"
          align="center"
          color="textSecondary"
          component="p"
        >
          By Ashlyn Levinson and Matthew Wilkinson
        </Typography>
      </footer>
      {/* End footer */}
    </React.Fragment>
  )
}
