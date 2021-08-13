<template>
  <div>
  <h1>{{ subject }}</h1>
  <button @click="kuji">くじびき</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      subject: "",
    };
  },
  methods: {
    kuji: async function () {
      let subject = ""
      let token = this.$auth.strategy.token.get()
      console.log(token);
      await axios
        .get("http://localhost:8080/api/subject/random", {
          headers: {
            Authorization: token,
          }
        })
        .then(function (response) {
          subject = response.data.subject;
        });
      this.subject = subject
      console.log(subject)
    },
  },
};
</script>
