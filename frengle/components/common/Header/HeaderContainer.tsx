import { FC, useEffect, useState } from 'react'
import throttle from 'lodash.throttle'
import s from './Header.module.css'
import { Container } from '@components/ui'

const HeaderContainer: FC = ({ children }) => {
  const [hasScrolled, setHasScrolled] = useState(false)

  useEffect(() => {
    const handleScroll = throttle(() => {
      const offset = 0
      const { scrollTop } = document.documentElement
      const scrolled = scrollTop > offset

      if (hasScrolled !== scrolled) {
        setHasScrolled(scrolled)
      }
    }, 200)

    document.addEventListener('scroll', handleScroll)
    return () => {
      document.removeEventListener('scroll', handleScroll)
    }
  }, [hasScrolled])

  return (
    <header className={[s.Header, hasScrolled && s.Header_scrolled].join(' ')}>
      <Container className={s.Container}>{children}</Container>
    </header>
  )
}

export default HeaderContainer
