<template>
  <div class="auth-container">
    <div class="auth-wrapper">
      <!-- 左侧装饰面板 -->
      <div class="decoration-panel">
        <div class="music-note">&#9835;</div>
        <div class="music-note">&#9834;</div>
        <div class="music-note">&#9836;</div>
        <div class="music-note">&#9835;</div>
        <div class="music-note">&#9834;</div>
        <div class="decoration-content">
          <h1>开启你的音乐之旅</h1>
          <p>加入我们，发现更多心动旋律</p>
        </div>
      </div>

      <!-- 右侧表单区域 -->
      <div class="form-panel">
        <!-- 标签页切换 -->
        <div class="tab-switcher">
          <button
            @click="activeTab = 'login'"
            :class="['tab-button', { active: activeTab === 'login' }]">
            登 录
          </button>
          <button
            @click="activeTab = 'register'"
            :class="['tab-button', { active: activeTab === 'register' }]">
            注 册
          </button>
        </div>

        <!-- 登录表单 -->
        <transition name="fade" mode="out-in">
          <div v-if="activeTab === 'login'" key="login" class="form-content">
            <div class="form-header">
              <h2>欢迎回来!</h2>
              <p>进入你的音乐小窝</p>
            </div>
            <el-form
              :model="loginForm"
              :rules="loginRules"
              ref="loginFormRef"
              class="space-y-1"
            >
              <el-form-item prop="username">
                <el-input v-model="loginForm.username" placeholder="请输入用户名" size="large" :prefix-icon="User"
                          @keyup.enter="handleLogin" />
              </el-form-item>
              <el-form-item prop="password">
                <el-input v-model="loginForm.password" type="password" placeholder="请输入密码" size="large"
                          :prefix-icon="Lock" show-password @keyup.enter="handleLogin" />
              </el-form-item>
              <div class="form-options">
                <el-checkbox v-model="rememberMe">记住我</el-checkbox>
                <a href="#" class="link-button">忘记密码?</a>
              </div>
              <el-button type="primary" :loading="loading" @click="handleLogin" class="submit-button" size="large">
                {{ loading ? '登录中...' : '登 录' }}
              </el-button>
            </el-form>
          </div>

          <!-- 注册表单 -->
          <div v-else key="register" class="form-content">
            <div class="form-header">
              <h2>创建新账号</h2>
              <p>只需一步，开启音乐之旅</p>
            </div>
            <el-form
              :model="registerForm"
              :rules="registerRules"
              ref="registerFormRef"
              class="space-y-1"
            >
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-x-4 gap-y-2">
                <el-form-item prop="username">
                  <el-input v-model="registerForm.username" placeholder="登录用户名" size="large" :prefix-icon="User" />
                </el-form-item>
                <el-form-item prop="name">
                  <el-input v-model="registerForm.name" placeholder="社区昵称" size="large" :prefix-icon="User" />
                </el-form-item>
                <el-form-item prop="password">
                  <el-input v-model="registerForm.password" type="password" placeholder="登录密码" size="large"
                            :prefix-icon="Lock" show-password />
                </el-form-item>
                <el-form-item prop="confirmPassword">
                  <el-input v-model="registerForm.confirmPassword" type="password" placeholder="确认密码" size="large"
                            :prefix-icon="Lock" show-password />
                </el-form-item>
              </div>

              <!-- 可选信息区域 -->
              <div class="optional-toggle">
                <el-button link type="primary" @click="showOptionalFields = !showOptionalFields" class="link-button">
                  {{ showOptionalFields ? '收起可选信息' : '填写更多信息 (选填)' }}
                  <el-icon class="el-icon--right transition-transform" :class="{'rotate-180': showOptionalFields}">
                    <ArrowDown />
                  </el-icon>
                </el-button>
              </div>

              <transition name="slide-fade">
                <div v-if="showOptionalFields" class="space-y-2">
                  <div class="grid grid-cols-1 sm:grid-cols-2 gap-x-4 gap-y-2">
                    <el-form-item prop="email">
                      <el-input v-model="registerForm.email" placeholder="邮箱" size="large" :prefix-icon="Message" />
                    </el-form-item>
                    <el-form-item prop="phone">
                      <el-input v-model="registerForm.phone" placeholder="手机号码" size="large" :prefix-icon="Phone" />
                    </el-form-item>
                    <el-form-item prop="gender">
                      <el-select v-model="registerForm.gender" placeholder="性别" size="large" style="width: 100%;">
                        <el-option label="保密" :value="0" />
                        <el-option label="男" :value="1" />
                        <el-option label="女" :value="2" />
                      </el-select>
                    </el-form-item>
                    <el-form-item prop="birthday">
                      <el-date-picker v-model="registerForm.birthday" type="date" placeholder="生日" size="large"
                                      style="width: 100%;" />
                    </el-form-item>
                  </div>
                  <el-form-item prop="location">
                    <el-input v-model="registerForm.location" placeholder="所在地" size="large"
                              :prefix-icon="Location" />
                  </el-form-item>
                  <el-form-item prop="bio">
                    <el-input v-model="registerForm.bio" type="textarea" placeholder="个人简介" :rows="2" />
                  </el-form-item>
                </div>
              </transition>

              <el-button type="primary" :loading="loading" @click="handleRegister" class="submit-button" size="large">
                {{ loading ? '创建中...' : '立即注册' }}
              </el-button>
            </el-form>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage, type FormInstance, type FormRules } from 'element-plus';
