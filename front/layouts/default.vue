<template>
  <v-app dark>
    <v-app-bar
      dark
      fixed
      app
    >
      <v-toolbar-title v-text="title" />
      <v-btn href="/">ホーム</v-btn>
      <v-btn href="/register-subject">科目登録</v-btn>
      <v-spacer />
      <v-btn @click="$auth.logout()">logout</v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data () {
    return {
      title: 'Coach'
    }
  }
}
</script>
