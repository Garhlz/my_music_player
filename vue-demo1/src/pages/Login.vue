<template>
  <div class="login-container">
    <div class="login-wrapper">
      <!-- 左侧装饰区域 -->
      <div class="login-decoration">
        <div class="decoration-content">
          <h1>音乐，让生活更美好</h1>
          <p>享受无限音乐，尽在于此</p>
        </div>
      </div>

      <!-- 右侧登录表单 -->
      <div class="login-form-container">
        <el-card class="login-card" :body-style="{ padding: '40px' }">
          <div class="login-header">
            <h2>欢迎回来</h2>
            <p class="subtitle">请登录您的账号</p>
          </div>

          <el-form
            :model="loginForm"
            :rules="rules"
            ref="loginFormRef"
            class="login-form"
          >
            <el-form-item prop="username">
              <el-input
                v-model="loginForm.username"
                placeholder="请输入用户名"
                :prefix-icon="User"
                size="large"
                @keyup.enter="handleLogin"
              />
            </el-form-item>

            <el-form-item prop="password">
              <el-input
                v-model="loginForm.password"
                type="password"
                placeholder="请输入密码"
                :prefix-icon="Lock"
                size="large"
                show-password
                @keyup.enter="handleLogin"
              />
            </el-form-item>

            <div class="login-options">
              <el-checkbox v-model="rememberMe">
                <span class="remember-text">记住我</span>
              </el-checkbox>
              <el-button link type="primary" @click="handleForgotPassword">
                忘记密码？
              </el-button>
            </div>

            <el-button
              type="primary"
              :loading="loading"
              @click="handleLogin"
              class="login-button"
              size="large"
            >
              {{ loading ? '登录中...' : '登录' }}
            </el-button>

            <div class="register-link">
              还没有账号？
              <router-link to="/register" class="register-button">
                立即注册
              </router-link>
            </div>
          </el-form>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
// 1. 引入自动生成的 API Client 和 Element Plus 图标
import { authApi } from '@/api'; // 从我们统一的出口引入
import type { ModelsLoginRequest } from '@/api-client'; // 引入请求的类型定义！
import { Lock, User } from '@element-plus/icons-vue';
import { ElMessage, ElMessageBox, type FormInstance } from 'element-plus';
import { onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { AxiosError } from 'axios'; // 引入 AxiosError 用于更精确的错误类型判断

const router = useRouter();
const loginFormRef = ref<FormInstance>(null);
const loading = ref(false);
const rememberMe = ref(false);

// 表单数据
const loginForm = reactive<ModelsLoginRequest>({
  username: '',
  password: '',
});

// 表单验证规则
const rules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 5, message: '密码长度不能少于5位', trigger: 'blur' },
  ],
};

// 处理忘记密码
const handleForgotPassword = () => {
  ElMessageBox.alert('请联系管理员重置密码', '忘记密码', {
    confirmButtonText: '确定',
    type: 'info',
  });
};

// 3. 核心改造：登录处理函数
const handleLogin = () => {
  if (!loginFormRef.value) return;

  // 1. 先进行表单验证
  loginFormRef.value.validate(async (valid) => {
    // 2. 如果验证通过 (valid 为 true)
    if (valid) {
      loading.value = true;
      try {
        // 3. 在 try...catch 中只包裹 API 请求
        // 【修正】: 直接传递 loginForm，而不是 loginForm.value
        const response = await authApi.authLoginPost(loginForm);

        const { token, user_id } = response.data;
        localStorage.setItem('token', token);
        localStorage.setItem('userId', JSON.stringify(user_id));

        if (rememberMe.value) {
          localStorage.setItem('remembered_username', loginForm.username);
        } else {
          localStorage.removeItem('remembered_username');
        }

        ElMessage({
          type: 'success',
          message: '登录成功，欢迎回来！',
          duration: 2000,
        });

        // 成功后跳转
        router.push('/pub');

      } catch (error) {
        // 4. 这里的 catch 块现在只处理 API 请求的错误
        console.error('API请求失败:', error);
        let errorMessage = '登录失败，请稍后重试';

        if (error instanceof AxiosError) {
          const status = error.response?.status;
          const responseError = error.response?.data?.error;
          switch (status) {
            case 404:
              errorMessage = responseError || '用户不存在';
              break;
            case 401:
              errorMessage = responseError || '密码不正确';
              break;
            default:
              errorMessage = responseError || `发生未知错误 (code: ${status})`;
              break;
          }
        }

        ElMessageBox.alert(errorMessage, '登录失败', {
          confirmButtonText: '确定',
          type: 'error',
        });
      } finally {
        loading.value = false;
      }
    } else {
      // 5. 如果表单验证失败
      console.log('表单验证失败!');
      return false;
    }
  });
};


