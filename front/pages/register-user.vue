<template>
  <center>
    <v-alert
      v-if="isSuccessRegister === 'success'"
      type="success"
    >ユーザ登録成功しました</v-alert>
    <v-alert
      v-if="isSuccessRegister === 'error'"
      type="error"
    >ユーザは既に存在しています</v-alert>
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
      isSuccessRegister: "",
    };
  },
  methods: {
    register: async function () {
      try {
	      const response = await this.$axios({
						method: "post",
						url: '/auth/register',
						data: {
							id: this.id,
							pw: this.password,
						},
				});
				if(response.status === 200){
						this.isSuccessRegister="success";
						this.id="";
						this.password="";
				}
      } catch (error) {
        this.id="";
				this.password="";
				this.isSuccessRegister="error";
      	console.log(error);
      }
    },
  },
};
</script>
