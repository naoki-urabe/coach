<template>
  <form>
    科目名:<input v-model="subject" />
    <br />
    科目コード:<input v-model="subjectCode" />
    <br />
    <button @click="register">登録</button>
  </form>
  <center>
  <table>
    <thead>
      <th>科目名</th>
      <th>科目コード</th>
    </thead>
    <tbody>
      <tr v-for="(value, key) in subjects" :key="key">
      <td>{{ value.subject }} </td>
      <td>{{ value.subject_code }}</td>
      </tr>
    </tbody>
  </table>
  </center>
</template>

<script>
export default {
  data() {
    return {
      subject: "",
      subjectCode: "",
      subjects: []
    };
  },
  methods: {
    register: function () {
      const subject = {
        subject: this.subject,
        subject_code: this.subjectCode,
      };
      this.axios
        .post("http://localhost:8080/api/subject", subject)
        .then(function (response) {
          console.log(response);
        });
    },
  },
  created: function () {
    this.axios.get("http://localhost:8080/api/subject").then((res) => {
      this.subjects = res.data;
    });
  },
};
</script>
