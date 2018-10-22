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
        <div class="has-text-centered">
          <a class="button is-primary is-fullwidth" @click="clicked">
            next
          </a>
        </div>
    </section>
  </div>
</template>

<script>
import AwesomeItem from '~/components/AwesomeItem.vue';
import '~/plugins/amplify';
import {API, graphqlOperation} from "aws-amplify";
export default {
  components: {
    AwesomeItem
  },
  async asyncData (ctx) {
    const query = `query {
        querySort(limit: 20) {
          items {
            mergedAt
            url
            title
            description
          },
          nextToken
        }
      }`;
    const res = await API.graphql(graphqlOperation(query));
    // console.log(res.data.querySort.items[4]);
    return {
      json: res.data.querySort.items,
      nextToken: res.data.querySort.nextToken
    };
  },
  methods: {
    async clicked() {
      const query = `query list {
        querySort(limit: 20, nextToken: "${this.nextToken}") {
          items {
            mergedAt
            url
            title
            description
          },
          nextToken
        }
      }`;
      const res = await API.graphql(graphqlOperation(query));
      // console.log(res.data.listAwesomeLinks.items);
      this.json = this.json.concat(res.data.querySort.items);
      this.nextToken = res.data.querySort.nextToken;
      return res.data.querySort;
    }
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
