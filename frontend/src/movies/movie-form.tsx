import { Autocomplete, Box, Button, FormControl, FormLabel, MenuItem, Rating, Select, TextField } from "@mui/material"
import ImageIcon from '@mui/icons-material/Image';
import { Genre } from "../genres/genres.interface"
import { useState } from "react";
import { useGenres } from "../genres/genres";

interface Props {
  value?: MovieFormData
  onSubmit: (value: MovieFormData) => void
  submitLabel?: string
}

export interface MovieFormData {
  title: string
  rating: string
  cast: string[]
  director: string
  poster?: string
  userRating?: number
  genres: string[]
}

export const MovieForm: React.FC<Props> = (props) => {
  const { value, onSubmit, submitLabel = "Create" } = props
  const { genres: availableGenres } = useGenres()

  const parsedGenres = value?.genres
    .map(id => availableGenres.find(ag => ag.id === id))
    .filter(Boolean) as Genre[]
  
  const [title, setTitle] = useState(value?.title ?? '')
  const [rating, setRating] = useState(value?.rating ?? '')
  const [cast, setCast] = useState(value?.cast ?? [])
  const [director, setDirector] = useState(value?.director ?? '')
  const [poster, setPoster] = useState(value?.poster ?? '')
  const [userRating, setUserRating] = useState(value?.userRating ?? null)
  const [genres, setGenres] = useState<Genre[]>(parsedGenres ?? [])

  const onCastChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCast(e.target.value.split('\n'))
  }

  const onGenreChange = (values: Genre[]) => {
    setGenres(values)
  }

  const onFormSubmit = (e: React.ChangeEvent<HTMLFormElement>) => {
    e.preventDefault()

    onSubmit({
      title,
      rating,
      cast,
      director,
      poster,
      userRating: userRating ?? undefined,
      genres: genres.map(g => g.id)
    })
  }

  return (
    <Box sx={{ display: 'flex'}}>
      <Box sx={{
        width: 300,
        height: 450,
        mr: 2,
        borderWidth: 1,
        borderStyle: 'solid',
        borderColor: 'primary.main',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        borderRadius: 2,
        overflow: 'hidden'
      }}>
        {
          poster
          ? <img src={poster} style={{ width: '100%' }}/>
          : <ImageIcon color="primary"/>
        }
      </Box>
      <Box component="form" sx={{
        '.MuiFormControl-root': {
          display: 'block',
          mb: 2,
        }
      }} onSubmit={onFormSubmit}>
        <FormControl>
          <FormLabel htmlFor="title">Title</FormLabel>
          <TextField id="title" fullWidth size="small" value={title} required onChange={e => setTitle(e.target.value)}/>
        </FormControl>
        <FormControl>
          <FormLabel htmlFor="rating">Movie Rating</FormLabel>
          <Select id="rating" fullWidth size="small" value={rating} required onChange={e => setRating(e.target.value)}>
            <MenuItem value="G">G</MenuItem>
            <MenuItem value="PG">PG</MenuItem>
            <MenuItem value="PG-13">PG-13</MenuItem>
            <MenuItem value="R">R</MenuItem>
            <MenuItem value="NR">NR</MenuItem>
          </Select>
        </FormControl>
        <FormControl>
          <FormLabel htmlFor="cast">Cast</FormLabel>
          <TextField id="cast" multiline fullWidth size="small" value={cast.join('\n')} required onChange={onCastChange}/>
        </FormControl>
        <FormControl>
          <FormLabel htmlFor="director">Director</FormLabel>
          <TextField id="director" fullWidth size="small" value={director} required onChange={e => setDirector(e.target.value)}/>
        </FormControl>
        <FormControl>
          <FormLabel htmlFor="poster">Poster Image</FormLabel>
          <TextField id="poster" fullWidth size="small" value={poster} onChange={e => setPoster(e.target.value)}/>
        </FormControl>
        <FormControl>
          <FormLabel htmlFor="genres">Genres</FormLabel>
          <Autocomplete multiple id="genres" options={availableGenres} getOptionLabel={opt => opt.name} renderInput={params => (
            <TextField
              {...params}
              size="small"
            />
          )} onChange={(_, values) => onGenreChange(values)}/>
        </FormControl>
        <FormControl>
          <FormLabel htmlFor="userRating">Your Rating</FormLabel>
          <Box><Rating id="userRating" value={userRating} onChange={(_, value) => setUserRating(value)}/></Box>
        </FormControl>
        <FormControl>
          <Button variant="contained" type="submit">{submitLabel}</Button>
        </FormControl>
      </Box>
    </Box>
  )
}