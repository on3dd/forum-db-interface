<template>
  <div>
    <Navbar></Navbar>
    <b-container class="my-3">
      <b-row>
        <b-col>
          <h1>{{rootCategory.name}}</h1>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-list-group>
            <b-list-group-item v-for="(subcategory, index) in subcategories" :key="index">
              <div class="d-flex justify-content-between align-items-center">
                <h6 class="my-0"><a href="">{{subcategory.name}}</a></h6>
                <div class="d-flex">
                  <div class="mr-2 mr-md-3" style="text-align: right">
                    <small class="d-block">Total themes: 2228</small>
                    <small class="d-block">Total messages: 14880</small>
                  </div>
                  <div style="text-align: right">
                    <small class="d-block text-truncate">Last message text</small>
                    <small class="d-block text-truncate"><a href="">Author name</a></small>
                  </div>
                </div>
              </div>
            </b-list-group-item>

          </b-list-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <div class="mt-3">
            <b-pagination
              v-model="currentPage"
              :total-rows="rows"
              :per-page="perPage">
            </b-pagination>
          </div>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-list-group>
            <b-list-group-item v-for="(message, index) in pageMessages" :key="index">
              <div>
                <h6 class="my-0">{{message.text}}</h6>
                <div>
                  <small><a href="">{{message.author_name}}</a></small>
                  <small>{{message.posted_at}}</small>
                </div>
              </div>
            </b-list-group-item>
          </b-list-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <div class="mt-3">
            <b-pagination
              v-model="currentPage"
              :total-rows="rows"
              :per-page="perPage">
            </b-pagination>
          </div>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
  import Navbar from "~/components/Navbar";

  export default {
    async asyncData ({ $axios }) {
      const rootCategory = await $axios.$get('http://localhost:8080/forum')
      let subcategories, messages
      if (rootCategory) {
        const response_1 = await $axios.$get('http://localhost:8080/forum/subcategories',
          {params:{id: rootCategory.id}})
        const response_2 = await $axios.$get('http://localhost:8080/forum/messages',
          {params:{id: rootCategory.id}})
        subcategories = response_1
        messages = response_2
      }

      return { rootCategory: rootCategory, subcategories: subcategories, messages: messages }
    },
    data: () => ({
      currentPage: 1,
      perPage: 20,
    }),
    components: {
      Navbar
    },
    computed: {
      rows() {
        return this.messages.length
      },
      pageMessages() {
        const from = (this.currentPage - 1) * 20
        return this.messages.slice().splice(from, 20)
      }
    },
  }
</script>

<style scoped>

</style>
