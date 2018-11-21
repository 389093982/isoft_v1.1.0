<template>
  <div style="margin: 20px;">
    <h2 style="font-size: 20px;font-weight: 400;margin-bottom: 10px;">友情链接</h2>
    <hr/>
    <ul>
      <li v-for="frindLink in frindLinks"><a :href="frindLink.link_addr">{{frindLink.link_name}}</a></li>
    </ul>
  </div>
</template>

<script>
  import {QueryRandomFrinkLink} from "../../../api"

  export default {
    name: "FrindLink",
    data(){
      return {
        frindLinks:[],
      }
    },
    methods:{
      async refreshRandomFrinkLink (){
        const data = await QueryRandomFrinkLink();
        if(data.status == "SUCCESS"){
          this.frindLinks = data.frind_links;
        }
      }
    },
    mounted:function () {
      this.refreshRandomFrinkLink();
      setInterval(this.refreshRandomFrinkLink, 30000);
    }
  }
</script>

<style scoped>
li{
  height: 32px;
  line-height: 32px;
  margin: 0 4px 5px;
  text-align: center;
  color: #333;
  float: left;
  display: inline;
}
a{
  color: #333;
  display: block;
  height: inherit;
  padding: 0 8px;
}
a:hover{
  color: red;
}
</style>
