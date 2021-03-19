import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import AccountCircleIcon from '@material-ui/icons/AccountCircle'
import Box from '@material-ui/core/Box'
import Button from '@material-ui/core/Button'

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
          <Button>
            <AccountCircleIcon></AccountCircleIcon>
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  )
}
export default NavBar
