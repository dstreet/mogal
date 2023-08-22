import { Alert, Dialog, DialogContent, DialogTitle } from "@mui/material"
import { MovieForm, MovieFormData } from "./movie-form"
import { useCreateMovie } from "./create-movie"
import { Movie } from "./movie.interface"
import { useEffect } from "react"

interface Props {
  open: boolean
  onCancel: () => void
  onCreate: (movie: Movie) => void
}

export const CreateMovieDialog: React.FC<Props> = (props) => {
  const { open, onCancel, onCreate } = props
  const { loading, success, error, movie, createMovie } = useCreateMovie()

  const onSubmit = (value: MovieFormData, newGenres: string[]) => {
    console.log(newGenres, value.genres)
    createMovie(value, newGenres)
  }

  useEffect(() => {
    if (success) {
      onCreate(movie as Movie)
    }
  }, [success])

  return (
    <Dialog onClose={() => onCancel()} open={open}>
      <DialogTitle>Create movie</DialogTitle>
      <DialogContent>
        {
          error
            ? <Alert severity="error" sx={{ mb: 2 }}>Failed to create movie</Alert>
            : null
        }
        <MovieForm onSubmit={onSubmit}/>
      </DialogContent>
    </Dialog>
  )
}