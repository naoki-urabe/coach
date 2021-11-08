<template>
  <div>
    <v-alert
      v-if="isRegister=='success'"
      type="success"
    >
      科目登録に成功しました
    </v-alert>
    <v-alert
      v-if="isRegister=='error'"
      type="error"
    >
      科目登録に失敗しました
    </v-alert>
    <v-container>
      <v-row>
        <v-col cols="4">
          <v-text-field
            v-model="subjectCode"
            label="科目コード"
          />
        </v-col>
        <v-col cols="4">
          <v-text-field
            v-model="subjectName"
            label="科目名"
          />
        </v-col>
        <v-col cols="4">
          <v-btn @click="registerSubject">
            登録
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <v-card>
      <v-data-table
        :headers="headers"
        :items="subjects"
        :item-per-page="20"
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
            編集しますか？
          </v-card-title>
          <v-card-actions class="justify-center">
            <v-container>
              <v-row>
                <v-col cols="6">
                  <v-text-field
                    v-model="editSubjectCode"
                    label="科目コード"
                  />
                </v-col>
                <v-col cols="6">
                  <v-text-field
                    v-model="editSubjectName"
                    label="科目名"
                  />
                </v-col>
              </v-row>
              <v-btn @click="editConfirm">
                はい
              </v-btn>
              <v-btn @click="dialogEdit=false">
                いいえ
              </v-btn>
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
            削除しますか？
          </v-card-title>
          <v-card-actions class="justify-center">
            </v-row>
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
export default {
  data() {
    return {
      dialogEdit: false,
      dialogDelete: false,
      editId: null,
      deleteId: null,
      editSubject: null,
      editSubjectCode: null,
      editSubjectName: null,
      subjectCode: "",
      subjectName: "",
      isRegister: "",
      headers: [
        {text: "科目コード", value: "subjectCode"},
        {text: "科目名", value: "subjectName"},
        {text: "編集/削除", value: "actions"},
      ],
      subjects: [],
    };
  },
  mounted: async function() {
    await this.getAllSubject();
  },
  methods: {
    editItem: async function(item) {
      this.dialogEdit=true;
      this.editId=item.subjectCode;
      this.editSubject = (await this.$axios.get(`/subject/${this.editId}`)).data;
      this.editSubjectCode = this.editSubject.subject_code;
      this.editSubjectName = this.editSubject.subject_name;
    },
    editConfirm: async function() {
      const bodyParameter = {subject_code: this.editSubjectCode, subject_name: this.editSubjectName};
      await this.$axios.post(`/subject/edit/${this.editId}`,bodyParameter);
      this.dialogEdit=false;
      await this.getAllSubject();
    },
    deleteItem: async function(item) {
      this.dialogDelete=true;
      this.deleteId=item.subjectCode;
    },
    deleteConfirm: async function() {
      await this.$axios.post(`/subject/delete/${this.deleteId}`);
      await this.getAllSubject();
      this.dialogDelete=false;
    },
    registerSubject: async function () {
      try {
        const bodyParameters = {
          subject_code: this.subjectCode,
          subject_name: this.subjectName,
        };
        await this.$axios.post("/subject", bodyParameters, );
        this.isRegister="success";
        this.subjectCode = "";
        this.subjectName = "";
        await this.getAllSubject();
      } catch(e) {
        this.isRegister="error";
        this.subjectCode = "";
        this.subjectName = "";
      }
    },
    getAllSubjectFromApi: async function() {
      const res = await this.$axios.get("/subject");
      return res.data;
    },
    setSubject: async function(subjects) {
      if(subjects == null){
        return;
      }
      for(let i=0;i<subjects.length;i++) {
        if(subjects[i] == null) {
          continue;
        }
        this.subjects.splice(0,0,{subjectCode: subjects[i].subject_code, subjectName: subjects[i].subject_name});
      }
    },
    getAllSubject: async function() {
      this.subjects = [];
      const subjects = await this.getAllSubjectFromApi();
      await this.setSubject(subjects);
    }
  }
};
</script>
