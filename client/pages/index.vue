<template>
  <div>
    <Navbar></Navbar>
    <b-container>

    </b-container>
  </div>
</template>

<script>
    import Navbar from "../components/Navbar";

    export default {
        data: () => ({
           messages: [],
        }),
        components: {
            Navbar
        },
    }
</script>

<style>
</style>
