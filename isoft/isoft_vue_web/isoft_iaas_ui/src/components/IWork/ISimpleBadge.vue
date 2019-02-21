<template>
  <span>
    <span class="badge" :style="mapStyle" @click="resetColor"></span>
    <span >
      <Modal style="float: right;"
        v-model="showFormModal"
        width="500"
        title="选择颜色标记"
        @on-ok="submitColor"
        :transfer="false"
        :mask-closable="false">
        <div>
          <RadioGroup v-model="currentColor">
            <Radio v-for="color in colors" :label="color">
              <span :style="{width:'100px', height:'25px', backgroundColor: color, display:'inline-block'}"></span>
            </Radio>
          </RadioGroup>
        </div>
      </Modal>
    </span>
  </span>

</template>

<script>
  export default {
    name: "ISimpleBadge",
    props: {
      marginRightStyle: {
        type: Number,
        default: 0,
      },
      backgroundColorStyle: {
        type: String,
        default: 'red',
      },
    },
    data(){
      return {
        currentColor:"",
        colors:["#FF0000","#660099","#FFFF00","#FF99FF","#99FF66"],
        showFormModal:false,
        mapStyle:{
          marginRight: `${this.marginRightStyle}px`,
          backgroundColor:this.backgroundColorStyle,
        }
      }
    },
    methods:{
      resetColor:function () {
        this.showFormModal = true;
      },
      submitColor:function () {
        this.$emit("submitColor", this.currentColor)
      }
    }
  }
</script>

<style scoped>
.badge{
  border-radius: 50%;
  width: 10px;
  height: 10px;
  display: inline-block;
}
</style>
