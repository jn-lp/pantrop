import {cloneElement, FC, JSXElementConstructor, ReactElement} from 'react'

import {useRouter} from 'next/router'
import Link, {LinkProps} from 'next/link'

import s from './Footer.module.css'

interface Props extends LinkProps {
  children: ReactElement<any, string | JSXElementConstructor<any>>
}

const NavLink: FC<Props> = ({href, children}) => {
  const {pathname} = useRouter();

  return <Link href={href}>
    {cloneElement(children, {
      className: [s.NavLink, pathname == href && s.NavLink_active].join(' ')
    })}
  </Link>;
};

export default NavLink;
