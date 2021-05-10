import React, {FC, useMemo} from 'react'
import {ThemeProvider} from 'next-themes'

export interface State {
}

const initialState: State = {}

type Action =
  | {
  type: 'EXAMPLE'
}

export const UIContext = React.createContext<State | any>(initialState)

UIContext.displayName = 'UIContext'

function uiReducer(state: State, action: Action) {
  switch (action.type) {
    case 'EXAMPLE': {
      return {
        ...state,
      }
    }
  }
}

export const UIProvider: FC = (props) => {
  const [state] = React.useReducer(uiReducer, initialState)

  const value = useMemo(
    () => ({
      ...state,
    }),
    [state]
  )

  return <UIContext.Provider value={value} {...props} />
}

export const useUI = () => {
  const context = React.useContext(UIContext)
  if (context === undefined) {
    throw new Error(`useUI must be used within a UIProvider`)
  }
  return context
}

export const ManagedUIContext: FC = ({children}) => (
  <UIProvider>
    <ThemeProvider forcedTheme={"dark"}>{children}</ThemeProvider>
  </UIProvider>
)
