<template>
  <div>
    <v-alert v-if="isRegister=='success'" type="success">科目登録に成功しました</v-alert>
    <v-alert v-if="isRegister=='error'" type="error">科目登録に失敗しました</v-alert>
    <div>科目コード</div>
    <v-text-field v-model="subjectCode"></v-text-field>
    <div>科目名</div>
    <v-text-field v-model="subjectName"></v-text-field>
    <v-btn @click="registerSubject">登録</v-btn>
  </div>
</template>
<script>
export default {
  data() {
    return {
      subjectCode: "",
      subjectName: "",
      isRegister: "",
    };
  },
  methods: {
    registerSubject: async function () {
    try {
	  let token = this.$auth.strategy.token.get();
	  const headers = {
		Authorization: token,
	  };
	  const bodyParameters = {
		subject_code: this.subjectCode,
		subject_name: this.subjectName,
	  };
	  await this.$axios.post("/subject", bodyParameters, );
	  this.isRegister="success";
	  this.subjectCode = "";
	  this.subjectName = "";
	  } catch(e) {
		this.isRegister="error";
		this.subjectCode = "";
		this.subjectName = "";
	  }
    },
  },
};
</script>
