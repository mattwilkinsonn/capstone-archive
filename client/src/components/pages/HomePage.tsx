import React from 'react'
import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'
import CardContent from '@material-ui/core/CardContent'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { fade, makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import NavBar from '../NavBar'
import { useCapstonesQuery } from '../../generated/graphql'
import { createClient } from '../../graphql/createClient'
import { Link } from 'react-router-dom'

const useStyles = makeStyles((theme) => ({
  heroContent: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(8, 0, 6),
  },
  noProjContent: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(12),
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
  cardDescription: {
    margin: theme.spacing(1, 0, 0, 0),
  },
  heroButtons: {
    marginTop: theme.spacing(2),
  },
}))

export default function Homepage(): JSX.Element {
  const classes = useStyles()

  // creates the client needed to query data.
  const rqClient = createClient()

  // run the Capstones query, get the data back from that.
  const { data } = useCapstonesQuery(
    rqClient,
    { limit: 20 },
    { staleTime: 300000 }
  )

  const truncateStr = (str: string, num: number): string => {
    if (str.length <= num) {
      return str
    } else {
      return str.slice(0, num) + '...'
    }
  }

  // Log the array of capstones to the console
  console.log(data?.capstones.capstones)

  const renderHeader = (): any => {
    return (
      <div className={classes.heroContent}>
        <Container maxWidth="sm">
          <Typography
            component="h1"
            variant="h2"
            align="center"
            color="textPrimary"
            gutterBottom
          >
            Capstone Projects
          </Typography>
          <Typography
            variant="h5"
            align="center"
            color="textSecondary"
            paragraph
          >
            Past projects can be seen here.
          </Typography>
          <div className={classes.heroButtons}>
            <Grid container spacing={2} justify="center">
              <Grid item>
                <Button
                  size="small"
                  variant="contained"
                  component={Link}
                  to={{
                    pathname: '/search',
                  }}
                >
                  Search for a Capstone
                </Button>
              </Grid>
              <Grid item>
                <Button
                  size="small"
                  variant="contained"
                  component={Link}
                  to={{
                    pathname: '/add',
                  }}
                >
                  Add a Capstone
                </Button>
              </Grid>
            </Grid>
          </div>
        </Container>
      </div>
    )
  }

  if (data?.capstones.capstones) {
    const cards = data?.capstones.capstones
    return (
      <React.Fragment>
        <NavBar></NavBar>
        <main>
          {renderHeader()}
          <Container className={classes.cardGrid} maxWidth="md">
            <Grid container spacing={4}>
              {cards.map((card) => (
                <Grid item key={cards.indexOf(card)} xs={12} sm={6} md={4}>
                  <Card className={classes.card}>
                    <CardContent className={classes.cardContent}>
                      <Typography gutterBottom variant="h5" component="h2">
                        {truncateStr(card!.title, 50)}
                      </Typography>
                      <Typography color="textSecondary">
                        {card?.semester}
                      </Typography>
                      <Typography className={classes.cardDescription}>
                        {truncateStr(card!.description, 100)}
                      </Typography>
                    </CardContent>
                    <CardActions>
                      <Button
                        size="small"
                        variant="contained"
                        component={Link}
                        to={{
                          pathname: 'view/' + card?.title,
                        }}
                      >
                        View
                      </Button>
                    </CardActions>
                  </Card>
                </Grid>
              ))}
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
  } else {
    return (
      <React.Fragment>
        <NavBar></NavBar>
        <main>
          {renderHeader()}
          <div className={classes.noProjContent}>
            <Container maxWidth="sm">
              <Typography
                component="h3"
                variant="h4"
                align="center"
                color="textPrimary"
                gutterBottom
              >
                Cannot load projects
              </Typography>
            </Container>
          </div>
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
}
