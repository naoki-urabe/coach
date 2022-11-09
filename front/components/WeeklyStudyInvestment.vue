<script>
import { Bar } from "vue-chartjs";
import dayjs from "dayjs";
export default {
  extends: Bar,
  data() {
    return {
      chartdata: {
        labels: [],
        datasets: [
          {
            label: ["週別学習投資量"],
            backgroundColor: "#4169e1",
            data: [],
          },
        ],
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
      username: ""
    };
  },
  mounted: async function () {
    this.username = this.$auth.$storage.getLocalStorage("user");
    const response = await this.getWeeklyPeriodDiff();
    this.setPeriodDiff(response);
    this.renderChart(this.chartdata, this.options);
  },
  methods: {
    getWeeklyPeriodDiff: async function () {
      const weeklyDiff = await this.$axios.post(
        "/study-log/aggregation/weekly",
        { user: this.username },
      );
      return weeklyDiff.data;
    },
    setPeriodDiff: async function (weeklyPeriodDiff) {
      if(weeklyPeriodDiff == null) {
        return;
      }
      for (let i = 0; i < weeklyPeriodDiff.length; i++) {
        this.chartdata["labels"].splice(i, 0, dayjs(weeklyPeriodDiff[i]["period"]).format("YYYY-MM-DD"));
        this.chartdata["datasets"][0]["data"].splice(
          i,
          0,
          weeklyPeriodDiff[i]["diff"]
        );
      }
    },
  },
};
</script>
