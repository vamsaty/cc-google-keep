'use client';

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import {
    Button, CircularProgress,FormControl, FormLabel,
    Input, Skeleton, Snackbar, ToggleButtonGroup,
    Typography
} from "@mui/joy";
import {doLogin, doRegister} from "../apis/auth";
import { useAuthV2 } from "../context/AuthV2";
import { UserCreds } from "../hooks/useUser";


const initialAuthState = {
    username: "",
    password: ""
}

const initialMessage = {
    show: false,
    body: '',
    color: 'success'
}

const LoginV2 = () => {
    const { isAuthenticated, isLoading, login, register, result } = useAuthV2()
    const [cred, setCred] = useState(initialAuthState);
    const [value, setValue] = useState('login');
    const [message, setMessage] = useState(initialMessage);
    const router = useRouter();


    useEffect(()=>{
        if(!isLoading && isAuthenticated) {
            router.replace('/')
        }
    }, [isLoading]);
    const updateUsername = (e) => { setCred({...cred, username: e.target.value}) }
    const updatePassword = (e) => { setCred({...cred, password: e.target.value}) }

    const handleLogin = () => {
        const res = login(cred)
        res.then( data => {
            if (data.status === 200) {
                router.replace('/')
            } else {
                setMessage({
                    ...message,
                    show: true,
                    color: 'danger',
                    body: 'failed to login ' + data.response.data.message
                })
            }
        })
    }

    const handleLogout = () => { console.log("handle-logout") }

    const handleRegister = () => {
        const res = register(cred)
        res.then(data => {
            if(data.status === 200) {
                setMessage({
                    ...message,
                    show: true,
                    color: 'success',
                    body: 'Registered successfully! Kindly Login'
                })
                setValue('login')
                setCred(initialAuthState)
            } else {
                setMessage({
                    ...message,
                    show: true,
                    color: 'danger',
                    body: 'failed to create account ' + data.response.data.message
                })
            }
        });
    }

    return (
        <>
            <Typography>
                <Skeleton loading={isLoading}>
                    {isLoading? 'loading': <b>Welcome!</b>}
                </Skeleton>
            </Typography>
            <Snackbar open={message.show}
                      autoHideDuration={1000}
                      color={message.color}
                      onClose={()=>{setMessage(initialMessage)}}>
                {message.body}
            </Snackbar>
            {
                !isLoading &&
                <ToggleButtonGroup color="neutral">
                    <Button aria-pressed={value === 'login'}
                            variant="outlined"
                            color="neutral"
                            onClick={(e)=>setValue('login')}
                            sx={(theme) => ({
                                [`&[aria-pressed="true"]`]: {
                                    ...theme.variants.outlinedActive.neutral,
                                    borderColor: theme.vars.palette.neutral.outlinedHoverBorder,
                                },
                            })}
                    >Login</Button>
                    <Button aria-pressed={value === 'register'}
                            variant="outlined"
                            onClick={(e)=>setValue('register')}
                            color="neutral"
                            sx={(theme) => ({
                                [`&[aria-pressed="true"]`]: {
                                    ...theme.variants.outlinedActive.neutral,
                                    borderColor: theme.vars.palette.neutral.outlinedHoverBorder,
                                },
                            })}
                    >Register</Button>
                </ToggleButtonGroup>
            }
            <Typography level="body-sm" >
                {isLoading?
                    <CircularProgress/>
                    :
                    (value=='login')?
                        "Sign in to continue!"
                        :
                        "Create a new account"
                }
            </Typography>

            {!isLoading && <FormControl orientation={"vertical"}>
                <FormLabel>username</FormLabel>
                <Input placeholder={"username"}
                       type={"text"}
                       value={cred.username}
                       onChange={(e) => updateUsername(e)}
                />
            < /FormControl>}

            {!isLoading && <FormControl orientation={"vertical"}>
                <FormLabel>password</FormLabel>
                <Input placeholder={"password"}
                       type={"password"}
                       value={cred.password}
                       onChange={(e) => updatePassword(e)}
                />
            </FormControl>}
            {
                !isLoading && (value === 'login' ?
                    <Button disabled={isLoading} variant={'soft'} sx={{ mt: 1 }} onClick={handleLogin}>Login</Button>
                    :
                    <Button disabled={isLoading} variant={'soft'} sx={{ mt: 1 }} onClick={handleRegister}>Register</Button>)
            }
        </>
    )
}


export {
    LoginV2
}