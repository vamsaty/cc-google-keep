import {useContext, useState} from "react";
import {useLocalStorage} from "./useLocalStorage";
import {AuthContext} from "../context/AuthV2";


export interface UserCreds {
    username: string;
    password: string;
}

export interface User {
    id: string;
    username: string;
    authToken: string;
}

export const useUser = () => {
    const [user, setUser] = useState<User|null>(null);
    const { setItem } = useLocalStorage();

    const addUser = (userInput: User) => {
        console.log('user-hook-addUser-start')
        setUser(userInput);
        setItem("user", JSON.stringify(userInput));
        console.log('user-hook-addUser-end: ', user)
    }

    const removeUser = () => {
        console.log('user-hook-removeUser-start')
        setUser(null);
        setItem("user", null)
        console.log('user-hook-removeUser-end')
    }

    return { user, addUser, removeUser };
}