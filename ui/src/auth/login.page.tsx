import { FormEvent, useCallback, useEffect, useRef } from "react"
import { useLogin } from "./login"
import { Alert, Box, Button, Card, Input, Typography } from "@mui/material"
import { useAuth } from "./auth.context"
import { Link, Navigate } from "react-router-dom"

export const LoginPage: React.FC = () => {
  const { login, loading, error, token, expiresIn } = useLogin()
  const auth = useAuth()

  const emailRef = useRef<HTMLInputElement>(null)
  const passwordRef = useRef<HTMLInputElement>(null)

  const onSubmit = useCallback((e: FormEvent) => {
    e.preventDefault()

    const email = emailRef.current?.value
    const pass = passwordRef.current?.value

    if (!email || !pass) {
      return
    }

    login(email, pass)
  }, [login])

  useEffect(() => {
    if (!auth.authenticated && token) {
      auth.login(token, Number(expiresIn))
    }
  }, [auth.authenticated, auth.login, token, expiresIn])

  if (auth.authenticated) {
    return <Navigate to="/" replace/>
  }

  return (
    <Box sx={{
      height: '100vh',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center'
    }}>
      <Card sx={{ p: 4 }}>
        {error ? <Alert severity="error">Unable to login</Alert> : null}
        <Typography variant="h1">Mogal</Typography>
        <Box component="form" onSubmit={onSubmit} sx={{
          mt: 2,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'flex-start'
        }}>
          <Input type="email" placeholder="Email" inputRef={emailRef} required/>
          <Input type="password" placeholder="Password" inputRef={passwordRef} required/>
          <Button type="submit" variant="contained" sx={{ mt: 2, mb: 2 }} disabled={loading}>
            {loading ? '...' : 'Login'}
          </Button>
        </Box>
        <Link to="/register">Or register</Link>
      </Card>
    </Box>
  )
}