<template>
  <div>
    <v-container>
      <v-row>
        <v-col cols="4" v-if="!this.$store.state.studyLog.isStart">
          <v-select
            v-model="subjectCode"
            :items="subjects"
            item-text="subject_name"
            item-value="subject_code"
          ></v-select>
        </v-col>
        <v-col cols="4" v-if="this.$store.state.studyLog.isStart">
          <v-text-field v-model="content"></v-text-field>
        </v-col>
        <v-col cols="4" v-if="this.$store.state.studyLog.isStart">
          <v-text-field v-model="comment"></v-text-field>
        </v-col>
        <v-col cols="4">
          <v-btn @click="recordTime">submit</v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-data-table :headers="headers" :items="studyLogs" :items-per-page="20">
    </v-data-table>
  </div>
</template>

<script>
import axios from "axios";
import dayjs from "dayjs";
export default {
  data() {
    return {
      subjectCode: "",
      subjects: [],
      comment: "",
      content: "",
      startTime: null,
      finishTime: null,
      headers: [
        { text: "科目名", value: "subject" },
        { text: "内容", value: "content" },
        { text: "感想", value: "comment" },
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
        content: this.content,
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
      this.content = "";
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
      this.studyLogs.push({
        id: latestStudyLog.id,
        subject: latestStudyLog.subject_code,
        content: null,
        comment: null,
        start: latestStudyLog.study_start_time,
        end: null,
        time: null,
      });
    },
    findIdx: function (latestStudyLog) {
      const id = latestStudyLog.id;
      for (let i = 0; i < this.studyLogs.length; i++) {
        if (id === this.studyLogs[i].id) {
          return i;
        }
      }
    },
    updateLatestStudyLog: function (latestStudyLog) {
      const idx = this.findIdx(latestStudyLog);
      this.$set(
        this.studyLogs[idx],
        "content",
        latestStudyLog.content
      );
      this.$set(
        this.studyLogs[idx],
        "comment",
        latestStudyLog.comment
      );
      this.$set(
        this.studyLogs[idx],
        "end",
        latestStudyLog.study_finish_time
      );
      const end = dayjs(latestStudyLog.study_finish_time);
      const start = dayjs(latestStudyLog.study_start_time);
      const time = parseInt(end.diff(start) / 1000 / 60, 10);
      this.$set(this.studyLogs[idx], "time", time);
    },
    setStudyLogs: async function (studyLogs) {
      for (let i = 0; i < studyLogs.length; i++) {
        let log = studyLogs[i];
        const end = dayjs(log.study_finish_time);
        const start = dayjs(log.study_start_time);
        const time = parseInt(end.diff(start) / 1000 / 60, 10);
        this.studyLogs.push({
          id: log.id,
          subject: log.subject_code,
          content: log.content,
          comment: log.comment,
          start: log.study_start_time,
          end: log.study_finish_time,
          time: time,
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
      if (allStudyLogs.data === null) {
        return [];
      }
      return allStudyLogs.data;
    },
    getAllSubjects: async function (token) {
      const allSubjects = await axios.get("http://localhost:8080/api/subject", {
        headers: { Authorization: token },
      });
      if (allSubjects.data === null) {
        return [];
      }
      return allSubjects.data;
    },
  },
  mounted: async function () {
    let token = this.$auth.strategy.token.get();
    const response = await this.getAllStudyLogs(token);
    this.subjects = await this.getAllSubjects(token);
    this.setStudyLogs(response);
  },
};
</script>