import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "./auth/auth.context"
import { AppBar, Box, Button, IconButton, Toolbar, Typography } from "@mui/material"
import { LogoutOutlined } from "@mui/icons-material"
import { MovieForm, MovieFormData } from "./movies/movie-form"
import { CreateMovieDialog } from "./movies/create-movie-dialog"
import { useState } from "react"
import { Movie } from "./movies/movie.interface"

interface Props {
  redirect: string
}

export const AuthenticatedPage: React.FC<Props> = (props) => {
  const auth = useAuth()
  const [showCreateDialog, setShowCreateDialog] = useState(false)

  if (!auth.authenticated) {
    return <Navigate to={props.redirect} replace/>
  }

  const onCreateNew = () => {
    setShowCreateDialog((show) => !show)
  }

  const onCloseCreateNew = () => {
    setShowCreateDialog(false)
  }

  const onCreateSuccess = (movie: Movie) => {
    setShowCreateDialog(false)
    console.log(movie)
  }

  return (
    <Box>
      <AppBar position="static">
        <Toolbar sx={{ display: 'flex', justifyContent: 'space-between'}}>
          <Typography variant="h6">Mogal</Typography>
          <Box>
            <Button variant="contained" color="secondary" onClick={onCreateNew}>Create Movie</Button>
            <IconButton>
              <LogoutOutlined sx={{ color: 'white' }}/>
            </IconButton>
          </Box>
        </Toolbar>
      </AppBar>
      <Box component="main" sx={{ p: 2 }}>
        <CreateMovieDialog open={showCreateDialog} onCancel={onCloseCreateNew} onCreate={onCreateSuccess}/>
        <Outlet/>
      </Box>
    </Box>
  )
}