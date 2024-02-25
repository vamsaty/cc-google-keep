'use client';

import {SingleNote} from "./note";
import {Sheet} from "@mui/joy";

export function NoteSheet(props) {
    return (
        <Sheet sx={{width: 500, height: 200}}>
            {
                props.data.map(
                    (note, index)=> (
                        <SingleNote
                            key={note._id}
                            noteId={note._id}
                            title={note.title}
                            content={note.content}
                        />
                    )
                )
            }
        </Sheet>
    )
}