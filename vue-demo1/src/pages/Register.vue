<template>
  <div class="register-container">
	<div class="register-wrapper">
	  <!-- 左侧装饰区域 -->
	  <div class="register-decoration">
		<div class="decoration-content">
		  <h1>开启音乐之旅</h1>
		  <p>注册账号，探索更多精彩</p>
		</div>
	  </div>

	  <!-- 右侧注册表单 -->
	  <div class="register-form-container">
		<el-card class="register-card" :body-style="{ padding: '40px' }">
		  <div class="register-header">
			<h2>创建账号</h2>
			<p class="subtitle">请填写以下信息完成注册</p>
		  </div>

		  <el-form
			  :model="registerForm"
			  :rules="rules"
			  ref="registerFormRef"
			  class="register-form"
		  >
			<el-form-item prop="username">
			  <el-input
				  v-model="registerForm.username"
				  placeholder="用户名"
				  :prefix-icon="User"
				  size="large"
			  />
			</el-form-item>

			<el-form-item prop="password">
			  <el-input
				  v-model="registerForm.password"
				  type="password"
				  placeholder="密码"
				  :prefix-icon="Lock"
				  size="large"
				  show-password
			  />
			</el-form-item>

			<el-form-item prop="confirmPassword">
			  <el-input
				  v-model="registerForm.confirmPassword"
				  type="password"
				  placeholder="确认密码"
				  :prefix-icon="Lock"
				  size="large"
				  show-password
			  />
			</el-form-item>

			<el-form-item prop="name">
			  <el-input
				  v-model="registerForm.name"
				  placeholder="昵称"
				  :prefix-icon="User"
				  size="large"
			  />
			</el-form-item>

			<el-form-item prop="email">
			  <el-input
				  v-model="registerForm.email"
				  placeholder="邮箱"
				  :prefix-icon="Message"
				  size="large"
			  />
			</el-form-item>

			<el-form-item prop="phone">
			  <el-input
				  v-model="registerForm.phone"
				  placeholder="手机号码"
				  :prefix-icon="Phone"
				  size="large"
			  />
			</el-form-item>

			<el-form-item prop="gender">
			  <el-select
				  v-model="registerForm.gender"
				  placeholder="性别"
				  size="large"
				  style="width: 100%"
			  >
				<el-option label="保密" :value="0"/>
				<el-option label="男" :value="1"/>
				<el-option label="女" :value="2"/>
			  </el-select>
			</el-form-item>

			<el-form-item prop="birthday">
			  <el-date-picker
				  v-model="registerForm.birthday"
				  type="date"
				  placeholder="选择生日"
				  size="large"
				  style="width: 100%"
			  />
			</el-form-item>

			<el-form-item prop="location">
			  <el-input
				  v-model="registerForm.location"
				  placeholder="所在地"
				  :prefix-icon="Location"
				  size="large"
			  />
			</el-form-item>

			<el-form-item prop="bio">
			  <el-input
				  v-model="registerForm.bio"
				  type="textarea"
				  placeholder="个人简介"
				  :rows="3"
			  />
			</el-form-item>

			<el-button
				type="primary"
				:loading="loading"
				@click="handleRegister"
				class="register-button"
				size="large"
			>
			  {{ loading ? '注册中...' : '立即注册' }}
			</el-button>

			<div class="login-link">
			  已有账号？
			  <router-link to="/login" class="login-button">
				立即登录
			  </router-link>
			</div>
		  </el-form>
		</el-card>
	  </div>
	</div>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {useRouter} from 'vue-router'
import {ElMessage} from 'element-plus'
import {User, Lock, Message, Phone, Location} from '@element-plus/icons-vue'
import {register} from 'src/api/axiosFile'

const router = useRouter()
const registerFormRef = ref(null)

const registerForm = ref({
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
  avatar: '/assets/avatars/default-user.jpg',
  status: 1
})

const validatePass2 = (rule, value, callback) => {
  if (value === '') {
	callback(new Error('请再次输入密码'))
  } else if (value !== registerForm.value.password) {
	callback(new Error('两次输入密码不一致!'))
  } else {
	callback()
  }
}

