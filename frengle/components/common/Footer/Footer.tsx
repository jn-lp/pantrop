import {Feed, User} from "@components/icons";

import NavLink from "./NavLink";

import s from './Footer.module.css'
import Explore from "@components/icons/Explore";

export default function Footer() {
  return (
    <footer className={s.Footer}>
      {/*<Container>*/}
      <div className={s.Inner}>
        <NavLink href={'/feed'}>
          <a>
            <Feed/>
            Feed
          </a>
        </NavLink>
        <NavLink href={'/explore'}>
          <a>
            <Explore/>
            Explore
          </a>
        </NavLink>
        <NavLink href={'/profile'}>
          <a>
            <User/>
            Profile
          </a>
        </NavLink>
      </div>
      {/*</Container>*/}
    </footer>
  )
}
