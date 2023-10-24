import { Box, useColorMode } from "@chakra-ui/react";
import React from "react";
import useGetColorModeHooks from "../hooks/useColorModeHooks";

interface CategoryComponentProps {
    id: string;
    name: string;
}


export default function CategoryComponent({ id, name }: CategoryComponentProps) {
    return <Box 
                id={id} width="200px" 
                color={useGetColorModeHooks({ light: "black", dark: "white" })}
                backgroundColor={useGetColorModeHooks({ light: "gray.100", dark: "#2D3748" })} 
                borderColor={useGetColorModeHooks({ light: "gray.100", dark: "#2D3748" })}
                height="30px" 
                borderWidth="1px" 
                borderRadius="md"
                marginTop="2"
            >{name}</Box>
}