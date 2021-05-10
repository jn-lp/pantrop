import React, {FC} from 'react'

import dynamic from 'next/dynamic'
import {useRouter} from 'next/router'
import NextLink from "next/link";

import {useAcceptCookies} from '@utils/hooks/useAcceptCookies'
import {Header} from '@components/common'
import {Plus} from "@components/icons";

const Loading = () => <>{/*<LoadingDots />*/}</>

const dynamicProps = {
  loading: () => <Loading/>,
}

const FeatureBar = dynamic(
  () => import('@components/common/FeatureBar'),
  dynamicProps
)

interface Props {
  pageProps: {
    sc00lFeatures: Record<string, boolean>
  }
}

const Layout: FC<Props> = ({
                             children,
                             pageProps: {sc00lFeatures, ...pageProps},
                           }) => {
  const {acceptedCookies, onAcceptCookies} = useAcceptCookies()

  const {locale = 'ru-RU', asPath} = useRouter()

  return (
    <>
      <Header/>
      <main className="fit">
        {children}
      </main>
      {/*<Footer/>*/}
      <NextLink href={"/create"}>
        <button className={"create"}>
          <Plus width={"32"} height={"32"}/>
        </button>
      </NextLink>

      <FeatureBar
        title={'Наши сайты используют куки'}
        action={
          <button
            // size={'sm'}
            // kind={'secondary'}
            onClick={() => onAcceptCookies()}
          >
            Ладно
          </button>
        }
        open={!acceptedCookies}
      />
    </>
  )
}

export default Layout
