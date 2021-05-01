<template>
  <form>
    科目名:<input v-model="subject" />
    <br />
    科目コード:<input v-model="subjectCode" />
    <br />
    <button @click="register">登録</button>
  </form>
</template>

<script>
export default {
  data() {
    return {
      subject: "",
      subjectCode: "",
    };
  },
  methods: {
    register: function () {
      console.log("post")
      console.log(this.subject);
      console.log(this.subjectCode);
      const subject = {
        subject: this.subject,
        subject_code: this.subjectCode
      }
      this.axios.post("http://localhost:8080/api/subject", subject).then(function (response) {
        console.log(response);
      });
    },
  },
};
</script>
