<!--展示模板-->
<template>
  <!--这里用的是id选择器来绑定css样式的-->
  <div id="app">
    <img src="./assets/logo.png">
    <router-view></router-view>
  </div>
</template>

// 函数和()之间必须有空格；每个字段后不能有空格

<script>
import Hello from './components/ApiData'
export default {
  name: 'xxb',
  components: {
    Hello
  }
}
</script>

<script>
import HelloWorld from './components/HelloWorld'
export default {
  name: 'helloworld',
  components: {
    HelloWorld
  }
}
</script>

<!--样式代码-->
<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