import { User, Lock, Message, Phone, Location, ArrowDown } from '@element-plus/icons-vue';
import { authApi } from '@/api';
import { type ModelsLoginRequest, type ModelsRegisterRequest } from '@/api-client';
import { AxiosError } from 'axios';
import { useUserStore } from '@/stores/user';
// --- 通用状态 ---
const router = useRouter();
const userStore = useUserStore();
const loading = ref(false);
const activeTab = ref<'login' | 'register'>('login');

// --- 登录逻辑 ---
const loginFormRef = ref<FormInstance | null>(null);
const rememberMe = ref(false);
const loginForm = reactive<ModelsLoginRequest>({
  username: '',
  password: '',
});

const loginRules = reactive<FormRules>({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
});

const handleLogin = async () => {
  if (!loginFormRef.value) return;
  await loginFormRef.value.validate(async (valid) => {
    if (valid) {
      loading.value = true;
      try {
        const response = await authApi.authLoginPost(loginForm);
        const { token, user_id } = response.data;
        // localStorage.setItem('token', token);
        // localStorage.setItem('userId', JSON.stringify(user_id));
        userStore.setCredentials(token, JSON.stringify(user_id));
        if (rememberMe.value) {
          localStorage.setItem('remembered_username', loginForm.username);
        } else {
          localStorage.removeItem('remembered_username');
        }
        ElMessage.success('登录成功，欢迎回来！');
        router.push('/pub');
      } catch (error) {
        let errorMessage = '登录失败';
        if (error instanceof AxiosError && error.response) {
          const status = error.response.status;
          const dataError = error.response.data?.error;
          if (status === 404) errorMessage = dataError || '用户不存在';
          else if (status === 401) errorMessage = dataError || '密码不正确';
          else errorMessage = dataError || `服务器错误 (代码: ${status})`;
        }
        ElMessage.error(errorMessage);
      } finally {
        loading.value = false;
      }
    }
  });
};

onMounted(() => {
  const rememberedUsername = localStorage.getItem('remembered_username');
  if (rememberedUsername) {
    loginForm.username = rememberedUsername;
    rememberMe.value = true;
  }
});

// --- 注册逻辑 ---
const registerFormRef = ref<FormInstance | null>(null);
const showOptionalFields = ref(false);

interface RegisterFormState extends ModelsRegisterRequest {
  confirmPassword?: string;
}

const registerForm = reactive<RegisterFormState>({
  username: '',
  password: '',
  confirmPassword: '',
  name: '',
  email: '',
  phone: '',
  gender: 0,
  birthday: '',
  location: '',
  bio: '',
  avatar: '',
});

const validatePass2 = (rule: any, value: any, callback: (error?: Error) => void) => {
  if (!value) {
    callback(new Error('请再次输入密码'));
  } else if (value !== registerForm.password) {
    callback(new Error('两次输入密码不一致!'));
  } else {
    callback();
  }
};

