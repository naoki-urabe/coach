<template>
  <div>
    <h1>{{ subjectName }}</h1>
    <v-btn @click="kuji">
      くじびき
    </v-btn>
  </div>
</template>

<script>
export default {
  data() {
    return {
      subjectName: "",
    };
  },
  methods: {
    kuji: async function () {
      let subjectName = "";
      let token = this.$auth.strategy.token.get();
      await this.$axios
        .get("/subject/random", {
          headers: {
            Authorization: token,
          }
        })
        .then(function (response) {
          subjectName = response.data.subject_name;
        });
      this.subjectName = subjectName;
    },
  },
};
</script>
