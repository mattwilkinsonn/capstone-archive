import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'
import Box from '@material-ui/core/Box'
import Button from '@material-ui/core/Button'
import { Link } from 'react-router-dom'
import { useLogoutMutation, useMeQuery } from '../generated/graphql'
import { createClient } from '../graphql/createClient'
import { useQueryClient } from 'react-query'

const NavBar = (): JSX.Element => {
  const rqClient = createClient()
  const queryClient = useQueryClient()

  const { data: meData } = useMeQuery(rqClient, {}, { staleTime: 300000 })

  const { mutateAsync: logout } = useLogoutMutation(rqClient, {
    onSuccess: () => {
      queryClient.invalidateQueries('Me')
    },
  })

  return (
    <div>
      <AppBar position="static">
        <Toolbar>
          <Box display="flex" flexGrow={1}>
            <Typography variant="h5" color="inherit">
              Capstone Archive
            </Typography>
          </Box>
          {meData?.me ? (
            <Button
              size="small"
              variant="contained"
              color="primary"
              onClick={() => logout({})}
            >
              Sign Out
            </Button>
          ) : (
            <Button
              size="small"
              variant="contained"
              color="primary"
              component={Link}
              to={{
                pathname: '/login',
              }}
            >
              Admin Login
            </Button>
          )}
        </Toolbar>
      </AppBar>
    </div>
  )
}
export default NavBar
