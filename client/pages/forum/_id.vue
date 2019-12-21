<template>
  <div>
    <Navbar></Navbar>
    <b-container class="my-3">
      <b-row>
        <b-col>
          <h1>Category name</h1>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-list-group>
            <b-list-group-item>
              <div class="d-flex justify-content-between align-items-center">
                <h6 class="my-0"><a href="">Subcategory name</a></h6>
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
            <b-list-group-item>
              <div class="d-flex justify-content-between align-items-center">
                <h6 class="my-0"><a href="">Subcategory name</a></h6>
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
            <b-list-group-item>
              <div class="d-flex justify-content-between align-items-center">
                <h6 class="my-0"><a href="">Subcategory name</a></h6>
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
            <b-pagination v-model="currentPage" :total-rows="rows"></b-pagination>
          </div>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <b-list-group>
            <b-list-group-item>
              <div>
                <h6 class="my-0">Message text</h6>
                <div>
                  <small><a href="">Author name</a></small>
                  <small>21.12.2019 19:00</small>
                </div>
              </div>
            </b-list-group-item>
            <b-list-group-item>
              <div>
                <h6 class="my-0">Message text</h6>
                <div>
                  <small><a href="">Author name</a></small>
                  <small>21.12.2019 19:00</small>
                </div>
              </div>
            </b-list-group-item>
            <b-list-group-item>
              <div>
                <h6 class="my-0">Message text</h6>
                <div>
                  <small><a href="">Author name</a></small>
                  <small>21.12.2019 19:00</small>
                </div>
              </div>
            </b-list-group-item>
            <b-list-group-item>
              <div>
                <h6 class="my-0">Message text</h6>
                <div>
                  <small><a href="">Author name</a></small>
                  <small>21.12.2019 19:00</small>
                </div>
              </div>
            </b-list-group-item>
            <b-list-group-item>
              <div>
                <h6 class="my-0">Message text</h6>
                <div>
                  <small><a href="">Author name</a></small>
                  <small>21.12.2019 19:00</small>
                </div>
              </div>
            </b-list-group-item>
          </b-list-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <div class="mt-3">
            <b-pagination v-model="currentPage" :total-rows="rows"></b-pagination>
          </div>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script>
  import Navbar from "~/components/Navbar";

  export default {
    data: () => ({
      messages: [],
      rows: 100,
      currentPage: 1
    }),
    components: {
      Navbar
    },
  }
</script>

<style scoped>

</style>
