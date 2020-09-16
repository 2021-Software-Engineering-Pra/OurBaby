import http from './public.js'
const baseUrl = '/api/'
//增加商品
export const addProduct = (params) => {
    return http.fetchPost(`${baseUrl}addproduct`, params)
  }
  export const getAllProduct = (params) => {
    return http.fetchGet(`${baseUrl}getallproduct`, params)
  }
  export const delProduct = (params) => {
    return http.fetchPost(`${baseUrl}delproduct`, params)
  }
  export const updateProduct = (params) => {
    return http.fetchPost(`${baseUrl}updateproduct`, params)
  }
  export const getallOrder = (params) => {
    return http.fetchGet(`${baseUrl}getallorder`, params)
  }
  export const updateOrder = (params) => {
    return http.fetchPost(`${baseUrl}updateorder`, params)
  }
  //删除用户
export const deluser =(params)=>{
    return http.fetchPost(`${baseUrl}deluser`,params)
  }
  //获取所有用户信息
  export const getalluser=()=>{
    return http.fetchGet(`${baseUrl}getalluser`)
  }
  //更新用户信息
  export const updateuser=(params)=>{
    return http.fetchPost(`${baseUrl}updateuser`,params)
  }
  //重置用户密码
  export const resetuser=(params)=>{
    return http.fetchPost(`${baseUrl}resetuser`,params)
  }