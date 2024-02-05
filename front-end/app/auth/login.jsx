'use client';

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import TextField from '@mui/material/TextField';
import { Button, Grid } from "@mui/joy";
import {doLogin, doRegister, getAuthToken} from "./auth";

const initialAuthState = {
    username: "",
    password: ""
}

function Login(props) {
    const router = useRouter();
    const [authState, setAuthState] = useState(initialAuthState)
    const [loading, setLoading] = useState(true)

    useEffect(()=>{
        if(getAuthToken()) { router.replace('/') }
        setLoading(false)
    }, [router])

    const updateUsername = (e) => {
        setAuthState({...authState, username: e.target.value})
    }

    const updatePassword = (e) => {
        setAuthState({...authState, password: e.target.value})
    }

    if(loading) {
        return <h4>loading...</h4>
    }


    const handleLogin = () => {
        try {
            doLogin(authState.username, authState.password).then(res => {
                if(res.ok) {
                    handleToken(res.headers)
                    router.replace('/')
                }
            }).catch(err=> {
                console.log(err)
                return {
                    error: err,
                    message: "FAILED"
                }
            })
        } catch (err) {
            console.log("Error: ", err)
            return {
                error: err,
                message: "LOGIN_EXCEPTION"
            }
        }
    }

    const handleRegister = () => {
        try {
            doRegister(authState.username, authState.password).then(res => {
                if(res.ok) {
                    router.replace('/')
                }
            }).catch(err => {
                console.log(err)
                return {
                    error: err,
                    message: "FAILED"
                }
            })
        } catch (err) {
            console.log("Error: ", err)
            return {
                error: err,
                message: "EXCEPTION"
            }
        }
    }

    return (
        <Grid container direction={"column"} alignItems={"center"} gap={"10px"}>
            <TextField
                id="outlined-text-input" defaultValue={""}
                label="username"  onChange={(e) => updateUsername(e)}
            />
            <TextField
                id="outlined-password-input" label="Password" type="password"
                autoComplete="current-password" onChange={(e) => {updatePassword(e)}}
            />
            {
                props.isLogin ?
                    <Button variant="outlined" onClick={handleLogin}> Login </Button> :
                    <Button variant="outlined" onClick={handleRegister}> Register </Button>
            }
        </Grid>
    )
}


export {
    Login
}