module.exports = {
  i18n: {
    locales: ['en-US', 'ru-RU', 'uk-UA'],
    defaultLocale: 'ru-RU',
  },
  future: {
    webpack5: true
  },
  reactStrictMode: true,
  typescript: {
    ignoreBuildErrors: true,
  },
  productionBrowserSourceMaps: true,
}

console.log('next.config.js', JSON.stringify(module.exports, null, 2))
