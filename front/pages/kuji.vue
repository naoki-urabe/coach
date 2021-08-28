<template>
  <div>
  <h1>{{ subjectName }}</h1>
  <button @click="kuji">くじびき</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      subjectName: "",
    };
  },
  methods: {
    kuji: async function () {
      let subjectName = ""
      let token = this.$auth.strategy.token.get()
      await axios
        .get("http://localhost:8080/api/subject/random", {
          headers: {
            Authorization: token,
          }
        })
        .then(function (response) {
          subjectName = response.data.subject_name;
        });
      this.subjectName = subjectName
    },
  },
};
</script>
