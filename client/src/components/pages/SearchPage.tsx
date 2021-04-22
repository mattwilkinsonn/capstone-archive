import React, { useState } from 'react'
import Grid from '@material-ui/core/Grid'
import Typography from '@material-ui/core/Typography'
import { fade, makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container'
import NavBar from '../NavBar'
import SearchIcon from '@material-ui/icons/Search'
import { Link } from 'react-router-dom'
import Button from '@material-ui/core/Button'
import Card from '@material-ui/core/Card'
import CardActions from '@material-ui/core/CardActions'
import CardContent from '@material-ui/core/CardContent'
import { useSearchCapstonesQuery } from '../../generated/graphql'
import { createClient } from '../../graphql/createClient'
import TextField from '@material-ui/core/TextField'
import InputAdornment from '@material-ui/core/InputAdornment'

const useStyles = makeStyles((theme) => ({
  heroContent: {
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    padding: theme.spacing(8, 0, 6),
  },
  noProjContent: {
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
    backgroundColor: fade(theme.palette.primary.light, 0.1),
    '&:hover': {
      backgroundColor: fade(theme.palette.primary.light, 0.2),
    },
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
    marginTop: theme.spacing(4),
  },
  cardDescription: {
    margin: theme.spacing(1, 0, 0, 0),
  },
  heroButtons: {
    marginTop: theme.spacing(2),
  },
}))

export default function SearchPage(): JSX.Element {
  const classes = useStyles()
  const rqClient = createClient()
  const [searchTerm, setSearchTerm] = useState('')
  const [queryEnabled, setQueryEnabled] = useState(false)

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    setSearchTerm(event.target.value)
    if (!queryEnabled) setQueryEnabled(true)
  }

  // run the Capstones query, get the data back from that.
  const { data, isLoading } = useSearchCapstonesQuery(
    rqClient,
    { limit: 500, query: searchTerm },
    { enabled: queryEnabled }
  )

  const truncateStr = (str: string, num: number): string => {
    if (str.length <= num) {
      return str
    } else {
      return str.slice(0, num) + '...'
    }
  }

  // Log the array of capstones to the console
  console.log(data?.searchCapstones.capstones)

  const renderHeader = (): JSX.Element => {
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
            Search Capstones
          </Typography>
          <Typography
            variant="h5"
            align="center"
            color="textSecondary"
            paragraph
          >
            Search for Capstone Projects here.
          </Typography>
          <form noValidate autoComplete="off" className={classes.search}>
            <Grid container spacing={1} justify="center">
              <Grid item>
                <TextField
                  id="input"
                  placeholder="Search"
                  value={searchTerm}
                  onChange={handleChange}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <SearchIcon />
                      </InputAdornment>
                    ),
                  }}
                />
              </Grid>
            </Grid>
          </form>
          <div className={classes.heroButtons}>
            <Grid container spacing={2} justify="center">
              <Grid item>
                <Button
                  size="small"
                  variant="contained"
                  className={classes.heroButtons}
                  component={Link}
                  to={{
                    pathname: '/',
                  }}
                >
                  Back to Main
                </Button>
              </Grid>
            </Grid>
          </div>
        </Container>
      </div>
    )
  }

  if (queryEnabled && data?.searchCapstones.capstones) {
    const cards = data?.searchCapstones.capstones
    if (cards.length !== 0) {
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
                            pathname: 'view/' + card?.slug,
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
                  No Projects Found
                </Typography>
              </Container>
            </div>
          </main>
        </React.Fragment>
      )
    }
  } else {
    return (
      <React.Fragment>
        <NavBar></NavBar>
        <main>{renderHeader()}</main>
      </React.Fragment>
    )
  }
}
