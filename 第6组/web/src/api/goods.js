import http from './public.js'
const baseUrl = '/api/'
// 电脑列表
export const getNewGoods = (...params) => {
  return http.fetchGet(`${baseUrl}getnew`, params)
}
export const getHotGoods = (...params) => {
  return http.fetchGet(`${baseUrl}gethot`, params)
}
// 获取购物车列表
export const getCart= (params) => {
  return http.fetchGet(`${baseUrl}getcart`, params)
}
// 加入购物车
export const addCart = (params) => {
  return http.fetchGet(`${baseUrl}addcart`, params)
}

// 删除购物车
export const delCart = (params) => {
  return http.fetchGet(`${baseUrl}delcart`, params)
}
// 编辑购物车
export const setCartNum = (params) => {
  return http.fetchGet(`${baseUrl}setcart`, params)
}
// 支付
export const toPay = (params) => {
  return http.fetchGet(`${baseUrl}topay`, params)
}
export const cancelOrder = (params) => {
  return http.fetchGet(`${baseUrl}cancelorder`, params)
}
export const receiveProduct = (params) => {
  return http.fetchGet(`${baseUrl}receive`, params)
}
// 订单
export const getOrder = (params) => {
  return http.fetchGet(`${baseUrl}getorder`, params)
}
export const addOrder = (params) => {
  return http.fetchPost(`${baseUrl}addorder`, params)
}
// 商品详情
export const getProduct = (params) => {
  return http.fetchGet(`${baseUrl}getproduct`, params)
}
export const Products = (params) => {
  return http.fetchPost(`${baseUrl}products`, params)
}
export const getKind = (params) => {
  return http.fetchGet(`${baseUrl}getkind`, params)
}
// 删除订单
export const delOrder = (params) => {
  return http.fetchGet(`${baseUrl}delorder`, params)
}
//搜索
export const searchProduct = (params) => {
  return http.fetchGet(`${baseUrl}search`, params)
}
