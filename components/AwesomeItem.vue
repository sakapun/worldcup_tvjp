<template>
  <div class="box">
    <div class="column">
      <div>
        <a :href="item.url"
         class="title"
         target="_blank">{{item.title}}</a>
      </div>
      <div class="subtitle">{{item.mergedAt | formatedDate}}</div>
      <div class="description">{{item.description}}</div>
    </div>
  </div>
</template>

<script>
  import moment from "moment";
  // const lang = navigator && navigator.language
  // moment.locale(lang)
  export default {
    name: "AwesomeItem",
    props: {
      item: {
        title: {
          type: String
        },
        url: {
          type: String
        },
        mergedAt: {
          type: String
        },
        description: {
          type: String
        }
      }
    },
    filters: {
      formatedDate(date) {
        return moment(date).format("ll");
      }
    },
  }
</script>

<style scoped>

</style>
