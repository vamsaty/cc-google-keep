// interceptors
import axios from "axios";
import {AUTH_PATH, AUTH_TOKEN} from "./base";
import {getAuthToken} from "../context/AuthV2";
import {useLocalStorage} from "../hooks/useLocalStorage";

axios.interceptors.response.use(
    response => {
        console.log("RESPONSE_INTERCEPTOR_<RESPONSE>");
        return {
            raw: response,
            status: response.status,
            message: "",
            data: response.data
        }
    },
    error => {
        console.log("RESPONSE_INTERCEPTOR_<ERROR>", error)
        if (error.response && error.response.status === 401) {
            window.location = `/auth`;
        }
        return Promise.reject(error);
    }
);

axios.interceptors.request.use(config => {
    // if authentication call
    if (config.url.includes('login') || config.url.includes('register')) {
        console.log("here we are interceptor")
        return config;
    }
    // include token in other calls
    const token = getAuthToken();
    console.log("TOKEN", getAuthToken())
    if (token) {
        config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
})

// what are interceptors?
// interceptors are functions that are called before the request is sent and after the response is received
// they are used to modify the request or response
// in this case, we are using them to redirect the user to the login page if the request returns a 401 error



const doLogin = (username, password) => {
    return axios.post(`${AUTH_PATH}/login`, {
        username: username,
        password: password
    });
}

const doRegister = (username, password) => {
    return axios.post(`${AUTH_PATH}/register`, {
        username: username,
        password: password
    })
}

const doLogout = () => {
    return axios.post(`${AUTH_PATH}/logout`);
}


export {
    doLogout,
    doLogin,
    doRegister
}