<template>
  <div class="person-card">
    <div class="person-info">
      <img :src="person.picture" alt="Profile Picture" class="person-picture" />
      <div class="person-details">
        <h2>{{ person.name }}</h2>
        <p>Email: {{ person.email }}</p>
        <p>Age: {{ person.age }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    person: Object, // Pass the person object as a prop
  },
};
</script>

<style scoped>
.person-card {
  width: 400px; /* Set a fixed width */
  height: 102px; /* Set a fixed height */
  border: 1px solid #ccc;
  margin: 1rem;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  display: inline-block;
}

.person-info {
  display: flex;
  align-items: center;
  height: 100%; /* Make sure the content fills the entire card */
}

.person-picture {
  width: 100px;
  height: 100px;
  margin-right: 1rem;
}

.person-details {
  flex: 1;
}
</style>
