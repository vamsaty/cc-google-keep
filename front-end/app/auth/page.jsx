import {Login} from "./login";


export default function App() {
    return (
        <>
            <Login isLogin={true} />
            <Login isLogin={false} />
        </>
    )
}