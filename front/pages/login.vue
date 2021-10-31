<template>
  <v-form>
    <center>
      <v-alert v-if="isLogin==='error'" type="error">ログインに失敗しました</v-alert>
      <h1>ログイン</h1>
      <v-container>
        <v-col cols="5">
          <v-text-field v-model="id" label="ID"> </v-text-field>
          <v-text-field v-model="password" label="PASSWORD" type="password"></v-text-field>
          <v-btn @click="login">submit</v-btn>
        </v-col>
        <v-btn :href="registerUserURL">新規ユーザ登録</v-btn>
      </v-container>
    </center>
  </v-form>
</template>
<script>
export default {
  data() {
    return {
      registerUserURL: `http://${this.$config.appHost}:${this.$config.frontPort}/register-user`,
      id: "",
      password: "",
      isLogin: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await this.$auth.loginWith('local', {data: {"Id": this.id, "Pw": this.password}});
        this.$auth.strategy.token.set(response.data.token)
        this.$auth.setUser(response.data.user)
        this.$auth.$storage.setLocalStorage("user", response.data.user)
      } catch(error) {
        this.isLogin="error";
        this.id="";
        this.password="";
        console.log(error)
      }
    },
  },
};
</script>
