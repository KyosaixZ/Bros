<template>
    <header>
      <Navbar />
    </header>
    <div>
        <ul>
          <DataFetcher />
        </ul>
    </div>
  </template>
  
  <script>
  import DataFetcher from '~~/components/DataFetcher.vue';
  import Navbar from '~~/components/Navbar.vue';
  
  export default {
    components: {
      DataFetcher,
      Navbar,
    },
  };
  </script>