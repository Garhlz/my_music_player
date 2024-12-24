import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
  // 用户信息
  const userInfo = ref({
    id: null,
    username: '',
    name: '',
    avatar: '',
    email: '',
    phone: '',
    sex: ''
  })

  // 设置用户信息
  const setUserInfo = (info) => {
    userInfo.value = {
      ...userInfo.value,
      ...info
    }
  }

  // 清除用户信息
  const clearUserInfo = () => {
    userInfo.value = {
      id: null,
      username: '',
      name: '',
      avatar: '',
      email: '',
      phone: '',
      sex: ''
    }
  }

  return {
    userInfo,
    setUserInfo,
    clearUserInfo
  }
}) 