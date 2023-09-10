import { Alert, Box } from "@mui/material"
import { useMovies } from "./movies"
import { MovieView } from "./movie.view"

export const MoviesPage: React.FC = () => {
  const { loading, error, movies } = useMovies()

  console.log(movies)

  return (
    <Box>
      {
        error
          ? <Alert severity="error" sx={{ mb: 2 }}>Could not get movies</Alert>
          : null
      }
      <Box sx={{
        display: 'grid',
        gridTemplateColumns: '1fr 1fr 1fr 1fr',
        gridAutoRows: 'auto',
        columnGap: 2,
        rowGap: 2
      }}>
        {
          movies.map(m => (
            <MovieView key={m.id} movie={m}/>
          ))
        }
      </Box>
    </Box>
  )
}