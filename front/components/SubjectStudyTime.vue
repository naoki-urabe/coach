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
          }],
          xAxes:[{
            ticks: {
              maxRotation: 90,
              minRotation: 90
            }
          }
          ]
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
      for(let i=0;i<this.subjects.length;i++){
        const label = this.subjects[i].subject_code;
        let isExists = true;
        for(let j=0;j<aggSubjectStudyTime.length;j++){
          if(label === aggSubjectStudyTime[j].subject_code){
            tmp.push(parseInt(aggSubjectStudyTime[j].AggregationSubjectStudyTime));
            isExists = false;
            break;
          }
        }
        if(isExists) {
          tmp.push(0);
        }
      }
      this.chartdata.datasets[0].data = tmp;
      this.chartdata.datasets[0].data.splice();
    },
    getAggregationSubjectStudyTime: async function() {
      const response = await this.$axios.get("/study-log/aggregation/subject");
      return response.data;
    }
  }
};
</script>
