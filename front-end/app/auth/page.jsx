'use client';

import {LoginV2} from "./login";
import { useEffect, useState } from "react";
import { Button, extendTheme, Grid, Sheet, useColorScheme } from "@mui/joy";
import CustomContainer from "../components/container";

export default function App() {
    console.log('app-page')

    return (
        <CustomContainer>
            <LoginV2 />
        </CustomContainer>
    )
}