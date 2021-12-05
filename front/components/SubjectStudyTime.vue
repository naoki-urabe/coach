<script>
import { Bar } from "vue-chartjs";
import Mixin from "../mixins/mixin.js";
export default {
  extends: Bar,
  mixins: [ Mixin ],
  data() {
    return {
      chartdata: {
        labels: [],
        datasets: [
          {
            label: ["勉強時間"],
            backgroundColor: [
              "rgba(255, 99, 132, 0.2)",
              "rgba(54, 162, 235, 0.2)",
              "rgba(255, 206, 86, 0.2)",
              "rgba(75, 192, 192, 0.2)",
              "rgba(153, 102, 255, 0.2)",
              "rgba(255, 159, 64, 0.2)"],
            data: []
          },
        ]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
            ticks: {
              beginAtZero: true
            }
          }]
        }
      },
      subjects: [],
    };
  },
  mounted: async function() {
    this.subjects = await this.getAllSubjects();
    this.setChartDataLabels();
    await this.setAggregationSubjectStudyTime();
    this.renderChart(this.chartdata, this.options);
  },
  methods: {
    setChartDataLabels: function() {
      const tmp = [];
      for(let i=0;i<this.subjects.length;i++){
        tmp.push(this.subjects[i].subject_name);
      }
      this.chartdata.labels=tmp;
    },
    setAggregationSubjectStudyTime: async function(){
      const tmp = [];
      const aggSubjectStudyTime = await this.getAggregationSubjectStudyTime();
      console.log(aggSubjectStudyTime);
      for(let i=0;i<this.chartdata.labels.length;i++){
        const label = this.chartdata.labels[i];
        for(let j=0;j<aggSubjectStudyTime.length;j++){
          if(label === aggSubjectStudyTime[j].subject_code){
            tmp.push(parseInt(aggSubjectStudyTime[j].AggregationSubjectStudyTime));
          }
        }
      }
      this.chartdata.datasets[0].data = tmp;
    },
    getAggregationSubjectStudyTime: async function() {
      const response = await this.$axios.get("/study-log/aggregation/subject");
      return response.data;
    }
  }
};
</script>
