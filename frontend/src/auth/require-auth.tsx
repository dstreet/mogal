import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "./auth.context"

interface Props {
  redirect: string
}

export const RequireAuth: React.FC<Props> = (props) => {
  const auth = useAuth()

  if (!auth.authenticated) {
    return <Navigate to={props.redirect} replace/>
  }

  return (
    <>
      Require auth...
      <Outlet/>
    </>
  )
}