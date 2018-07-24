<template>
  <div>
    <div class="hero is-medium is-primary is-bold">
      <div class="hero-body">
        <h1 class="fuck title">
          Awesome Vue New Update
        </h1>
      </div>
    </div>
    <section class="container section">
      <AwesomeItem
        v-for="item in json"
        v-bind:key="item.Id"
        :item="item" />
    </section>
  </div>
</template>

<script>
import AwesomeItem from '~/components/AwesomeItem.vue';
export default {
  components: {
    AwesomeItem
  },
  async asyncData (ctx) {
    const {data} = await ctx.$axios.get("/api");
    const json = data
      .map(d => ({...d, MergedAt: new Date(d.MergedAt)}))
      .sort((a, b) => b.MergedAt.getTime() - a.MergedAt.getTime());
    return {json};
  },
}
</script>

<style scoped>
.fuck.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 300;
  font-size: 50px;
  letter-spacing: 1px;
  text-align: center;
}
</style>
