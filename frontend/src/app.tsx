import { CssBaseline } from '@mui/material';
import { Route, Routes } from 'react-router-dom';
import { LoginPage } from './auth/login.page';
import { RegisterPage } from './auth/register.page';
import { RequireAuth } from './auth/require-auth';
import { MoviesPage } from './movies/movies.page';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthContextProvider } from './auth/auth.context';

const queryClient = new QueryClient()

function App() {
  return (
    <>
      <CssBaseline/>
      <QueryClientProvider client={queryClient}>
        <AuthContextProvider>
          <Routes>
            <Route path="/login" element={<LoginPage/>}/>
            <Route path="/register" element={<RegisterPage/>}/>
            <Route path="/" element={<RequireAuth redirect="/login"/>}>
              <Route index element={<MoviesPage />}/>
            </Route>
          </Routes>
        </AuthContextProvider>
      </QueryClientProvider>
    </>
  );
}

export default App;
