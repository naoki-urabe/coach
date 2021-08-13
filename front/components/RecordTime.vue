<template>
  <v-data-table :headers="headers" :items="studyLogs" :items-per-page="5">
  </v-data-table>
</template>

<script>
export default {
  data() {
    return {
      headers: [
        { text: "科目名", value: "subject" },
        { text: "内容", value: "content" },
        { text: "開始時刻", value: "start" },
        { text: "終了時刻", value: "end" },
        { text: "勉強時間", value: "time" },
      ],
      studyLogs: [
        {
          subject: "network",
          content: "TCP/IP 1-12p",
          start: 1,
          end: 2,
          time: 3,
        },
        {
          subject: "データベース",
          content: "DB 1-12p",
          start: 4,
          end: 5,
          time: 5,
        },
        {
          subject: "Docker",
          content: "docker 1-12p",
          start: 6,
          end: 7,
          time: 8,
        },
      ],
    };
  },
};
</script>