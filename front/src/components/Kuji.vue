<template>
  <h1>{{count}}</h1>
  <button @click="kuji">くじびき</button>
</template>

<script>
export default ({
    data() {
        return {count: 0}
    },
    methods: {
        kuji:function() {
            this.count++;
        }
    }
})
</script>
