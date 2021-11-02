<template>
  <div>
    <p>{{ username }}さんの学習記録</p>
    <v-container>
      <v-row>
        <v-col cols="4" v-if="!this.$store.state.studyLog.isStart">
	  <v-select
            v-model="subjectCode"
            :items="subjects"
            item-text="subject_name"
            item-value="subject_code"
            label="科目名"
            @change="getStudyLogs"
	    no-data-text="科目を登録してください"
	    :menu-props="{
		closeOnClick: true,
		closeOnContentClick: true,
	    }"
	    :error="isError"
	    :error-messages="errMsg"
          >
		  <template v-slot:prepend-item>
			<v-list-item ripple @change="getAllStudyLogs" @click="subjectCode=''" value="">
				<v-list-item-content>
					<v-list-item-title>
						SELECT ALL
					</v-list-item-title>
				</v-list-item-content>
			</v-list-item>
		  </template>
	  </v-select>
        </v-col>
        <v-col cols="4" v-if="this.$store.state.studyLog.isStart">
          <v-text-field v-model="content" label="内容"></v-text-field>
        </v-col>
        <v-col cols="4" v-if="this.$store.state.studyLog.isStart">
          <v-text-field v-model="comment" label="感想"></v-text-field>
        </v-col>
        <v-col cols="4">
          <v-btn @click="recordTime">submit</v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-card>
      <v-data-table :headers="headers" :items="studyLogs" :items-per-page="20" :options="pagination">
      </v-data-table>
    </v-card>
  </div>
</template>

<script>
import dayjs from "dayjs";
export default {
  data() {
    return {
      subjectCode: "",
      subjects: [{subject_code: "",subject_name:"SELECT ALL"}],
      comment: "",
      content: "",
      startTime: null,
      finishTime: null,
      headers: [
        { text: "日付", value: "date"},
        { text: "科目名", value: "subject" },
        { text: "内容", value: "content" },
        { text: "感想", value: "comment" },
        { text: "開始時刻", value: "start" },
        { text: "終了時刻", value: "end" },
        { text: "勉強時間", value: "time" },
      ],
      studyLogs: [],
      username: "",
      pagination: {
        sortBy: ['date'],
        sortDesc: [true]
      },
      isError: false,
      errMsg: ""
    };
  },
  computed: {
    getId() {
      return this.$store.getters["studyLog/getId"];
    },
    /*getIsStart() {
      return this.$store.getters["studyLog/getIsStart"];
    }*/
  },
  methods: {
    selectAll: function() {
    	this.subjectCode="";
    },
    createRecord: async function (token) {
      try {
	      const bodyParameters = {
		subject_code: this.subjectCode,
		study_start_time: new Date(),
		user: this.username
	      };
	      const response = await this.$axios.post(
		"/study-log/start",
		bodyParameters,
	      );
	      this.$store.commit("studyLog/isStartStateChange");
	      this.$store.commit("studyLog/changeCurrentId", response.data.id);
	      this.subjectCode = "";
	      this.isError=false;
	      this.errMsg="";
	      return response;
      } catch(err) {
      	console.log(err);
	this.isError=true;
	this.errMsg="科目の値が不正です";
      }
    },
    updateCurrentRecord: async function (token, currentId) {
      const bodyParameters = {
        id: currentId,
        content: this.content,
        comment: this.comment,
        study_finish_time: new Date(),
      };
      const response = await this.$axios.post(
        "/study-log/finish",
        bodyParameters,
      );
      this.$store.commit("studyLog/isStartStateChange");
      this.$store.commit("studyLog/changeCurrentId", response.data.id);
      this.content = "";
      this.comment = "";
      return response;
    },
    recordTime: async function () {
      try {
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
      } catch(err) {
      	console.log(err);
      }
    },
    addLatestStudyLog: async function (latestStudyLog) {
      this.studyLogs.push({
        id: latestStudyLog.id,
        date: dayjs(latestStudyLog.study_start_time).format("YYYY-MM-DD"),
        subject: latestStudyLog.subject_code,
        content: null,
        comment: null,
        start: dayjs(latestStudyLog.study_start_time).format("HH:mm"),
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
        dayjs(latestStudyLog.study_finish_time).format("HH:mm")
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
          date: dayjs(log.study_start_time).format("YYYY-MM-DD"),
          subject: log.subject_code,
          content: log.content,
          comment: log.comment,
          start: dayjs(log.study_start_time).format("HH:mm"),
          end: dayjs(log.study_finish_time).format("HH:mm"),
          time: time,
        });
      }
    },
    getStudyLogs: async function() {
      this.studyLogs = []
      if(this.subjectCode == "") {
        this.getAllStudyLogs();
      } else {
        const response = await this.getSubjectStudyLogs();
        this.setStudyLogs(response);
      }
    },
    getSubjectStudyLogs: async function() {
      const subjectStudyLogs = await this.$axios.post(
        "/study-log/subject",
        {user: this.username,subjectCode: this.subjectCode},
      )
      if (subjectStudyLogs === null) {
        return []
      }
      return subjectStudyLogs.data;
    },
    getAllStudyLogsFromApi: async function () {
      const allStudyLogs = await this.$axios.post(
        "/study-log/all",
        {user: this.username},
      );
      if (allStudyLogs.data === null) {
        return [];
      }
      return allStudyLogs.data;
    },
    getAllSubjects: async function (token) {
      const allSubjects = await this.$axios.get("/subject", {
      });
      if (allSubjects.data === null) {
        return [];
      }
      return allSubjects.data;
    },
    getAllStudyLogs: async function() {
	    const response = await this.getAllStudyLogsFromApi();
	    this.setStudyLogs(response);
    }
  },
  mounted: async function () {
    this.username = this.$auth.$storage.getLocalStorage("user");
    let token = this.$auth.strategy.token.get();
    this.subjects = await this.getAllSubjects(token);
    this.getAllStudyLogs();
  },
};
</script>
