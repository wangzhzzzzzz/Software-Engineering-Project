<template>
  <div class="login">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="登录">
          <b-form>
            <b-form-group label="学号">
              <b-form-input
                v-model="$v.user.netid.$model"
                type="number"
                placeholder="输入学号"
                :state="validateState('netid')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('netid')">
                学号不符合要求
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                学号符合要求
              </b-form-valid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入密码"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码不符合要求
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                密码符合要求
              </b-form-valid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                @click="register"
                variant="info"
                block
              >登录</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import {
  required, minLength, maxLength,
} from 'vuelidate/lib/validators';
// import customValidator from '@/helper/validator';// 自定义的验证器

export default {
  data() {
    return {
      user: {
        netid: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      netid: {
        required,
        minLength: minLength(8),
        maxLength: maxLength(8),
      },
      password: {
        minLength: minLength(6),
        maxLength: maxLength(20),
      },
    },
  },
  methods: {
    validateState(name) {
      // 这里是es6 解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      console.log('register');
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
