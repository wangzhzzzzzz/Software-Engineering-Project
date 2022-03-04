<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="注册">
          <b-form>
            <b-form-group label="用户名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入名称"
                :state="validateState('name')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('name')">
                用户名不符合要求
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                用户名符合要求
              </b-form-valid-feedback>
            </b-form-group>
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
              >注册</b-button>
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

export default {
  data() {
    return {
      user: {
        name: '',
        netid: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      name: {
        required,
        minLength: minLength(3),
        maxLength: maxLength(16),
      },
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
      // 验证数据
      this.$v.user.$touch();// 没有交互的情况下，也触发表单验证
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      const api = 'http://localhost:1016/api/auth/register';
      this.axios.post(api, { ...this.user }).then((res) => {
        // 保存token
        console.log(res.data);
        // 跳转到主页
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
      });
      console.log('register');
    },
  },
};
</script>

<style lang="scss" scoped>
</style>