const rules = {
  username: [
	{required: true, message: '请输入用户名', trigger: 'blur'},
	{min: 3, max: 20, message: '长度在 3 到 20 个字符', trigger: 'blur'}
  ],
  password: [
	{required: true, message: '请输入密码', trigger: 'blur'},
	{min: 6, message: '密码长度不能少于6位', trigger: 'blur'}
  ],
  confirmPassword: [
	{required: true, validator: validatePass2, trigger: 'blur'}
  ],
  email: [
	{type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur'}
  ],
  phone: [
	{pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码', trigger: 'blur'}
  ],
  name: [
	{max: 50, message: '姓名长度不能超过50个字符', trigger: 'blur'}
  ],
  bio: [
	{max: 500, message: '个人简介不能超过500个字符', trigger: 'blur'}
  ],
  location: [
	{max: 100, message: '地址长度不能超过100个字符', trigger: 'blur'}
  ]
}

const loading = ref(false)

const handleRegister = async () => {
  if (!registerFormRef.value) return

  try {
	await registerFormRef.value.validate()

	loading.value = true
	const {confirmPassword, ...registerData} = registerForm.value

	// 转换日期格式
	if (registerData.birthday) {
	  registerData.birthday = registerData.birthday.toISOString().split('T')[0]
	}

	const result = await register(registerData)
	if (result.data.message) {
	  ElMessage.success('注册成功')
	  router.push('/login')
	} else {
	  throw new Error(result.data.error || '注册失败')
	}
  } catch (error) {
	console.error('注册失败:', error)
	ElMessage.error(error.message || '注册失败，请稍后重试')
  } finally {
	loading.value = false
  }
}
</script>

<style scoped>
.register-container {
  min-height: 100vh;
  background: linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  animation: gradientBG 15s ease infinite;
  background-size: 400% 400%;
}

.register-wrapper {
  display: flex;
  width: 1200px;
  max-width: 100%;
  background: rgba(255, 255, 255, 0.95);
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1), 0 5px 15px rgba(0, 0, 0, 0.07);
  backdrop-filter: blur(10px);
  transform: translateY(0);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.register-wrapper:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.12), 0 8px 20px rgba(0, 0, 0, 0.09);
}

.register-decoration {
  flex: 0.8;
  background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
  padding: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  position: relative;
  overflow: hidden;
}

.register-decoration::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(45deg, rgba(255, 255, 255, 0.1) 0%, rgba(255, 255, 255, 0) 100%);
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

.register-form-container {
  flex: 1.2;
  padding: 60px;
  background: #ffffff;
  overflow-y: auto;
  max-height: 90vh;
}

.register-card {
  box-shadow: none !important;
  border: none !important;
  width: 100%;
  max-width: 500px;
  margin: 0 auto;
  animation: fadeIn 0.8s ease-out;
}

.register-header {
  text-align: center;
  margin-bottom: 40px;
}

.register-header h2 {
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

.register-form :deep(.el-input__wrapper) {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.03);
  border-radius: 12px;
  padding: 0 15px;
  transition: all 0.3s ease;
}

.register-form :deep(.el-input__wrapper:hover) {
  transform: translateY(-1px);
  box-shadow: 0 4px 16px 0 rgba(0, 0, 0, 0.06);
}

.register-form :deep(.el-input__inner) {
  height: 48px;
  font-size: 16px;
}

.register-form :deep(.el-form-item) {
  margin-bottom: 25px;
  animation: slideInRight 0.6s ease-out calc(var(--el-transition-duration) * 0.1s) backwards;
}

.register-button {
  width: 100%;
  height: 48px;
  border-radius: 12px;
  font-size: 16px;
  margin-top: 30px;
  font-weight: 500;
  background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
  border: none;
  transition: all 0.3s ease;
  animation: slideInUp 0.6s ease-out 0.4s backwards;
}

.register-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(106, 17, 203, 0.3);
}

.register-button:active {
  transform: translateY(0);
}

.login-link {
  text-align: center;
  margin-top: 25px;
  color: #7f8c8d;
  font-size: 15px;
  animation: fadeIn 0.6s ease-out 0.5s backwards;
}

.login-button {
  color: #6a11cb;
  text-decoration: none;
  font-weight: 500;
  margin-left: 5px;
  transition: all 0.3s ease;
}

.login-button:hover {
  color: #2575fc;
  text-decoration: none;
}

/* 滚动条样式 */
.register-form-container::-webkit-scrollbar {
  width: 6px;
}

.register-form-container::-webkit-scrollbar-thumb {
  background-color: #cbd5e0;
  border-radius: 3px;
}

.register-form-container::-webkit-scrollbar-track {
  background-color: #f7fafc;
}

/* 响应式设计 */
@media (max-width: 992px) {
  .register-wrapper {
	flex-direction: column;
  }

  .register-decoration {
	padding: 40px 20px;
	min-height: 200px;
  }

  .register-form-container {
	padding: 30px 20px;
  }
}

/* 动画关键帧 */
@keyframes gradientBG {
  0% {
	background-position: 0% 50%
  }
  50% {
	background-position: 100% 50%
  }
  100% {
	background-position: 0% 50%
  }
}

@keyframes shimmer {
  0% {
	transform: translateX(-100%)
  }
  100% {
	transform: translateX(100%)
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

@keyframes slideInRight {
  from {
	opacity: 0;
	transform: translateX(30px);
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