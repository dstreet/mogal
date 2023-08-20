import { Box, Button, Input } from "@mui/material"
import { useRegisterUser } from "./register"
import { useCallback } from "react"

export const RegisterPage: React.FC = () => {
  const { register, loading, error, token, expiresIn } = useRegisterUser()

  const onRegister = useCallback(() => {
    register({ email: "foo@bar.com", password: "blah" })
  }, [register])

  console.log(token, expiresIn)

  return (
    <Box>
      <Input type="email" placeholder="Email" />
      <Input type="password" placeholder="Password" />
      <Button onClick={onRegister}>Register</Button>
    </Box>
  )
}