import {User, UserCreds, useUser} from "../hooks/useUser";
import {createContext, useContext, useEffect, useState} from "react";
import {useMemo} from "react";
import {useLocalStorage} from "../hooks/useLocalStorage";
import {doLogin, doRegister} from "../apis/auth";
import {USER_PATH} from "../apis/base";


export interface AuthResponse {
    ok: boolean;
    err: any,
    message: string;
    isLoading: boolean;
}

const defaultAuthResponse = {
    ok: true,
    message: "",
    err: null,
    isLoading: true,
}

export interface AuthContext {
    isAuthenticated: boolean;
    isLoading: boolean;
    login: (creds: UserCreds) => any;
    logout: () => void;
    register: (creds: UserCreds) => any;
    user: User | null;
}


export const AuthContext = createContext<AuthContext>({
    isAuthenticated: false,
    isLoading: true,
    login: (creds) => {},
    logout: () => {},
    register: (creds) => {},
    user: null,
})

export const getAuthToken = () => {
    return JSON.parse(localStorage.getItem('user')).authToken
}

export const AuthProviderV2 = ({ children }) => {
    const { user, addUser, removeUser } = useUser();
    const [isLoading, setIsLoading] = useState(true);
    const { getItem, setItem } = useLocalStorage();

    useEffect(() => {
        console.log("here we are")
        const user = getItem("user")
        if(user) {
            addUser(JSON.parse(user))
        }
        setIsLoading(false);
    }, []);

    const login = (creds: UserCreds) => {
        return doLogin(creds.username, creds.password).then(res=>{
            addUser({
                username: creds.username,
                id  : creds.username,
                authToken: res.data,
            })
            return res
        }).catch(err => {
            return err
        })
    }

    const register = (creds: UserCreds) => {
        return doRegister(creds.username, creds.password).then(res=>{
            return res
        }).catch(err => {
            return err
        })
    }

    const logout = () => { removeUser(); }

    return <AuthContext.Provider value={{
        isAuthenticated: !!user,
        user, login, logout, register,
        isLoading
    }}>
        {children}
    </AuthContext.Provider>

}

export const useAuthV2 = () => {
    return useContext(AuthContext)
}