<template>
  <ul id="example-1">
    <li v-for="some in msgs">
      {{ some.msg }}
    </li>
  </ul>
</template>

<script>

import axios from 'axios';

export default {
  name: 'HelloWorld',
  data () {
    return {
      msgs: "msgs"
    }
  },
  created() {
    axios.get(`http://138.68.82.101:8080/api/msg/`)
    .then(response => {
      this.msgs = response.data
    })
    .catch(e => {
      this.errors.push(e)
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>