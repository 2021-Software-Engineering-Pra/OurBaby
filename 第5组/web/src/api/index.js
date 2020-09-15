import http from './public.js'
const baseUrl = '/api/'
// 登陆
export const userLogin = (params) => {
  return http.fetchGet(`${baseUrl}login`, params)
}
// 退出登陆
export const loginOut = (params) => {
  return http.fetchPost(`${baseUrl}logout`, params)
}
export const NothavePaypass = (params) => {
  return http.fetchPost(`${baseUrl}nothavepaypass`, params)
}
// 用户信息
export const userInfo = (params) => {
  return http.fetchGet(`${baseUrl}userinfo`, params)
}
export const isUname = (params) => {
  return http.fetchGet(`${baseUrl}isuname`, params)
}
export const updateEmail = (params) => {
  return http.fetchGet(`${baseUrl}updateemail`, params)
}
export const updatePhone = (params) => {
  return http.fetchGet(`${baseUrl}updatephone`, params)
}
export const updatePaypass = (params) => {
  return http.fetchGet(`${baseUrl}updatepaypasswd`, params)
}
// 注册账号
export const register = (params) => {
  return http.fetchPost(`${baseUrl}register`, params)
}
export const updatePasswd = (params) => {
  return http.fetchGet(`${baseUrl}updatepasswd`, params)
}
export const updateAddress = (params) => {
  return http.fetchPost(`${baseUrl}updateaddress`, params)
}
// 收货地址
export const getAddress = (params) => {
  return http.fetchGet(`${baseUrl}getaddress`, params)
}
export const addAddress = (params) => {
  return http.fetchPost(`${baseUrl}addaddress`, params)
}
export const delAddress = (params) => {
  return http.fetchGet(`${baseUrl}deladdress`, params)
}
  //删除用户
  export const deluser =(params)=>{
    return http.fetchPost(`${baseUrl}deluser`,params)
  }
