import {cloneElement, MouseEvent, useEffect, useRef} from 'react'
import hasParent from './hasParent'

interface ClickOutsideProps {
  active: boolean
  onClick: (e?: MouseEvent) => void
  children: any
}

const ClickOutside = ({
                        active = true,
                        onClick,
                        children,
                      }: ClickOutsideProps) => {
  const innerRef = useRef()

  const handleClick = (event: any) => {
    if (!hasParent(event.target, innerRef?.current)) {
      if (typeof onClick === 'function') {
        onClick(event)
      }
    }
  }

  useEffect(() => {
    if (active) {
      document.addEventListener('mousedown', handleClick)
      document.addEventListener('touchstart', handleClick)
    }

    return () => {
      if (active) {
        document.removeEventListener('mousedown', handleClick)
        document.removeEventListener('touchstart', handleClick)
      }
    }
  })

  return cloneElement(children, {ref: innerRef})
}

export default ClickOutside
