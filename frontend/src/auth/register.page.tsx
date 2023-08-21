import { Alert, Box, Button, Card, Input, Typography } from "@mui/material"
import { useRegisterUser } from "./register"
import { FormEvent, useCallback, useEffect, useRef } from "react"
import { useAuth } from "./auth.context"
import { Link, Navigate } from "react-router-dom"

export const RegisterPage: React.FC = () => {
  const { register, loading, error, token, expiresIn } = useRegisterUser()
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

    register(email, pass)
  }, [register])

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
        {error ? <Alert severity="error">Unable to register</Alert> : null}
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
            {loading ? '...' : 'Register'}
          </Button>
        </Box>
        <Link to="/login">Already have an account? Login</Link>
      </Card>
    </Box>
  )
}