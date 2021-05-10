import React, { FC } from 'react'
import s from './Container.module.css'

interface Props {
  className?: string
  children?: any
  el?: HTMLElement
  clean?: boolean
}

const Container: FC<Props> = ({ children, className, el = 'div', clean }) => {
  let Component: React.ComponentType<
    React.HTMLAttributes<HTMLDivElement>
  > = el as any

  return (
    <Component className={[className, !clean && s.Container].join(' ')}>
      {children}
    </Component>
  )
}

export default Container
