import { Card, CardActionArea, CardContent, CardMedia, Typography } from "@mui/material"
import { Movie } from "./movie.interface"
import { EditMovieDialog } from "./edit-movie-dialog"
import { useState } from "react"

interface Props {
  movie: Movie
}

export const MovieView: React.FC<Props> = (props) => {
  const { movie } = props
  const [openEdit, setOpenEdit] = useState(false)

  return (
    <>
      <Card>
        <CardActionArea onClick={() => setOpenEdit(true)}>
          <CardMedia
            image={movie.poster}
            title={movie.title}
            sx={{height: 200, backgroundPosition: 'top center'}}
          />
          <CardContent>
            <Typography variant="h6">{movie.title}</Typography>
          </CardContent>
        </CardActionArea>
      </Card>
      <EditMovieDialog open={openEdit} onCancel={() => setOpenEdit(false)} movie={movie}/>
    </>
  )
}