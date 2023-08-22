import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "./auth/auth.context"
import { AppBar, Box, Button, IconButton, Toolbar, Typography } from "@mui/material"
import { LogoutOutlined } from "@mui/icons-material"
import { MovieForm, MovieFormData } from "./movies/movie-form"

interface Props {
  redirect: string
}

export const AuthenticatedPage: React.FC<Props> = (props) => {
  const auth = useAuth()

  if (!auth.authenticated) {
    return <Navigate to={props.redirect} replace/>
  }

  const onFormSubmit = (values: MovieFormData) => {
    console.log(values)
  }

  return (
    <Box>
      <AppBar position="static">
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between'}}>
          <Typography variant="h6">Mogal</Typography>
          <Box>
            <Button variant="contained" color="secondary">Create Movie</Button>
            <IconButton>
              <LogoutOutlined sx={{ color: 'white' }}/>
            </IconButton>
          </Box>
        </Toolbar>
      </AppBar>
      <Box component="main" sx={{ p: 2 }}>
        <MovieForm onSubmit={onFormSubmit}/>
        <Outlet/>
      </Box>
    </Box>
  )
}