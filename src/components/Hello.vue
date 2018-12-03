<template>
  <!--使用class来绑定css的样式文件-->
  <div class="hello">
    <!--{{}} 输出对象属性和函数返回值-->
    <h1>{{ msg }}</h1>
    <h1>site : {{site}}</h1>
    <h1>url : {{url}}</h1>
    <h3>{{details()}}</h3>
  </div>
</template>
 
<script>
export default {
  name: 'hello',
  // data用来定义返回数据的属性
  data () {
    return {
      msg: 'hello,xxbandy！',
      site: "bgops",
      url: "https://xxbandy.github.io",
    }
  },
  // 用于定义js的方法
  methods: {
    details: function() {
      return this.site + "数据运维"
    }
  }  
}
</script>

<!--使用css的class选择器[多重样式的生效优先级]-->
<style>
.hello {
  font-weight: normal;
  text-align:center;
  font-size:8pt;
}
h3 
{
  text-align:right;
  font-size:20pt;
  color:red;
}
</style>
