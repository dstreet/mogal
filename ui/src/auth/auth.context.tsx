import { PropsWithChildren, createContext, useCallback, useContext, useState } from "react"

interface AuthContext {
  authenticated: boolean
  token?: string
  expiresIn: number

  login(token: string, expiresIn: number): void
  logout(): void
}

const AuthContext = createContext<AuthContext>({
  authenticated: false,
  expiresIn: 0,
  login: () => {},
  logout: () => {}
})

export const AuthContextProvider: React.FC<PropsWithChildren> = (props) => {
  const [token, setToken] = useState<string | null>(null)
  const [expiresIn, setExpiresIn] = useState<number>(0)

  const login = useCallback((token: string, expiresIn: number) => {
    console.log('do login', token, expiresIn)
    setToken(token)
    setExpiresIn(expiresIn)
  }, [setToken, setExpiresIn])

  const logout = useCallback(() => {
    setToken(null)
  }, [setToken])

  const value = {
    authenticated: !!token,
    token: token ?? undefined,
    expiresIn,
    login,
    logout
  }

  return (
    <AuthContext.Provider value={value}>
      {props.children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  return useContext(AuthContext)
}