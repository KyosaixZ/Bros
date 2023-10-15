<template>
  <div>
    <div class="person-cards">
      <PersonCard v-for="user in users" :key="user._id" :person="user" />
    </div>
  </div>
</template>

<script>
import PersonCard from '~/components/Card.vue'; // Import the PersonCard component

export default {
  components: {
    PersonCard,
  },
  data() {
    return {
      users: [],
    };
  },
  async created() {
    try {
      const response = await fetch('http://localhost:8000'); // Replace with your API endpoint
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      this.users = data;
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  },
};
</script>
