import * as React from 'react';
import Box from '@mui/joy/Box';
import Drawer from '@mui/joy/Drawer';
import Button from '@mui/joy/Button';
import List from '@mui/joy/List';
import Divider from '@mui/joy/Divider';
import ListItem from '@mui/joy/ListItem';
import ListItemButton from '@mui/joy/ListItemButton';
import Link from "next/link";
import {ListItemIcon} from "@mui/material";
import {IconButton} from "@mui/joy";
import {Menu} from "@mui/icons-material";


export function LinkDecor(props) {
    const {index, currentPath, path} = props
    return (
        <Link className={currentPath===path? 'always': 'hover'} href={path}>
            {props.children}
        </Link>
    )
}


export default function DrawerBasic(props) {
    const [open, setOpen] = React.useState(false);

    const toggleDrawer = (inOpen: boolean) => (event: React.KeyboardEvent | React.MouseEvent) => {
        if (event.type === 'keydown' &&
            ((event as React.KeyboardEvent).key === 'Tab' ||
                (event as React.KeyboardEvent).key === 'Shift')
        ) {
            return;
        }
        setOpen(inOpen);
    };

    return (
        <Box sx={{ display: 'flex' }}>
            <IconButton variant="outlined" color="neutral" onClick={toggleDrawer(true)}>
                <Menu>asdf</Menu>
            </IconButton>
            <Drawer sx={{width: '200px'}} size={"sm"} open={open} onClose={toggleDrawer(false)}>
                <Box
                    role="presentation"
                    onClick={toggleDrawer(false)}
                    onKeyDown={toggleDrawer(false)}
                >
                    <List>
                        {props.routes.map((r, index)=>(
                            <ListItem key={index}>
                                <Link href={r.path}> {r.icon} </Link>
                            </ListItem>
                        ))}
                    </List>
                </Box>
            </Drawer>
        </Box>
    );
}