const registerRules = reactive<FormRules<RegisterFormState>>({
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur' },
  ],
  name: [{ required: true, message: '请输入您的昵称', trigger: 'blur' }],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 5, message: '密码长度不能少于5位', trigger: 'blur' },
  ],
  confirmPassword: [{ required: true, validator: validatePass2, trigger: 'blur' }],
  email: [{ type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }],
  phone: [{ pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码', trigger: 'blur' }],
});

const handleRegister = async () => {
  if (!registerFormRef.value) return;
  try {
    await registerFormRef.value.validate();
    loading.value = true;
    const payload: ModelsRegisterRequest = {
      username: registerForm.username,
      password: registerForm.password || '',
      name: registerForm.name,
      ...(registerForm.email && { email: registerForm.email }),
      ...(registerForm.phone && { phone: registerForm.phone }),
      ...(registerForm.avatar && { avatar: registerForm.avatar }),
      ...(registerForm.bio && { bio: registerForm.bio }),
      ...(registerForm.location && { location: registerForm.location }),
      ...(registerForm.gender !== undefined && registerForm.gender !== null && { gender: registerForm.gender }),
      ...(registerForm.birthday && { birthday: new Date(registerForm.birthday).toISOString() }),
    };
    await authApi.authRegisterPost(payload);
    ElMessage.success('注册成功！请登录。');
    activeTab.value = 'login';
  } catch (error: any) {
    console.error('注册失败:', error);
    let errorMessage = '注册失败，请稍后重试';
    if (error.response) {
      const status = error.response.status;
      const dataError = error.response.data?.error;
      switch (status) {
        case 409:
          errorMessage = dataError || '该用户名已被注册';
          break;
        case 400:
          errorMessage = dataError || '提交的信息有误，请检查';
          break;
        default:
          errorMessage = dataError || `服务器错误 (代码: ${status})`;
      }
    } else if (error.request) {
      errorMessage = '网络连接失败，请检查您的网络';
    } else {
      errorMessage = error.message;
    }
    ElMessage.error(errorMessage);
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
/* --- Enhanced Kawaii Pink Theme --- */
.auth-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
  background: linear-gradient(135deg, #ffe4e9, #f5d7f0);
  background-size: 300% 300%;
  animation: gradientBG 12s ease infinite;
  overflow: hidden; /* 防止溢出滚动 */
}

.auth-wrapper {
  width: 100%;
  max-width: 68rem;
  display: flex;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 182, 193, 0.6);
  border-radius: 1.75rem;
  box-shadow: 0 20px 40px rgba(255, 105, 180, 0.2);
  overflow: hidden; /* 防止子元素溢出 */
}

.decoration-panel {
  display: none;
  flex: 1;
  align-items: center;
  justify-content: center;
  padding: 2.5rem;
  color: white;
  text-align: center;
  background: linear-gradient(135deg, #ff99ac, #d8b4fe);
  position: relative;
}

@media (min-width: 768px) {
  .decoration-panel {
    display: flex;
  }
}

.decoration-content h1 {
  font-size: 2rem;
  font-weight: 600;
  margin-bottom: 1rem;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.decoration-content p {
  font-size: 1rem;
  opacity: 0.85;
}

.form-panel {
  flex: 1;
  padding: 2.5rem 3.5rem;
  background-color: #fff;
  max-height: calc(100vh - 2rem); /* 减去顶部和底部 padding */
  overflow-y: auto; /* 仅在内容溢出时启用垂直滚动 */
  overflow-x: hidden; /* 防止水平滚动 */
}

.tab-switcher {
  display: flex;
  justify-content: center;
  margin-bottom: 2rem;
  background: #ffe4e9;
  padding: 0.5rem;
  border-radius: 1rem;
}

.tab-button {
  padding: 0.75rem 2rem;
  font-size: 1rem;
  font-weight: 500;
  color: #666;
  background: transparent;
  border: none;
  border-radius: 0.75rem;
  transition: all 0.3s ease;
  position: relative;
}

.tab-button:hover,
.tab-button:focus {
  color: #ff6699;
  background: rgba(255, 105, 180, 0.25);
  outline: none;
}

.tab-button.active {
  color: #fff;
  background: #ff6699;
  box-shadow: 0 2px 8px rgba(255, 105, 180, 0.35);
}

.form-content {
  width: 100%;
  max-width: 450px;
  margin: 0 auto;
  padding: 1rem;
}

.form-header {
  text-align: center;
  margin-bottom: 2rem;
}

.form-header h2 {
  font-size: 1.75rem;
  font-weight: 600;
  color: #2d1b3e;
}

.form-header p {
  margin-top: 0.5rem;
  color: #7b6b8e;
}

:deep(.el-input__wrapper) {
  background-color: #fff0f5 !important;
  border: 1px solid #ff99ac !important;
  border-radius: 0.75rem !important;
  padding: 0.3rem 1rem !important;
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper.is-focus) {
  border-color: #ff6699 !important;
  box-shadow: 0 0 0 3px rgba(255, 102, 153, 0.25) !important;
}

:deep(.el-input__inner) {
  height: 44px;
  color: #2d1b3e;
}

:deep(.el-select .el-input__wrapper) {
  padding: 0.3rem 1rem !important;
}

:deep(.el-textarea__inner) {
  background-color: #fff0f5 !important;
  border: 1px solid #ff99ac !important;
  border-radius: 0.75rem !important;
}

:deep(.el-date-editor .el-input__wrapper) {
  padding: 0.3rem 1rem !important;
}

.submit-button {
  width: 100%;
  height: 48px;
  margin-top: 1.5rem;
  color: white;
  font-weight: 600;
  font-size: 1rem;
  border: none;
  border-radius: 0.75rem;
  background: linear-gradient(135deg, #ff6699, #ff3385);
  box-shadow: 0 4px 12px rgba(255, 102, 153, 0.35);
  transition: all 0.3s ease;
}

.submit-button:hover {
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(255, 102, 153, 0.45);
}

.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1.5rem 0;
  font-size: 0.875rem;
}

.link-button {
  color: #ff3385;
  font-weight: 500;
  transition: color 0.3s ease;
}

.link-button:hover {
  color: #ff6699;
}

.optional-toggle {
  text-align: center;
  margin: 1.5rem 0;
}

/* --- 动画效果 --- */
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

.music-note {
  position: absolute;
  color: rgba(255, 255, 255, 0.4);
  font-size: 2.5rem;
  animation: float 8s infinite ease-in-out;
  user-select: none;
}

.music-note:nth-child(1) {
  top: 15%;
  left: 20%;
  animation-duration: 10s;
}

.music-note:nth-child(2) {
  top: 25%;
  left: 75%;
  animation-duration: 12s;
  animation-delay: -4s;
  font-size: 1.8rem;
}

.music-note:nth-child(3) {
  top: 55%;
  left: 15%;
  animation-duration: 9s;
  animation-delay: -2s;
  font-size: 3rem;
}

.music-note:nth-child(4) {
  top: 75%;
  left: 65%;
  animation-duration: 11s;
  animation-delay: -6s;
}

.music-note:nth-child(5) {
  top: 35%;
  left: 45%;
  animation-duration: 10s;
  animation-delay: -3s;
  font-size: 2rem;
}

@keyframes float {
  0% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.3;
  }
  50% {
    transform: translateY(-20px) rotate(10deg);
    opacity: 0.7;
  }
  100% {
    transform: translateY(0px) rotate(0deg);
    opacity: 0.3;
  }
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-8px);
  opacity: 0;
}

/* 响应式设计 */
@media screen and (max-width: 768px) {
  .auth-wrapper {
    max-width: 90%;
    flex-direction: column;
    border-radius: 1rem;
  }

  .decoration-panel {
    display: none; /* 移动端隐藏装饰面板 */
  }

  .form-panel {
    padding: 1.5rem 2rem;
    max-height: calc(100vh - 2rem); /* 调整移动端高度 */
  }

  .tab-switcher {
    margin-bottom: 1.5rem;
  }

  .form-content {
    max-width: 100%;
  }
}
</style>
