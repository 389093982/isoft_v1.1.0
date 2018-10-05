<template>
  <div>
    <Row>
      <Col span="4">
        <div style="border: 2px solid #ff5000;height: 400px;">
          AAA
        </div>
      </Col>
      <Col span="16">
        <Carousel v-model="value1" loop>
          <CarouselItem>
            <div class="demo-carousel">
              <img src="../../../assets/ilearning/course/1.jpg" style="width: 100%;height: 400px;"/>
            </div>
          </CarouselItem>
          <CarouselItem>
            <div class="demo-carousel">
              <img src="../../../assets/ilearning/course/2.jpg" style="width: 100%;height: 400px;"/>
            </div>
          </CarouselItem>
          <CarouselItem>
            <div class="demo-carousel">
              <img src="../../../assets/ilearning/course/3.jpg" style="width: 100%;height: 400px;"/>
            </div>
          </CarouselItem>
          <CarouselItem>
            <div class="demo-carousel">
              <img src="../../../assets/ilearning/course/4.jpg" style="width: 100%;height: 400px;"/>
            </div>
          </CarouselItem>
        </Carousel>
      </Col>
      <Col span="4">广告区</Col>
    </Row>
  </div>
</template>

<script>

  export default {
    name: "ChooseCourseType",
    data(){
      return {
        value1: 0
      }
    }
  }
</script>

<style scoped>

</style>
