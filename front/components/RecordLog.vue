<template>
  <div>
    <v-container>
      <v-row>
        <div v-if="!this.$store.state.studyLog.isStart">
          <v-col cols="5">
            <v-select v-model="subjectCode" :items="subjects"></v-select>
          </v-col>
        </div>
        <div v-if="this.$store.state.studyLog.isStart">
          <v-col cols="5">
            <v-text-field v-model="comment"></v-text-field>
          </v-col>
        </div>
        <v-btn @click="recordTime">submit</v-btn>
      </v-row>
    </v-container>
    <v-data-table :headers="headers" :items="studyLogs" :items-per-page="5">
    </v-data-table>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      subjectCode: "",
      subjects: ["network", "db"],
      comment: "",
      startTime: null,
      finishTime: null,
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
  computed: {
    getId() {
      return this.$store.getters["studyLog/getId"];
    },
  },
  methods: {
    createRecord: async function (token) {
      const bodyParameters = {
        subject_code: this.subjectCode,
        study_start_time: new Date(),
      };
      const response = await axios.post(
        "http://localhost:8080/api/study-log/start",
        bodyParameters,
        {
          headers: { Authorization: token },
        }
      );
      this.$store.commit("studyLog/isStartStateChange");
      this.$store.commit("studyLog/changeCurrentId", response.data.id);
      this.subjectCode = "";
    },
    updateCurrentRecord: async function (token, currentId) {
      const bodyParameters = {
        id: currentId,
        comment: this.comment,
        study_finish_time: new Date(),
      };
      const response = await axios.post(
        "http://localhost:8080/api/study-log/finish",
        bodyParameters,
        {
          headers: { Authorization: token },
        }
      );
      this.$store.commit("studyLog/isStartStateChange");
      this.$store.commit("studyLog/changeCurrentId", response.data.id);
      this.comment = "";
    },
    recordTime: async function () {
      let token = this.$auth.strategy.token.get();
      const currentId = this.$store.getters["studyLog/getId"];
      if (currentId === -1) {
        this.createRecord(token);
      } else {
        this.updateCurrentRecord(token, currentId);
      }
    },
  }
};
</script>