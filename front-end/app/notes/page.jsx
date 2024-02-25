'use client';

import {useEffect, useState} from "react";
import axios from "axios";
import {USER_PATH} from "../apis/base";
import {CircularProgress, Typography} from "@mui/joy";
import * as React from "react";
import {NoteSheet} from "./notes";
import {useAuthV2} from "../context/AuthV2";

const initNotesData = [];

export default function App() {
    const auth = useAuthV2();
    const [notesData, setNotesData] = useState(initNotesData)

    console.log("inside-notes-page[AUTH]", auth)
    useEffect(()=>{
        console.log('fetching data')
        setNotesData(["asf"])
        axios.get(`${USER_PATH}/notes`)
        //     .then(res => {
        //         if (res.status === 200) {
        //             setNotesData(res.data)
        //         }
        //     }).catch(err => {
        // })
        console.log("NOTES", notesData)
    }, [])

    return (
        <>
            <Typography>Notes Page</Typography>
            <h2>asfsdf</h2>
            {/*{auth.isLoading? <CircularProgress /> : <NoteSheet data={notesData} />}*/}
        </>
    )
}