import { Box, Text, useColorMode, useColorModeValue } from "@chakra-ui/react";
import React from "react";

export interface CategoryComponentProps {
    id: string;
    name: string;
}


export default function CategoryComponent({ id, name }: CategoryComponentProps) {
    return <Box 
                key={id} id={id} width="250px" 
                color={useColorModeValue("gray.100", "white")}
                backgroundColor={useColorModeValue("yellow.100", "#2D3748")} 
                borderColor={useColorModeValue("gray.100","#2D3748")}
                height="30px" 
                borderWidth="1px" 
                borderRadius="md"
                marginTop="2"
                paddingLeft="1"
            >
                <Text fontSize='lg' fontFamily='monolisa' color={useColorModeValue('black', 'white')}>{name}</Text>
            </Box>
}