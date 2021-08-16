<template>
  <div>
    <div>科目コード</div>
    <v-text-field v-model="subjectCode"></v-text-field>
    <div>科目名</div>
    <v-text-field v-model="subjectName"></v-text-field>
    <v-btn @click="gdgd">登録</v-btn>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      subjectCode: "",
      subjectName: "",
    };
  },
  methods: {
    gdgd: async function () {
      let token = this.$auth.strategy.token.get();
      console.log(token);
      const headers = {
        Authorization: token,
      };
      const bodyParameters = {
        subject_code: this.subjectCode,
        subject_name: this.subjectName,
      };
      await axios.post("http://localhost:8080/api/subject", bodyParameters, {
        headers: { Authorization: token },
      });
      this.subjectCode = "";
      this.subjectName = "";
    },
  },
};
</script>