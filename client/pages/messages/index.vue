<template>
  <div>
    <Navbar></Navbar>
    <b-container class="my-3">
      <div class="messages-table-header">
        <b-row>
          <b-col>
            <h1>Messages</h1>
          </b-col>
        </b-row>
        <b-row>
          <b-col>
            <b-pagination
              v-model="currentPage"
              :total-rows="rows"
              :per-page="perPage"
              aria-controls="messages-table"
            ></b-pagination>
          </b-col>
          <b-col>
            <b-button variant="outline-primary" style="float: right" v-b-modal.modal-prevent-closing>New message</b-button>
          </b-col>
        </b-row>
      </div>
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
            :state="categoryState"
            label="Category"
            label-for="category-input"
            invalid-feedback="Category is required"
          >
            <b-form-select
              id="category-input"
              v-model="category"
              :state="categoryState"
              :options="categories"
              size="sm"
              required
            ></b-form-select>
          </b-form-group>
          <b-form-group
            :state="textState"
            label="Text"
            label-for="text-input"
            invalid-feedback="Text is required"
          >
            <b-form-textarea
              id="text-input"
              v-model="text"
              :state="textState"
              placeholder="Enter something..."
              rows="2"
              size="sm"
              required
            ></b-form-textarea>
          </b-form-group>
        </form>
      </b-modal>
      <b-table
        striped hover
        id="messages-table"
        :items="messages"
        :per-page="perPage"
        :current-page="currentPage"
      ></b-table>
    </b-container>
  </div>
</template>

<script>
    import Navbar from "~/components/Navbar";

    export default {
        data: () => ({
            currentPage: 1,
            perPage: 20,
            category: '',
            categories: ["pudge", "obama", "puke"],
            categoryState: null,
            text: '',
            textState: null,
        }),
        async asyncData({ $axios }) {
            const messages = await $axios.$get('http://localhost:8080/messages')
            return { messages: messages }
        },
        components: {
            Navbar
        },
        mounted() {
            console.log(this.messages)
        },
        computed: {
            rows() {
                return this.messages.length
            }
        },
        methods: {
            checkFormValidity() {
                const valid = this.$refs.form.checkValidity()
                this.textState = valid ? 'valid' : 'invalid'
                return valid
            },
            resetModal() {
                this.category = ''
                this.categoryState = null
                this.text = ''
                this.text = null
            },
            handleOk(bvModalEvt) {
                // Prevent modal from closing
                bvModalEvt.preventDefault()
                // Trigger submit handler
                this.handleSubmit()
            },
            handleSubmit() {
                // Exit when the form isn't valid
                if (!this.checkFormValidity()) {
                    return
                }

                console.log("Message has been pushed")

                // Hide the modal manually
                this.$nextTick(() => {
                    this.$refs.modal.hide()
                })
            }
        }
    }
</script>

<style>
</style>
