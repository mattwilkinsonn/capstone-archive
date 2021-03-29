import React from 'react'
import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'
import CardContent from '@material-ui/core/CardContent'
import CardMedia from '@material-ui/core/CardMedia'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { fade, makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import NavBar from '../NavBar'
import SearchIcon from '@material-ui/icons/Search'
import InputBase from '@material-ui/core/InputBase'
import { useCapstonesQuery } from '../../generated/graphql'
import { createClient } from '../../graphql/createClient'
import { Link } from 'react-router-dom'

const useStyles = makeStyles((theme) => ({
  icon: {
    marginRight: theme.spacing(2),
  },
  heroContent: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(8, 0, 6),
  },
  noProjContent: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(12),
  },
  heroButtons: {
    marginTop: theme.spacing(4),
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
  search: {
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: fade(theme.palette.primary.main, 0.15),
    '&:hover': {
      backgroundColor: fade(theme.palette.primary.main, 0.2),
    },
    marginTop: theme.spacing(4),
    marginRight: 'auto',
    marginLeft: 'auto',
    width: '50%',
  },
  searchIcon: {
    padding: theme.spacing(0, 2),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  inputRoot: {
    color: 'inherit',
  },
  inputInput: {
    padding: theme.spacing(1, 1, 1, 0),
    paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
  },
  cardDescription: {
    margin: theme.spacing(0, 0, 1, 0),
  },
  cardSemester: {
    color: fade(theme.palette.primary.light, 0.5),
  },
}))

export default function Homepage(): JSX.Element {
  const classes = useStyles()

  const title = 'Project Title'
  const desc = 'Project Description'
  const img = 'https://source.unsplash.com/random'

  // creates the client needed to query data.
  const rqClient = createClient()

  // run the Capstones query, get the data back from that.
  const { data } = useCapstonesQuery(
    rqClient,
    { limit: 20 },
    { staleTime: 300000 }
  )

  // Log the array of capstones to the console
  console.log(data?.capstones.capstones)

  if (data?.capstones.capstones) {
    const cards = data?.capstones.capstones
    return (
      <React.Fragment>
        <NavBar></NavBar>
        <main>
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
              <div className={classes.search}>
                <div className={classes.searchIcon}>
                  <SearchIcon />
                </div>
                <InputBase
                  placeholder="Search…"
                  classes={{
                    root: classes.inputRoot,
                    input: classes.inputInput,
                  }}
                  inputProps={{ 'aria-label': 'search' }}
                />
              </div>
            </Container>
          </div>
          <Container className={classes.cardGrid} maxWidth="md">
            <Grid container spacing={4}>
              {cards.map((card) => (
                <Grid item key={cards.indexOf(card)} xs={12} sm={6} md={4}>
                  <Card className={classes.card}>
                    <CardMedia
                      className={classes.cardMedia}
                      image={img}
                      title="Image title"
                    />
                    <CardContent className={classes.cardContent}>
                      <Typography gutterBottom variant="h5" component="h2">
                        {card?.title}
                      </Typography>
                      <Typography className={classes.cardSemester}>
                        {card?.semester}
                      </Typography>
                      <Typography className={classes.cardDescription}>
                        {card?.description}
                      </Typography>
                    </CardContent>
                    <CardActions>
                      <Button
                        size="small"
                        color="primary"
                        component={Link}
                        to={{
                          pathname: '/View',
                          state: {
                            name: card?.title,
                            description: card?.description,
                            image: img,
                          },
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
          <div className={classes.noProjContent}>
            <Container maxWidth="sm">
              <Typography
                component="h1"
                variant="h2"
                align="center"
                color="textPrimary"
                gutterBottom
              >
                No Projects to Show
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
