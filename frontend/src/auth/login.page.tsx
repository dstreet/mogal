import { FormEvent, useCallback, useEffect, useRef } from "react"
import { useLogin } from "./login"
import { Box, Button, Input } from "@mui/material"
import { useAuth } from "./auth.context"
import { Navigate } from "react-router-dom"

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
    <Box>
      <form onSubmit={onSubmit}>
        <Input type="email" placeholder="Email" inputRef={emailRef} required/>
        <Input type="password" placeholder="Password" inputRef={passwordRef} required/>
        <Button type="submit">
          {loading ? '...' : 'Login'}
        </Button>
      </form>
    </Box>
  )
}