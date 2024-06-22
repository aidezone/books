<template>
  <div class="login-container">
    <el-card class="login-card">
      <h2 class="login-title">Login</h2>
      <el-form :model="loginForm" :rules="rules" ref="loginFormRef" label-width="100px">
        <el-form-item label="Username" prop="username">
          <el-input v-model="loginForm.username"></el-input>
        </el-form-item>
        <el-form-item label="Password" prop="password">
          <el-input type="password" v-model="loginForm.password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">Login</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { reactive, ref } from 'vue'
import { accountLogin } from '@/service/auth'
import { useRouter } from 'vue-router';

export default {
  name: 'Login',
  setup() {
    const loginForm = ref({
      username: '',
      password: ''
    })

    const rules = {
      username: [
        { required: true, message: 'Please input username', trigger: 'blur' }
      ],
      password: [
        { required: true, message: 'Please input password', trigger: 'blur' }
      ]
    }

    const loginFormRef = ref(null)
    const router = useRouter();
    const handleLogin = () => {
      loginFormRef.value.validate(async (valid) => {
        if (valid) {
          const resp = await accountLogin(loginForm.value)
          console.log("login response:",resp)
          if (resp.status == 200) {
            localStorage.setItem("token", resp.data.token)
            router.push({ name: 'home' });
            // router.push('/');
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }

    return {
      loginForm,
      rules,
      loginFormRef,
      handleLogin
    }
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-card {
  width: 400px;
  padding: 20px;
}

.login-title {
  text-align: center;
  margin-bottom: 20px;
}
</style>
