import axios, { type AxiosRequestConfig, type AxiosResponse } from "axios";

// 路径
axios.defaults.baseURL = `http://localhost:7003/`;
// 请求拦截器
axios.interceptors.request.use((config: AxiosRequestConfig | any) => {
    return config
})

// 响应拦截器
axios.interceptors.response.use((res: AxiosResponse) => {
    return res
}, (err: any) => {
    return Promise.reject(err)
});

export default axios;