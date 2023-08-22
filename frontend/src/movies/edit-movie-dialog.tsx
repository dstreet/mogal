import { Dialog, DialogContent, DialogTitle } from "@mui/material"
import { MovieForm, MovieFormData } from "./movie-form"
import { Movie } from "./movie.interface"

interface Props {
  open: boolean
  movie: Movie
  onCancel: () => void
}

export const EditMovieDialog: React.FC<Props> = (props) => {
  const { open, onCancel, movie } = props

  const onSubmit = (value: MovieFormData, newGenres: string[]) => {
    console.log(value, newGenres)
  }

  const movieValue = {
    ...movie,
    genres: movie.genres.map(g => g.id)
  }

  return (
    <Dialog onClose={() => onCancel()} open={open}>
      <DialogTitle>Update movie</DialogTitle>
      <DialogContent>
        <MovieForm onSubmit={onSubmit} submitLabel="Update" value={movieValue}/>
      </DialogContent>
    </Dialog>
  )
}