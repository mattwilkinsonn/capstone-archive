import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import Box from '@material-ui/core/Box'
import Button from '@material-ui/core/Button'
import { Link } from 'react-router-dom'

const NavBar = (): JSX.Element => {
  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          <Box display="flex" flexGrow={1}>
            <Typography variant="h5" color="inherit">
              Capstone Archive
            </Typography>
          </Box>
          <Button
            size="small"
            variant="contained"
            color="primary"
            component={Link}
            to={{
              pathname: '/login',
            }}
          >
            Sign Out
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  )
}
export default NavBar
