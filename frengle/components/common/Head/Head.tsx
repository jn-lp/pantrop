import { FC } from 'react'
import NextHead from 'next/head'
import { DefaultSeo } from 'next-seo'
import config from '@config/seo.json'
import { useTheme } from 'next-themes'

const Head: FC = () => {
  const { resolvedTheme } = useTheme()

  return (
    <>
      {/* @ts-ignore */}
      <DefaultSeo {...config} />
      <NextHead>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="manifest" href="/site.webmanifest" key="site-manifest" />
        {/* TODO: change theme-color based on theme */}
        <meta
          name="theme-color"
          content={resolvedTheme === 'light' ? 'white' : '#34353e'}
        />
        <link href='https://api.mapbox.com/mapbox-gl-js/v2.1.1/mapbox-gl.css' rel='stylesheet' />
      </NextHead>
    </>
  )
}

export default Head
