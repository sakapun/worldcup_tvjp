<template>
  <div class="box">
    <div class="column">
      <div>
        <a :href="item.Url"
         class="title"
         target="_blank">{{item.Title}}</a>
      </div>
      <div class="subtitle">{{item.MergedAt | formatedDate}}</div>
      <div class="description">{{item.Description}}</div>
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
        Title: {
          type: String
        },
        Url: {
          type: String
        },
        MergedAt: {
          type: String
        },
        Description: {
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
