<template>
  <div>
    <Navbar></Navbar>
    <b-container class="my-3">
      <b-row>
        <b-col>
          <div class="d-flex">
            <h1 class="text-truncate">{{rootCategory.name}}</h1>
          </div>
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
          <div class="mt-3 d-flex justify-content-between">
            <b-pagination
              v-model="currentPage"
              :total-rows="rows"
              :per-page="perPage">
            </b-pagination>
            <div>
              <b-button variant="outline-primary" v-b-modal.modal-prevent-closing>New message</b-button>
              <b-modal
                id="modal-prevent-closing"
                ref="modal"
                title="Create new message"
                @show="resetModal"
                @hidden="resetModal"
                @ok="handleOk"
              >
                <form ref="form" @submit.stop.prevent="handleSubmit">
                  <b-form-group
                    :state="textState"
                    label="Message text"
                    label-for="text-input"
                    invalid-feedback="Message is required"
                  >
                    <b-form-textarea
                      id="text-input"
                      v-model="text"
                      :state="textState"
                      placeholder="Enter something..."
                      rows="2"
                      size="sm"
                      required
                    />
                  </b-form-group>
                </form>
              </b-modal>
            </div>
          </div>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-list-group>
            <b-list-group-item v-for="(message, index) in pageMessages" :key="index">
              <div>
                <h6 class="font-weight-normal my-0">{{message.text}}</h6>
                <div>
                  <small><a href="">{{message.author_name}}</a>,</small>
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
    async asyncData({$axios, options}) {
      const rootCategory = await $axios.$get('http://localhost:8080/forum')
      let subcategories, messages
      if (rootCategory) {
        const response_1 = await $axios.$get('http://localhost:8080/forum/subcategories',
          {params: {id: rootCategory.id}})
        const response_2 = await $axios.$get('http://localhost:8080/forum/messages',
          {params: {id: rootCategory.id}})
        subcategories = response_1
        messages = response_2.map(el => {
          const options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            timezone: 'UTC',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric'
          }

          const timestamp = Date.parse(el.posted_at)
          el.posted_at = new Date(timestamp).toLocaleDateString("en-US", options)

          return el
        })
      }

      return {rootCategory: rootCategory, subcategories: subcategories, messages: messages}
    },
    data: () => ({
      currentPage: 1,
      perPage: 20,
      text: '',
      textState: null,
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
    methods: {
      checkFormValidity() {
        const valid = this.$refs.form.checkValidity()
        this.textState = valid ? true : false
        return valid
      },
      resetModal() {
        this.text = ''
        this.textState = null
      },
      handleOk(bvModalEvt) {
        // Prevent modal from closing
        bvModalEvt.preventDefault()
        // Trigger submit handler
        this.handleSubmit()
      },
      async handleSubmit() {
        // Exit when the form isn't valid
        if (!this.checkFormValidity()) {
          return
        }

        let data = new FormData()
        data.append("text", this.text)
        data.append("category_id", this.rootCategory.id)
        console.log(this.rootCategory.id, this.text)

        let msg = await this.$axios.$post('http://localhost:8080/messages', data)

        console.log(msg)

        this.resetModal()

        // Hide the modal manually
        this.$nextTick(() => {
          this.$refs.modal.hide()
        })
      }
    }
  }
</script>

<style scoped>

</style>
