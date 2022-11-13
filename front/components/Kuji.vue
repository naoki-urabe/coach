<template>
  <div>
    <v-card>
      <v-data-table
        :headers="headers"
        :items="subjects"
        :items-per-page="100"
      >
      </v-data-table>
    </v-card>
  </div>
</template>
<script>
export default {
  data() {
    return {
      subjects: [],
      headers: [
        { text: "科目名", value: "subject_code" },
        { text: "選択済み", value: "selected" },
        { text: "対象", value: "target" }
      ],
    };
  },
  mounted: async function() {
    await this.getKujiSubjects();
  },
  methods: {
    getAllKujiFromApi: async function() {
    const response = await this.$axios.get(
      "/subject/kuji/all"
    );
    const datas = response.data;
    for(let i=0;i<datas.length;i++) {
      let selected = datas[i].selected;
      let target = datas[i].target;
      datas[i].selected = selected?"〇":"×";
      datas[i].target = target?"〇":"×";
      this.subjects.push(datas[i]);
    }
    },
    getKujiSubjects: async function() {
      await this.getAllKujiFromApi();
    }
  }
}
</script>
