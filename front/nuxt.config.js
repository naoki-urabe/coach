import colors from "vuetify/es5/util/colors";

export default {
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: "%s - front",
    title: "front",
    htmlAttrs: {
      lang: "en"
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" },
      { name: "format-detection", content: "telephone=no" }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: ['~/plugins/axios.js',{src: '~/plugins/persistedState.client.js'}],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/vuetify
    "@nuxtjs/vuetify"
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: ["@nuxtjs/axios", "@nuxtjs/auth-next"],
  axios: {
    baseURL: `http://${process.env.NUXT_ENV_APP_HOST}:${process.env.NUXT_ENV_API_PORT}/api/`
  },
  auth: {
    redirect: {
      login: "/login",
      logout: "/login",
      callback: false,
      home: "/"
    },
    strategies: {
      local: {
        endpoints: {
          login: {
            url: '/auth/login',
            method: "post",
            propertyName: "token"
          },
          logout: false,
          user: false
          // user: {url: '/api/auth/user', method: 'get', propertyName: 'user'}
        }
      }
    }
  },
  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ["~/assets/variables.scss"],
    theme: {
      dark: false,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        }
      }
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},
  router: {
    middleware: ["auth"]
  },
  publicRuntimeConfig: {
    appHost: process.env.NUXT_ENV_APP_HOST,
    apiPort: process.env.NUXT_ENV_API_PORT,
    frontPort: process.env.NUXT_ENV_FRONT_PORT,
  },
  privateRuntimeConfig: {},
  server: {
    host: '0.0.0.0',
    port: process.env.NUXT_ENV_FRONT_PORT
  }
};
