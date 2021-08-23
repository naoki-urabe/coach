<template>
  <div>
    <v-container>
      <v-row>
        <div v-if="!this.$store.state.studyLog.isStart">
          <v-col cols="5">
            <v-select v-model="subjectCode" :items="subjects" item-text="subject_name" item-value="subject_code"></v-select>
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
    <v-data-table :headers="headers" :items="studyLogs" :items-per-page="20">
    </v-data-table>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      subjectCode: "",
      subjects: [],
      comment: "",
      startTime: null,
      finishTime: null,
      headers: [
        { text: "科目名", value: "subject" },
        { text: "内容", value: "comment" },
        { text: "開始時刻", value: "start" },
        { text: "終了時刻", value: "end" },
        { text: "勉強時間", value: "time" },
      ],
      studyLogs: [],
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
      return response;
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
      return response;
    },
    recordTime: async function () {
      let token = this.$auth.strategy.token.get();
      const currentId = this.$store.getters["studyLog/getId"];
      let response = null;
      if (currentId === -1) {
        response = await this.createRecord(token);
        const latestStudyLog = response.data;
        this.addLatestStudyLog(latestStudyLog);
      } else {
        response = await this.updateCurrentRecord(token, currentId);
        const latestStudyLog = response.data;
        this.updateLatestStudyLog(latestStudyLog);
      }
    },
    addLatestStudyLog: async function (latestStudyLog) {
      console.log(latestStudyLog);
      this.studyLogs.push({
        subject: latestStudyLog.subject_code,
        comment: null,
        start: latestStudyLog.study_start_time,
        end: null,
        time: null,
      });
    },
    updateLatestStudyLog: function (latestStudyLog) {
      this.$set(
        this.studyLogs[latestStudyLog.id - 1],
        "comment",
        latestStudyLog.comment
      );
      this.$set(
        this.studyLogs[latestStudyLog.id - 1],
        "end",
        latestStudyLog.study_finish_time
      );
    },
    setStudyLogs: async function (studyLogs) {
      for (let i = 0; i < studyLogs.length; i++) {
        let log = studyLogs[i];
        this.studyLogs.push({
          subject: log.subject_code,
          comment: log.comment,
          start: log.study_start_time,
          end: log.study_finish_time,
          time: null,
        });
      }
    },
    getAllStudyLogs: async function (token) {
      const allStudyLogs = await axios.get(
        "http://localhost:8080/api/study-log/all",
        {
          headers: { Authorization: token },
        }
      );
      return allStudyLogs.data;
    },
    getAllSubjects: async function (token) {
      const allSubjects = await axios.get("http://localhost:8080/api/subject", {
        headers: { Authorization: token },
      });
      return allSubjects.data;
    },
  },
  mounted: async function () {
    let token = this.$auth.strategy.token.get();
    const response = await this.getAllStudyLogs(token);
    this.subjects = await this.getAllSubjects(token);
    this.setStudyLogs(response);
    this.subjects.splice(0,0);
  },
};
</script>