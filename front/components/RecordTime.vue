<template>
  <v-data-table :headers="headers" :items="studyLogs" :items-per-page="5">
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        { text: "開始時刻", value: "start" },
        { text: "終了時刻", value: "end" },
        { text: "勉強時間", value: "time" },
      ],
      studyLogs: [
        { start: 1, end: 2, time: 3 },
        { start: 4, end: 5, time: 5 },
        { start: 6, end: 7, time: 8 },
      ],
    };
  },
};
</script>