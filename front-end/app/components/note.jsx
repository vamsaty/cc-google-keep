'use client';

import {CenterGrid, CenterTopGrid, LeftGrid, LeftStrip, RightGrid} from "./container";
import {DeleteRounded, EditNote, EditNoteRounded} from "@mui/icons-material";
import {Checkbox, IconButton, Sheet} from "@mui/joy";
import {useAuthV2} from "../context/AuthV2";
import {useEffect, useState} from "react";
import axios from "axios";
import {USER_PATH} from "../apis/base";

function NoteBody(props) {
    return (
        <>
            {
                props.content.map((item, index) => {
                    return (
                        <ContentLine key={index} label={item.text} text={item.text} />
                    )
                })
            }
        </>
    );
}

export function ContentLine(props) {
    const [striked, setStriked] = useState(false);
    return (
        <Checkbox onChange={()=>setStriked(!striked)}
                  label={props.label}
                  sx={{textDecoration: striked?'line-through': ''}}
        />
    );
}

function NoteUtility() {
    return (
        <>
            <LeftStrip>
                <IconButton color={'primary'} sx={{borderRadius: '50%'}} variant={'plain'}>
                    <EditNoteRounded></EditNoteRounded>
                </IconButton>
                <IconButton color={'danger'} sx={{borderRadius: '50%'}} variant={'plain'}>
                    <DeleteRounded></DeleteRounded>
                </IconButton>
            </LeftStrip>
        </>
    )
}

export default function SingleNote() {
    const [notesData, setNotesData] = useState([]);
    const auth = useAuthV2();

    useEffect(()=>{
        axios.get(`${USER_PATH}/notes`)
            .then(res => {
                if (res.status === 200) {
                    setNotesData(res.data)
                }
            }).catch(err => {
        })
        console.log("NOTES", notesData)
    }, [auth.isAuthenticated])

    return (
        <LeftGrid>
            <RightGrid boxShadow={"xs"} style={{padding: '5px'}}>
                <NoteBody content={{notesData}}></NoteBody>
                <NoteUtility></NoteUtility>
            </RightGrid>
        </LeftGrid>
    );
}