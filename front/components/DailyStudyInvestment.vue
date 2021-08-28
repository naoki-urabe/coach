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
            label: ["日別学習投資量"],
            backgroundColor: "#4169e1",
            data: [],
          },
        ],
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
      },
    };
  },
  methods: {
    getDailyPeriodDiff: async function (token) {
      const dailyDiff = await axios.get(
        "http://localhost:8080/api/study-log/aggregation/daily",
        { headers: { Authorization: token } }
      );
      return dailyDiff.data;
    },
    setPeriodDiff: async function (dailyPeriodDiff) {
      for (let i = 0; i < dailyPeriodDiff.length; i++) {
        this.chartdata["labels"].splice(i, 0, dailyPeriodDiff[i]["period"]);
        this.chartdata["datasets"][0]["data"].splice(
          i,
          0,
          dailyPeriodDiff[i]["diff"]
        );
      }
    },
  },
  mounted: async function () {
    const token = this.$auth.strategy.token.get();
    const response = await this.getDailyPeriodDiff(token);
    this.setPeriodDiff(response);
    this.renderChart(this.chartdata, this.options);
  },
};
</script>