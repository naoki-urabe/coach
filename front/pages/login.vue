<template>
  <v-form>
    <center>
      <h1>ログイン</h1>
      <v-container>
        <v-col cols="5">
          <v-text-field v-model="id" label="ID"> </v-text-field>
          <v-text-field v-model="password" label="PASSWORD"></v-text-field>
          <v-btn @click="login">submit</v-btn>
        </v-col>
        <v-btn href="http://localhost:3001/register-user">新規ユーザ登録</v-btn>
      </v-container>
    </center>
  </v-form>
</template>
<script>
export default {
  data() {
    return {
      id: "",
      password: "",
    };
  },
  methods: {
    async login() {
      try {
        const response = await this.$auth.loginWith('local', {data: {"Id": this.id, "Pw": this.password}});
        this.$auth.strategy.token.set(response.data)
      } catch(error) {
        console.log(error)
      }
    },
  },
};
</script>