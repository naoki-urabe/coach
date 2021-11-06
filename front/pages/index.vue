<template>
  <div>
    <WeeklyStudyInvestment />
    <DailyStudyInvestment />
    <RecordLog />
  </div>
</template>

<script>
import RecordLog from "@/components/RecordLog";
import DailyStudyInvestment from "@/components/DailyStudyInvestment.vue";
import WeeklyStudyInvestment from "@/components/WeeklyStudyInvestment.vue";
export default {
    components: { RecordLog, DailyStudyInvestment, WeeklyStudyInvestment },
};
</script>