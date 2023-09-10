import { createTheme } from '@mui/material/styles';

export const theme = createTheme({
  palette: {
    mode: 'light',
    primary: {
      main: '#3f51b5',
    },
    secondary: {
      main: '#f50057',
    },
  },
  typography: {
    h1: {
      fontSize: '3.7rem',
    },
    h2: {
      fontSize: '3rem',
    },
    h3: {
      fontSize: '2.5rem',
    },
  },
})