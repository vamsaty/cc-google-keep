'use client';

import axios from 'axios';
import React, {createContext, useEffect, useState} from "react";

const BASE_PATH = 'http://localhost:8099';
const AUTH_PATH = `${BASE_PATH}/auth`;


export const getAuthToken = () => {
    return localStorage.getItem('token');
}

// interceptors
axios.interceptors.response.use(
    response => response,
    error => {
        if (error.response.status === 401) {
            window.location = `${AUTH_PATH}/login`;
        }
        return Promise.reject(error);
    }
);

axios.interceptors.request.use(config => {
    if (config.url.includes('login') || config.url.includes('register')) {
        return config;
    }
    const token = localStorage.getItem('token');
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
    });
}

const doLogout = () => {
    return axios.post(`${AUTH_PATH}/logout`);
}

const AuthContext = createContext()

const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            axios.get(`${AUTH_PATH}/user`)
                .then(response => {
                    setUser(response.data);
                })
                .catch(error => {
                    console.log(error);
                });
        }
    }, []);

    const login = (username, password) => {
        doLogin(username, password)
            .then(response => {
                localStorage.setItem('token', response.data.token);
                setUser(response.data.user);
            })
            .catch(error => {
                console.log(error);
            });
    }

    const register = (username, password) => {
        doRegister(username, password)
            .then(response => {
                localStorage.setItem('token', response.data.token);
                setUser(response.data.user);
            })
            .catch(error => {
                console.log(error);
            });
    }

    const logout = () => {
        doLogout()
            .then(response => {
                localStorage.removeItem('token');
                setUser(null);
            })
            .catch(error => {
                console.log(error);
            });
    }
    console.log("AuthProvider loaded")
    return (
        <AuthContext.Provider value={{ isAuthenticated: !!user, user, login, register, logout }}>
            {children}
        </AuthContext.Provider>
    );
}

console.log("auth.jsx loaded");

const useAuth = () => {return React.useContext(AuthContext);}

export {
    doLogout,
    doLogin,
    doRegister,
    useAuth,
    AuthProvider
}