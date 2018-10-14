<template>
  <div>
    <div class="hero is-medium is-primary is-bold">
      <div class="hero-body">
        <h1 class="fuck title" @click="clicked">
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
import '~/plugins/amplify';
import {API, graphqlOperation} from "aws-amplify";
export default {
  components: {
    AwesomeItem
  },
  async asyncData (ctx) {
    const query = `query list {
        orderDesc {
          items {
            MergedAt
            Url
            Title
          },
          nextToken
        }
      }`;
    const res = await API.graphql(graphqlOperation(query));
    // console.log(res.data);
    return {
      json: res.data.orderDesc.items,
      nextToken: res.data.orderDesc.nextToken
    };
  },
  methods: {
    async clicked() {
      const query = `query list {
        orderDesc(nextToken: "${this.nextToken}") {
          items {
            MergedAt
            Url
            Title
          },
          nextToken
        }
      }`;
      const res = await API.graphql(graphqlOperation(query));
      // console.log(res.data.listAwesomeLinks.items);
      this.json = this.json.concat(res.data.orderDesc.items);
      this.nextToken = res.data.orderDesc.nextToken;
      return res.data.orderDesc;
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
