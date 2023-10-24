import { Box, Checkbox, Flex, Heading, Text, useColorModeValue } from "@chakra-ui/react";
import React from "react";

export interface PostComponentProps {
    id: string;
    title: string;
    content: string;
    isChecked: boolean;
    createdAt: string;
    tags: string[];
}


export default function PostComponent({id, title, content, isChecked, createdAt}: PostComponentProps) {

    // make card ui using chakra ui
    return <Flex id={id} 
                width="600px" 
                height="150px" 
                borderWidth="1px" 
                borderRadius="md"
                borderColor={useColorModeValue('gray.300', 'gray.500')}
                marginTop="2" 
                padding="2"
            justifyContent="space-between"
        >
            <Box>
            <Heading fontFamily='monolisa'>{title}</Heading>
            <Text fontSize='lg' fontFamily='monolisa'>{content}</Text>
            <Text fontFamily='monolisa'>{createdAt}</Text>
            </Box>
            <Box>
                <Checkbox colorScheme="green" defaultChecked={isChecked} />
            </Box>
        </Flex>
}