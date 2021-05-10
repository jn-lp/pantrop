import '@styles/chrome-bug.css'
import '@styles/colors.css'
import '@styles/themes.css'
import '@styles/reset.css'
import '@styles/globals.css'
import 'mapbox-gl/dist/mapbox-gl.css';

import {FC, useEffect} from 'react'
import type {AppProps} from 'next/app'
import {Head} from '@components/common'
import {ManagedUIContext} from '@components/ui/context'

const Noop: FC = ({children}) => <>{children}</>

export default function MyApp({Component, pageProps}: AppProps) {
  const Layout = (Component as any).Layout || Noop

  useEffect(() => {
    document.body.classList?.remove('loading')
  }, [])

  return (
    <>
      <ManagedUIContext>
        <Head/>
        <Layout pageProps={pageProps}>
          <Component {...pageProps} />
        </Layout>
      </ManagedUIContext>
    </>
  )
}
