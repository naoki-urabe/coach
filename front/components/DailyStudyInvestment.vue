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
                        label: ["日別学習投資量"],
                        backgroundColor: "#4169e1",
                        data: []
                    }
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
            username: ""
        };
    },
    mounted: async function() {
        this.username = this.$auth.$storage.getLocalStorage("user");
        const response = await this.getDailyPeriodDiff();
        this.setPeriodDiff(response);
        this.renderChart(this.chartdata, this.options);
    },
    methods: {
        getDailyPeriodDiff: async function() {
            const bodyParameters = { "user": this.username };
            const dailyDiff = await this.$axios.post(
                "/study-log/aggregation/daily",
                bodyParameters
            );
            return dailyDiff.data;
        },
        setPeriodDiff: async function(dailyPeriodDiff) {
            for (let i = 0; i < dailyPeriodDiff.length; i++) {
                this.chartdata["labels"].splice(i, 0, dayjs(dailyPeriodDiff[i]["period"]).format("YYYY-MM-DD"));
                this.chartdata["datasets"][0]["data"].splice(
                    i,
                    0,
                    dailyPeriodDiff[i]["diff"]
                );
            }
        }
    }
};
</script>
