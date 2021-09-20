<template>
  <div>
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
    };
  },
  methods: {
    registerSubject: async function () {
      let token = this.$auth.strategy.token.get();
      const headers = {
        Authorization: token,
      };
      const bodyParameters = {
        subject_code: this.subjectCode,
        subject_name: this.subjectName,
      };
      await this.$axios.post("/subject", bodyParameters, {
        headers: { Authorization: token },
      });
      this.subjectCode = "";
      this.subjectName = "";
    },
  },
};
</script>