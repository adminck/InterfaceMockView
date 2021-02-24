import axios from 'axios'; // 引入axios
import { Message, Loading } from 'element-ui';
import { store } from '@/store/index'


const service = axios.create({
    baseURL: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1' : '',
    timeout: 99999,
})

let acitveAxios = 0
let loadingInstance
let timer
const showLoading = () => {
    acitveAxios++
    if (timer) {
        clearTimeout(timer)
    }
    timer = setTimeout(() => {
        if (acitveAxios > 0) {
            loadingInstance = Loading.service({ fullscreen: true })
        }
    }, 400);
}

const closeLoading = () => {
    acitveAxios--
    if (acitveAxios <= 0) {
        clearTimeout(timer)
        loadingInstance && loadingInstance.close()
    }
}

//http request 拦截器
service.interceptors.request.use(
    config => {
        showLoading()
        const token = store.getters['user/token']
        if (Object.prototype.toString.call(config.data) == '[object FormData]') {
            config.headers = {
                'Content-Type': 'multipart/form-data',
                'x-token': token
            }
        }else{
            config.data = JSON.stringify(config.data);
            config.headers = {
                'Content-Type': 'application/json',
                'x-token': token
            }
        }
        return config;
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error);
    }
);


//http response 拦截器
service.interceptors.response.use(
    response => {
        closeLoading()
        if (response.data.code == 0 || response.headers.success === "true") {
            return response.data
        } else {
            Message({
                showClose: true,
                message: response.data.msg || decodeURI(response.headers.msg),
                type: 'error',
            })
            if (response.data.data && response.data.data.reload) {
                store.commit('user/LoginOut')
            }
            return Promise.reject(response.data.msg)
        }
    },
    error => {
        closeLoading()
        Message({
            showClose: true,
            message: error,
            type: 'error'
        })
        return Promise.reject(error)
    }
)

export default service