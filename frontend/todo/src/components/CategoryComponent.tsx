import { Box, useColorMode, useColorModeValue } from "@chakra-ui/react";
import React from "react";

interface CategoryComponentProps {
    id: string;
    name: string;
}


export default function CategoryComponent({ id, name }: CategoryComponentProps) {
    return <Box 
                key={id} id={id} width="200px" 
                color={useColorModeValue("gray.100", "#2D3748")}
                backgroundColor={useColorModeValue("yellow.100", "#2D3748")} 
                borderColor={useColorModeValue("gray.100","#2D3748")}
                height="30px" 
                borderWidth="1px" 
                borderRadius="md"
                marginTop="2"
            >{name}</Box>
}