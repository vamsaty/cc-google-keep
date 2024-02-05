'use client';

import { Grid, Sheet } from "@mui/joy";
import { useAuth } from "./auth/auth";


function CreateNote() {
    return (
        <>
        <Sheet >
            <h2>create a new note</h2>
        </Sheet>
        </>
    )
}


export default function MyApp({ children }: any) {
    return (
        <Grid container className="flex align-content-center border-2 h-max border-black">
            {children}
        </Grid>
    )
}
