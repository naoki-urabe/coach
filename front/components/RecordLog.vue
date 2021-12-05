<template>
  <div>
    <WeeklyStudyInvestment />
    <DailyStudyInvestment />
    <SubjectStudyTime />
    <p>{{ username }}さんの学習記録</p>
    <v-container>
      <v-row>
        <v-col
          v-if="!$store.state.studyLog.isStart"
          cols="4"
        >
          <v-select
            v-model="subjectCode"
            :items="subjects"
            item-text="subject_name"
            item-value="subject_code"
            label="科目名"
            no-data-text="科目を登録してください"
            :menu-props="{
              closeOnClick: true,
              closeOnContentClick: true,
            }"
            :error="isError"
            :error-messages="errMsg"
            @change="getStudyLogs"
          >
            <template #prepend-item>
              <v-list-item
                ripple
                value=""
                @change="getAllStudyLogs"
                @click="subjectCode=''"
              >
                <v-list-item-content>
                  <v-list-item-title>
                    SELECT ALL
                  </v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </template>
          </v-select>
        </v-col>
        <v-col
          v-if="$store.state.studyLog.isStart"
          cols="4"
        >
          <v-text-field
            v-model="content"
            label="内容"
          />
        </v-col>
        <v-col
          v-if="$store.state.studyLog.isStart"
          cols="4"
        >
          <v-text-field
            v-model="comment"
            label="感想"
          />
        </v-col>
        <v-col cols="4">
          <v-btn @click="recordTime">
            submit
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-card>
      <v-data-table
        :headers="headers"
        :items="studyLogs"
        :items-per-page="20"
        :options="pagination"
      >
        <template #[`item.actions`]="{ item }">
          <v-icon
            small
            @click="editItem(item)"
          >
            mdi-pencil
          </v-icon>
          <v-icon
            small
            @click="deleteItem(item)"
          >
            mdi-delete
          </v-icon>
        </template>
      </v-data-table>
      <v-dialog
        v-model="dialogEdit"
        max-width="500px"
      >
        <v-card>
          <v-card-title class="justify-center">
            編集しますか?
          </v-card-title>
          <v-card-actions class="justify-center">
            <v-container>
              <v-row>
                <v-select
                  v-model="editStudyLog.subject_code"
                  :items="subjects"
                  item-text="subject_name"
                  item-value="subject_code"
                  label="科目名"
                  no-data-text="科目を登録してください"
                  :menu-props="{
                    closeOnClick: true,
                    closeOnContentClick: true,
                  }"
                  :error="isError"
                  :error-messages="errMsg"
                />
              </v-row>
              <v-row>
                <v-col cols="6">
                  <v-text-field
                    v-model="editStudyLog.content"
                    label="内容"
                  />
                </v-col>
                <v-col cols="6">
                  <v-text-field
                    v-model="editStudyLog.comment"
                    label="感想"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="6">
                  <v-text-field
                    v-model="editStudyStartTime"
                    label="開始時刻"
                  />
                </v-col>
                <v-col cols="6">
                  <v-text-field
                    v-model="editStudyFinishTime"
                    label="終了時刻"
                  />
                </v-col>
              </v-row>
              <v-row>
                <v-col>
                  <v-btn @click="editConfirm">
                    OK
                  </v-btn>
                  <v-btn @click="dialogEdit=false">
                    キャンセル
                  </v-btn>
                </v-col>
              </v-row>
            </v-container>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog
        v-model="dialogDelete"
        max-width="500px"
      >
        <v-card>
          <v-card-title class="justify-center">
            削除しますか?
          </v-card-title>
          <v-card-actions class="justify-center">
            <v-btn @click="deleteConfirm">
              はい
            </v-btn>
            <v-btn @click="dialogDelete=false">
              いいえ
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-card>
  </div>
</template>

<script>
import dayjs from "dayjs";
import ja from "dayjs/locale/ja";
import RecordLog from "@/components/RecordLog";
import DailyStudyInvestment from "@/components/DailyStudyInvestment.vue";
import WeeklyStudyInvestment from "@/components/WeeklyStudyInvestment.vue";
import SubjectStudyTime from "@/components/SubjectStudyTime.vue";
import Mixin from "../mixins/mixin.js";
dayjs.locale(ja);
export default {
  components: { DailyStudyInvestment, WeeklyStudyInvestment },
  mixins: [ Mixin ],
  data() {
    return {
      dialogEdit: false,
      dialogDelete: false,
      editId: null,
      editStudyLog: {subject_code: "",content: "", comment:"", study_start_time: null,study_finish_time:null},
      editStudyStartTime: null,
      editStudyFinishTime: null,
      deleteId: null,
      subjectCode: "",
      subjects: [],
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
        { text: "編集/削除", value: "actions"},
      ],
      studyLogs: [],
      username: "",
      pagination: {
        sortBy: ["date"],
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
  mounted: async function () {
    this.username = this.$auth.$storage.getLocalStorage("user");
    this.subjects = await this.getAllSubjects();
    this.getAllStudyLogs();
  },
  methods: {
    editItem: async function(item) {
      this.dialogEdit=true;
      this.editId=item.id;
      await this.getSpecificStudyLog();
      this.editStudyStartTime=dayjs(this.editStudyLog.study_start_time).format("YYYY-MM-DD HH:mm");
      this.editStudyFinishTime=dayjs(this.editStudyLog.study_finish_time).format("YYYY-MM-DD HH:mm");
    },
    deleteItem: function(item) {
      this.dialogDelete=true;
      this.deleteId=item.id;
    },
    editConfirm: async function() {
      console.log(dayjs(this.editStudyStartTime).toDate());
      console.log(this.editStudyStartTime);
      const bodyParameters = {
        id: this.editStudyLog.id,
        subject_code: this.editStudyLog.subject_code,
        content: this.editStudyLog.content,
        comment: this.editStudyLog.comment,
        study_start_time: dayjs(this.editStudyStartTime).toDate(),
        study_finish_time: dayjs(this.editStudyFinishTime).toDate(),
        user: this.editStudyLog.user
      };
      await this.$axios.post(
        `/study-log/edit/${this.editId}`,
        bodyParameters,
      );
      this.dialogEdit = false;
      this.getAllStudyLogs();
    },
    deleteConfirm: async function() {	
      await this.$axios.post(
        `/study-log/delete/${this.deleteId}`
      );
      this.dialogDelete=false;
      await this.getAllStudyLogs();
    },
    selectAll: function() {
      this.subjectCode="";
    },
    createRecord: async function () {
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
    updateCurrentRecord: async function (currentId) {
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
        const currentId = this.$store.getters["studyLog/getId"];
        let response = null;
        if (currentId === -1) {
          response = await this.createRecord();
          const latestStudyLog = response.data;
          this.addLatestStudyLog(latestStudyLog);
        } else {
          response = await this.updateCurrentRecord(currentId);
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
        date: dayjs(latestStudyLog.study_start_time).format("YYYY-MM-DD(ddd)"),
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
          date: dayjs(log.study_start_time).format("YYYY-MM-DD(ddd)"),
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
      this.studyLogs = [];
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
      );
      if (subjectStudyLogs === null) {
        return [];
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
    getAllStudyLogs: async function() {
      this.studyLogs = [];
      const response = await this.getAllStudyLogsFromApi();
      this.setStudyLogs(response);
    },
    getSpecificStudyLog: async function() {
      console.log(this.editId);
      const response = await this.$axios.post(`/study-log/${this.editId}`);
      console.log(response.data);
      this.editStudyLog = response.data;
    }
  },
};
</script>
