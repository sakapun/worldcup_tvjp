module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: 'worldcup_tvjp',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'worldcup 2018' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** Run ESLint on save
    */
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    },
    postcss: {
      plugins: {
        'postcss-custom-properties': false
      }
    }
  },
  css: [
    // node.js module but we specify the pre-processor
    { src: '~assets/main.sass', lang: 'sass' }
  ],
  modules: [
    "@nuxtjs/proxy",
    "@nuxtjs/axios"
  ],
  proxy: {
    "/api" : "http://localhost:8080"
  }
}
