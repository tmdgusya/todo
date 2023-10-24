import { Badge, Box, Text, useColorMode, useColorModeValue } from "@chakra-ui/react";
import React from "react";

export interface TagComponentProps {
    id: string;
    name: string;
    color: string;
}

export default function TagComponent({id, name, color}: TagComponentProps) {
    return <Box key={id} id={id}>
        <Badge 
            fontFamily='monolisa'
            color={'black'} 
            backgroundColor={`${color}.300`} 
            marginRight="1"
        >
            {name}
        </Badge>
    </Box>
}