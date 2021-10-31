<template>
  <center>
    <v-alert
      v-if="isSuccessRegister"
      type="success"
    >ユーザ登録成功しました</v-alert>
    <h1>ユーザ登録</h1>
    <v-container>
      <v-col cols="5">
        <v-text-field v-model="id" label="ID"></v-text-field>
        <v-text-field v-model="password" label="PASSWORD"></v-text-field>
        <v-btn @click="register">submit</v-btn>
      </v-col>
    </v-container>
  </center>
</template>
<script>
export default {
  auth: "guest",
  data() {
    return {
      id: "",
      password: "",
      isSuccessRegister: false,
    };
  },
  methods: {
    register: async function () {
      const response = await this.$axios({
        method: "post",
        url: '/auth/register',
        data: {
          id: this.id,
          pw: this.password,
        },
      });
      if(response.status === 200){
      	this.isSuccessRegister=true;
	this.id="";
	this.password="";
      }
      console.log(response);
    },
  },
};
</script>
