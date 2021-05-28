<template>
  <div class="hello">
    <ul>
      <li v-for="item in items" :key="item.id">
        <dl>
          <dt>id</dt>
          <dd>{{ item.id }}</dd>
          <dt>name</dt>
          <dd>{{ item.name }}</dd>
        </dl>
      </li>
    </ul>
    <div>
      <dl>
        <dt>id</dt>
        <dd><input v-model="input.id"/></dd>
        <dt>name</dt>
        <dd><input v-model="input.name"/></dd>
      </dl>
      <button @click="add">追加</button>
    </div>
  </div>
</template>

<script>
import {put, get} from "axios";

export default {
  name: "Pets",
  data() {
    return {
      input: {
        id: "",
        name: "",
      },
      items: [],
    };
  },
  async mounted() {
    await this.update()
  },
  methods: {
    async add() {
      await put("http://localhost:8081/v1/pets", {
        id: this.input.id,
        name: this.input.name,
      });
      await this.update()
    },
    async update() {
      const res = await get("http://localhost:8081/v1/pets");
      this.items = res.data;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
