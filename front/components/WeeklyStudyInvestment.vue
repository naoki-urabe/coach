<script>
import { Bar } from "vue-chartjs";
export default {
  extends: Bar,
  data() {
    return {
      chartdata: {
        labels: ["week1", "week2", "week3", "week4"],
        datasets: [
          {
            label: ["週別学習投資量"],
            backgroundColor: "#4169e1",
            data: [140, 200, 300, 240],
          },
        ],
      },
      options: { responsive: true, maintainAspectRatio: false },
    };
  },
  mounted() {
    this.renderChart(this.chartdata, this.options);
  },
};
</script>