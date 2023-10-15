<template>
    <header>
      <Navbar />
    </header>
    <div>
      <h1>Add a New Person</h1>
      <form @submit.prevent="submitData">
        <div>
          <label for="name">Picture:</label>
          <input type="text" id="picture" v-model="formData.picture" required />
        </div>
        <div>
          <label for="name">Name:</label>
          <input type="text" id="name" v-model="formData.name" required />
        </div>
        <div>
          <label for="email">Email:</label>
          <input type="email" id="email" v-model="formData.email" required />
        </div>
        <div>
          <label for="age">Age:</label>
          <input type="number" id="age" v-model="formData.age" required />
        </div>
        <button type="submit">Add Person</button>
      </form>
    </div>
  </template>
  
  <script>
  import Navbar from '~~/components/Navbar.vue';
  export default {
    components: {
      Navbar,
    },
    data() {
      return {
        formData: {
          picture: "",
          name: "",
          email: "",
          age: null,
        },
      };
    },
    methods: {
      async submitData() {
        try {
          const response = await fetch("http://localhost:8000/add", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(this.formData),
          });
          if (response.ok) {
            console.log("Data added successfully");
            // Optionally, you can redirect to a success page or perform other actions.
          } else {
            console.error("Failed to add data");
          }
        } catch (error) {
          console.error("Error:", error);
        }
      },
    },
  };
  </script>  