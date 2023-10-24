import { useColorMode } from "@chakra-ui/react";
import React from "react";

interface GetColorModeHooksProps {
    light: string;
    dark: string;
}

export default function useGetColorModeHooks({light, dark}: GetColorModeHooksProps) {
    const { colorMode } = useColorMode();
    
    const bgColor = { light: light, dark: dark };

    function getColor(): string {
        console.log(colorMode)
        if(colorMode == "light") return bgColor.light
        return bgColor.dark 
    }

    return getColor()
}