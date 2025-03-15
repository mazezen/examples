// api 接口地址
import axios from "../utils/requests";
import type { IRuleForm } from "../utils/types";

// 登录
export const login = (data: IRuleForm) => axios.post('/login', data);

// 获取导航、菜单
export const getMenu = () => axios.post('/ges/menu');