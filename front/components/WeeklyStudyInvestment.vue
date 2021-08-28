<script>
import { Bar } from "vue-chartjs";
import axios from "axios";
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
      options: { responsive: true, maintainAspectRatio: false },
    };
  },
  methods: {
    getWeeklyPeriodDiff: async function (token) {
      const weeklyDiff = await axios.get(
        "http://localhost:8080/api/study-log/aggregation/weekly",
        { headers: { Authorization: token } }
      );
      return weeklyDiff.data;
    },
    setPeriodDiff: async function (weeklyPeriodDiff) {
      for (let i = 0; i < weeklyPeriodDiff.length; i++) {
        this.chartdata["labels"].splice(i, 0, weeklyPeriodDiff[i]["period"]);
        this.chartdata["datasets"][0]["data"].splice(
          i,
          0,
          weeklyPeriodDiff[i]["diff"]
        );
      }
    },
  },
  mounted: async function () {
    const token = this.$auth.strategy.token.get();
    const response = await this.getWeeklyPeriodDiff(token);
    this.setPeriodDiff(response);
    this.renderChart(this.chartdata, this.options);
  },
};
</script>