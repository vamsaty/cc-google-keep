import {Grid, Sheet} from "@mui/joy";
import {IconButtonRoot} from "@mui/joy/IconButton/IconButton";
import {MenuBook} from "@mui/icons-material";


export function FullSheet({ children }) {
    return (
        <main>
            <Sheet sx={{
                width: '100%', height: '100vh', py: 10,
                display: 'flex', flexDirection: 'column',
                bgcolor: 'whitesmoke', borderRadius: 'sm',
                boxShadow: 'md'
            }}>
                {children}
            </Sheet>
        </main>
    )
}

export function RightGrid(props) {
    return (
        <Grid container
              sx={{
                  width: '100%',
                  height:'100%',
              }}
              direction={"row"} alignItems={"center"}
              justifyContent={"space-between"}
              {...props}
        >
            {props.children}
        </Grid>
    )
}

export function LeftGrid(props) {
    // set direction
    let colRow = props.direction;
    if (colRow === undefined) {
        colRow = "column"
    }

    // set width
    let width = '100%';
    if (props.width !== undefined) {
        width = props.width
    }

    // set height
    let height = 'auto';
    if(props.height !== undefined) {
        height = props.height
    }
    return (
        <Grid container
              sx={{width: width, height:height}}
              direction={colRow}
              alignItems={"start"}
              flexWrap={"nowrap"}
              justifyContent={"start"}
              {...props}
        >
            {props.children}
        </Grid>
    )
}


export function LeftStrip(props) {
    let colRow = props.direction;
    if (colRow === undefined) {
        colRow = "column"
    }
    return (
        <LeftGrid width={'auto'} direction={colRow}>
            {props.children}
        </LeftGrid>
    )
}

export function CenterGrid({ children }) {
    return (
        <Grid container sx={{width: '100%', height:'100%'}}
              direction={"column"} alignItems={"center"}
              justifyContent={"center"}
        >
            {children}
        </Grid>
    )
}

export function CenterTopGrid(props) {
    return (
        <Grid container sx={{width: '100%'}}
              direction={props.direction}
              justifyContent={"start"}
        >
            {props.children}
        </Grid>
    )
}

export function GridSheet({ children }) {
    return (
        <Grid container sx={{width: '100%', height:'100%'}}
              direction={"column"} alignItems={"center"}
              justifyContent={"center"}
        >
            <Sheet sx={{
                width: 300, mx: 'auto', my: 4, py: 3, px: 2, gap: 2,
                display: 'flex', flexDirection: 'column', alignItems: 'center',
                borderRadius: 'sm', boxShadow: 'md'
            }}>
                {children}
            </Sheet>
        </Grid>
    )
}

export default function CustomContainer({ children }) {
    return (
        <FullSheet>
            <GridSheet>
                {children}
            </GridSheet>
        </FullSheet>
    )
}

export function SmallContainer({ children }) {
    return (
        {children}
    )
}


export function TopLeftMarginPaper({ children }) {
    return (
        <LeftGrid sx={{bgcolor: 'red', height: '100%'}} direction={"row"}>
            <MenuBook/>
        </LeftGrid>
    )
}