// 生命周期钩子
onMounted(() => {
  const rememberedUsername = localStorage.getItem('remembered_username');
  if (rememberedUsername) {
    loginForm.username = rememberedUsername;
    rememberMe.value = true;
  }
});
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  background: linear-gradient(120deg, #84fab0 0%, #8fd3f4 100%);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  animation: gradientBG 15s ease infinite;
  background-size: 400% 400%;
}

.login-wrapper {
  display: flex;
  width: 1000px;
  max-width: 100%;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1), 0 5px 15px rgba(0, 0, 0, 0.07);
  backdrop-filter: blur(10px);
  transform: translateY(0);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.login-wrapper:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.12), 0 8px 20px rgba(0, 0, 0, 0.09);
}

.login-decoration {
  flex: 1;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  position: relative;
  overflow: hidden;
}

.login-decoration::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    45deg,
    rgba(255, 255, 255, 0.1) 0%,
    rgba(255, 255, 255, 0) 100%
  );
  transform: translateX(-100%);
  animation: shimmer 3s infinite;
}

.decoration-content {
  text-align: center;
  position: relative;
  z-index: 1;
}

.decoration-content h1 {
  font-size: 2.8em;
  margin-bottom: 20px;
  font-weight: 600;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);
  animation: slideInLeft 0.8s ease-out;
}

.decoration-content p {
  font-size: 1.4em;
  opacity: 0.95;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.15);
  animation: slideInLeft 0.8s ease-out 0.2s backwards;
}

.login-form-container {
  flex: 1;
  padding: 40px;
  background: #ffffff;
  position: relative;
}

.login-card {
  box-shadow: none !important;
  border: none !important;
  animation: fadeIn 0.8s ease-out;
}

.login-header {
  text-align: center;
  margin-bottom: 40px;
}

.login-header h2 {
  font-size: 28px;
  color: #2c3e50;
  margin-bottom: 12px;
  font-weight: 600;
  animation: slideInDown 0.6s ease-out;
}

.subtitle {
  color: #7f8c8d;
  font-size: 16px;
  animation: slideInDown 0.6s ease-out 0.1s backwards;
}

.login-form :deep(.el-input__wrapper) {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.03);
  border-radius: 12px;
  padding: 0 15px;
  transition: all 0.3s ease;
}

.login-form :deep(.el-input__wrapper:hover) {
  transform: translateY(-1px);
  box-shadow: 0 4px 16px 0 rgba(0, 0, 0, 0.06);
}

.login-form :deep(.el-input__inner) {
  height: 48px;
  font-size: 16px;
}

.login-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 25px 0;
  animation: fadeIn 0.6s ease-out 0.3s backwards;
}

.remember-text {
  color: #7f8c8d;
  font-size: 14px;
}

.login-button {
  width: 100%;
  height: 48px;
  border-radius: 12px;
  font-size: 16px;
  margin-top: 20px;
  font-weight: 500;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
  transition: all 0.3s ease;
  animation: slideInUp 0.6s ease-out 0.4s backwards;
}

.login-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
}

.login-button:active {
  transform: translateY(0);
}

.register-link {
  text-align: center;
  margin-top: 25px;
  color: #7f8c8d;
  font-size: 15px;
  animation: fadeIn 0.6s ease-out 0.5s backwards;
}

.register-button {
  color: #667eea;
  text-decoration: none;
  font-weight: 500;
  margin-left: 5px;
  transition: all 0.3s ease;
}

.register-button:hover {
  color: #764ba2;
  text-decoration: none;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .login-wrapper {
    flex-direction: column;
  }

  .login-decoration {
    padding: 30px 20px;
    min-height: 180px;
  }

  .decoration-content h1 {
    font-size: 2.2em;
  }

  .decoration-content p {
    font-size: 1.1em;
  }
}

/* 动画关键帧 */
@keyframes gradientBG {
  0% {
    background-position: 0% 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0% 50%;
  }
}

@keyframes shimmer {
  0% {
    transform: translateX(-100%);
  }
  100% {
    transform: translateX(100%);
  }
}

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideInDown {
  from {
    opacity: 0;
    transform: translateY(-20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
