import axios from "axios";
import {USER_PATH} from "./base";


export function getNotes(){
    const result = {
        "notes": [],
        "isError": false,
        "error": "",
    }
    axios.get(`${USER_PATH}/notes`)
        .then(res => {
            if (res.status == 200) {
                console.log("2222v22REPOSNE_JSS", res.data, res.status)
                result.notes = res.data.notes
            }
            return result
        }).catch(err => {
        result.isError = true
        result.error = "failed to fetch notes"
        return result
    })
    return result
}