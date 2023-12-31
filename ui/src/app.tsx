import { CssBaseline, ThemeProvider } from '@mui/material';
import { Route, Routes } from 'react-router-dom';
import { LoginPage } from './auth/login.page';
import { RegisterPage } from './auth/register.page';
import { AuthenticatedPage } from './authenticated.page';
import { MoviesPage } from './movies/movies.page';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthContextProvider } from './auth/auth.context';
import { theme } from './theme';

const queryClient = new QueryClient()

function App() {
  return (
    <>
      <ThemeProvider theme={theme}>
      <CssBaseline/>
      <QueryClientProvider client={queryClient}>
        <AuthContextProvider>
          <Routes>
            <Route path="/login" element={<LoginPage/>}/>
            <Route path="/register" element={<RegisterPage/>}/>
            <Route path="/" element={<AuthenticatedPage redirect="/login"/>}>
              <Route index element={<MoviesPage />}/>
            </Route>
          </Routes>
        </AuthContextProvider>
      </QueryClientProvider>
      </ThemeProvider>
    </>
  );
}

export default App;
