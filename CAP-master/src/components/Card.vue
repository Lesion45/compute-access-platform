<script >
import { get_computer, reserve_computer, relieve_computer } from "../API.js";
export default {
  data() {
    return {
      bonus_info: false,
      date: null,
      token:null,
      ssh:null,
      access: false,
      load: false,
    };
  },

  props:{
    computer_info:{}
  },

  computed: {
    async check_access(){
      try{
        this.load = true
        if(await get_computer(computer_info.id, token) != 'error'){
          this.access = true;
          this.load = false;
          return true;
        }
        this.load = false;
        return false;
      }
      catch{
        this.load = false;
        return false;
      }
    }
  },

  methods: {
    card_click(){
      this.bonus_info = true;
    },
    close(e) {
      if (!this.$el.contains(e.target) && !e.target.closest('.pop-up') && !e.target.closest('.stend1') && !e.target.closest('.grey1_div')) { 
        this.bonus_info = false;
      }
    },
    copy(){
      navigator.clipboard.writeText(this.ssh);
      this.copy_flag = true;
      setTimeout(() => {
        this.copy_flag = null;
      }, 2500);
    },
    async booking(){
      try{
        this.load = true
        if(await reserve_computer(this.computer_info.id, this.token) && this.token != null){
          const data = await get_computer(this.computer_info.id, this.token);
          if(data != "error"){
            this.ssh = data.ssh;
            this.access = true;
          }
        }
        this.load = false;
      }
      catch{
        this.load = false;
      }
    },
    async unbooking(){
      try{
        this.load = true
        if(await relieve_computer(this.computer_info.id, this.token) && this.token != null){
          this.access = false;
        }
        this.load = false;
      }
      catch{
        this.load = false;
      }
    }
  },
  
  mounted() {
    document.addEventListener('click', this.close.bind(this));
    this.token = localStorage.getItem('token');
  },

  beforeDestroy () {
    document.removeEventListener('click',this.close);
  },
};
</script>

<template>
<div @click="card_click" class="stend1"> 
    <img class="comp_img" src="../imgs/Codesandbox.svg">
    <div class="ch_div"> 
        <p> {{ computer_info.name }} </p>
        <div class="ch"> 
            <img class="os_img" src="../imgs/Monitor.svg"> <p> {{ computer_info.os }} </p>
            <img class="prc_img" src="../imgs/Cpu.svg"> <p> {{ computer_info.cpu }} </p>
            <img class="mem_img" src="../imgs/Disc.svg">  <p> {{ computer_info.ram }} </p>
            <div :class="{red_status: computer_info.status === 'Забронирован'}" class="status_div"> <p class="p_grey1"> {{ computer_info.status }} </p> </div>
        </div>    
    </div>
</div>

<div v-if="bonus_info" class="stend_info_bg"></div>
<div v-if="bonus_info" class="pop-up">
  <div class="stend_info"> 
      <div class="comp_img_ch1_div">
          <img class="comp_img" src="../imgs/Codesandbox.svg">
          <div class="ch_div"> 
              <p> <b> {{ computer_info.name }} </b> </p>
              <div class="ch"> 
                  <img class="os_img" src="../imgs/Monitor.svg"> <p> {{ computer_info.os }} </p>
                  <div> </div> <div> </div> <div> </div> <div> </div>
                  <img class="prc_img" src="../imgs/Cpu.svg"> <p> {{ computer_info.cpu }} </p>
                  <div> </div> <div> </div> <div> </div> <div> </div>
                  <img class="mem_img" src="../imgs/Disc.svg">  <p> {{ computer_info.ram }} </p>
              </div>    
          </div>
          </div>
          <div @click="booking" v-if="computer_info.status === 'Свободен' && !access && !load" class="grey1_div"> <p class="p_grey1"> Забронировать </p> </div>
          <div @click="unbooking" v-if="computer_info.status === 'Забронирован' && check_access && access" class="grey1_div"> <p class="p_grey1"> Снять бронь </p> </div>
          <img v-if="load" class="load_img" src="../imgs/Loader.svg">
          <div class="text_div" v-if="access">
            <!-- <p>id стенда: {{ computer_info.id }}</p> -->
            <p
            :class="{
              copy_animate_on: copy_flag,
              copy_animate_off: copy_flag,
            }"
            class="copy_alert"
            >
              Скопированно!
            </p>
            <p class="copy" @click="copy()"> {{ ssh }}</p>
            <p class="sub_text">нажмите чтобы скопировать ssh ключ и воспользуйтесь консолью для подключения</p>
          </div>
  </div>
</div>
</template>

<style src="../style.css">
</style>
