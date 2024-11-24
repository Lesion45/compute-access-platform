<script>
import { register, login } from "../API.js";
export default {
  data() {
    return {
      active: true,
      text:{
        email:"Электронная почта",
        password:"Пароль",
      },  
      form: {
        role:"student",
        email: "",
        password: "",
        password_valid: "",
      },
      errors: {
        email: null,
        password: null,
      },
    };
  },
  computed: {
    isFormValid() {
      return !this.errors.email && !this.errors.password;
    },
  },
  methods: {
    sign_in() {
      this.active = true;
      localStorage.setItem("active", "true");
    },

    sign_up() {
      this.active = false;
      localStorage.setItem("active", "false");
    },

    validateEmail() {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!this.form.email) {
        this.errors.email = "Электронная почта обязательна.";
      } else if (!emailRegex.test(this.form.email)) {
        this.errors.email = "Некорректный формат электронной почты.";
      } else {
        this.errors.email = null;
      }
    },

    validatePassword() {
      if (!this.form.password) {
        this.errors.password = "Пароль обязателен.";
      } else if (this.form.password.length < 6) {
        this.errors.password = "Пароль должен содержать минимум 6 символов.";
      } else if (!this.active && this.form.password !== this.form.password_valid) {
        this.errors.password = "Пароли не совпадают.";
      } else {
        this.errors.password = null;
      }
    },

    async handleSubmit() {
        this.validateEmail();
        this.validatePassword();

        if (!this.isFormValid) {
            alert("Исправьте ошибки в форме.");
            return;
        }

        if (!this.active) {
            this.resetForm();
            const register_data = await register(this.form.email, this.form.password, this.form.role);
            localStorage.setItem("token", register_data.token);
            localStorage.setItem("role", register_data.userRole);
            localStorage.setItem("email", register_data.userEmail);
            localStorage.setItem("id", register_data.userId);
            this.$router.push("/"); 
        } else {
            const login_data = await login(this.form.email, this.form.password);
            localStorage.setItem("token", login_data.token);
            localStorage.setItem("role", login_data.userRole);
            localStorage.setItem("email", login_data.userEmail);
            localStorage.setItem("id", login_data.userId);
            this.$router.push("/"); 
        }
    },

    resetForm() {
      this.form.email = "";
      this.form.password = "";
      this.form.password_valid = "";
      this.errors.email = null;
      this.errors.password = null;
    },
  },
  
  mounted() {
    const activeFromStorage = localStorage.getItem("active");
    this.active = activeFromStorage === "true";
    const token = localStorage.getItem("token");
    if(token && token != ""){
      this.$router.push("/"); 
    }
  },
};
</script>


<template>
    <menu>
      <div class="p_div">
        <p @click="sign_in" :class="{ p2: active }" class="p1">Вход</p>
        <p @click="sign_up" :class="{ p2: !active }" class="p1">Регистрация</p>
      </div>
    </menu>
  
    <main>
      <div class="circle_1"></div>
      <div class="circle_2"></div>
      <div class="circle_3"></div>
  
      <div class="div_form">
        <form class="registration-form" @submit.prevent="handleSubmit">
          <div v-if="!active" class="role_div"> 
            <label> Выберите роль </label>
            <select v-model="form.role" id="role" name="role">
                <option value="student">Студент</option>
                <option value="admin">Администратор</option>
            </select>
          </div>

          <div>
            <label for="email">{{ errors.email == null ? text.email : errors.email}}</label>
            <input
              v-model="form.email"
              type="email"
              id="email"
              name="email"
              placeholder="Введите email"
              @input="validateEmail"
            />
          </div>
  
          <div>
            <label for="password">{{ errors.password == null ? text.password : errors.password}}</label>
            <input
              v-model="form.password"
              type="password"
              id="password"
              name="password"
              placeholder="Введите пароль"
              @input="validatePassword"
            />
          </div>
  
          <div v-if="!active">
            <label for="confirm-password">Подтвердите пароль</label>
            <input
              v-model="form.password_valid"
              type="password"
              id="confirm-password"
              name="confirm-password"
              placeholder="Подтвердите пароль"
              @input="validatePassword"
            />
          </div>
  
          <button type="submit">
            {{ active ? "Вход" : "Регистрация" }}
          </button>
        </form>
      </div>
    </main>
  
    <footer></footer>
  </template>
  



<style src="../style.css"></style>
