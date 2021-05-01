<template>
  <h1>{{ subject }}</h1>
  <button @click="kuji">くじびき</button>
</template>

<script>
export default {
  data() {
    return {
      subject: "",
    };
  },
  methods: {
    kuji: async function () {
      let subject = ""
      await this.axios
        .get("http://localhost:8080/api/subject/random")
        .then(function (response) {
          subject = response.data.subject;
        });
      this.subject = subject
      console.log(subject)
    },
  },
};
</script>
