'use client';

import {AuthContext, AuthProviderV2, useAuth, useAuthV2} from "./context/AuthV2";
import {CircularProgress, Sheet, Typography} from "@mui/joy";
import {useEffect} from "react";
import {useRouter} from "next/navigation";
import {LoginV2} from "./auth/login";
import CustomContainer, {
    CenterGrid,
    FullSheet,
    GridSheet,
    LeftGrid, LeftStrip, RightGrid,
    TopLeftMarginPaper
} from "./components/container";
import { EditNote } from "@mui/icons-material";
import SingleNote from "./components/note";
import ListItem from "@mui/joy/ListItem";


export default function AppPage({ children }){

    const auth = useAuthV2();
    const router = useRouter();
    useEffect(()=>{
        if(!auth.isLoading && !auth.isAuthenticated) {
            router.replace('/auth')
        }
    }, [auth.isLoading])

    return (
        <main>
            <Typography>Welcome</Typography>
            {children}
        </main>
    )
}

export function NoteUtility() {
    return (
        <>
            <EditNote></EditNote>
        </>
    )
}

export const ProtectedApp = ({ children }) => {
    return (
        <AuthProviderV2>
            {children}
        </AuthProviderV2>
    )
}