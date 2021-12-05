const Mixin = {
  methods: {
    getAllSubjects: async function () {
      const allSubjects = await this.$axios.get("/subject", {
      });
      if (allSubjects.data === null) {
        return [];
      }
      return allSubjects.data;
    },
  }
};
export default Mixin;
