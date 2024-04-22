import { defineStore } from 'pinia'


export const useInfoStore= defineStore('info',{
  state: () => ({
    nickname: localStorage.getItem('nickname') || '',
    role: localStorage.getItem('role') || '',
  }),
  getters: {
    getNickname() {
      return this.nickname
    },
    getRole() {
      return this.role
    }
  },
  actions: {
    setNickname(nickname) {
      this.nickname = nickname
      localStorage.setItem('nickname', nickname)
    },
    setRole(role) {
      this.role = role
      localStorage.setItem('role', role)
    },
    clearAll(){
      this.nickname = ''
      this.role = ''
      localStorage.removeItem('nickname')
      localStorage.removeItem('role')
    }
  }
})