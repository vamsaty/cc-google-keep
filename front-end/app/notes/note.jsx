import {useState} from "react";
import {Card, CardContent, Checkbox} from "@mui/joy";
import Typography from "@mui/joy/Typography";
import CardActions from "@mui/joy/CardActions";
import Button from "@mui/joy/Button";
import {Delete, Save} from "@mui/icons-material";
import * as React from "react";
import {USER_PATH} from "../apis/base";
import axios from "axios";


export function ContentLine(props) {
    const [striked, setStriked] = useState(false);
    return (
        <Checkbox onChange={()=>setStriked(!striked)}
                  label={props.label}
                  sx={{textDecoration: striked?'line-through': ''}}
        />
    );
}

export function SingleNote(props) {

    const deleteNote = (event) => {
        event.preventDefault();
        console.log('here iam ', props.noteId)
        axios.delete(`${USER_PATH}/note`, {

        })
    }

    const updateNote = () => {
        console.log('here we are udpate note')
    }

    return (
        <Card sx={{ height: '100%' }}
              variant="outlined"
        >
            <CardContent>
                <Typography level="title-md">{props.title}</Typography>
            </CardContent>
            <CardContent>
                {
                    props.content.map((item, index) => {
                        return (
                            <ContentLine key={index} label={item.text} text={item.text} />
                        )
                    })
                }
            </CardContent>
            <CardActions>
                <Button variant="outlined"
                        size="sm"
                        color={"danger"}
                        onClick={(e)=>deleteNote(e)}
                >
                    <Delete />
                </Button>
                <Button variant="outlined" size="sm">
                    <Save />
                </Button>
            </CardActions>
        </Card>
    );
}