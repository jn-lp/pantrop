import {FC, Fragment} from 'react'

import {useRouter} from "next/router";
import NextLink from "next/link";

import HeaderContainer from '@components/common/Header/HeaderContainer'

import s from './Header.module.css'
import {User} from "@components/icons";

const Header: FC = ({}) => {
  const {pathname} = useRouter()

  return (
    <HeaderContainer>
      <div className={s.Inner}>
        <NextLink href={'/'}>
          <h1 className={s.Title}>Pantrop</h1>
        </NextLink>
        <NextLink href={'/profile'}>
          <a>
            <User/>
          </a>
        </NextLink>
      </div>
    </HeaderContainer>
  )
}

export default Header
