import { Box, Checkbox, Flex, Heading, Text, useColorModeValue } from "@chakra-ui/react";
import React from "react";
import TagComponentView from "../views/TagComponentView";
import { TagComponentProps } from "./TagComponent";

export interface PostComponentProps {
    id: string;
    title: string;
    content: string;
    isChecked: boolean;
    createdAt: string;
    tags: TagComponentProps[];
}


export default function PostComponent({id, title, content, isChecked, createdAt, tags}: PostComponentProps) {

    // make card ui using chakra ui
    return <Box id={id} 
                width="600px" 
                height="150px" 
                borderWidth="1px" 
                borderRadius="md"
                borderColor={useColorModeValue('gray.300', 'gray.500')}
                marginTop="2" 
                padding="2"
        >   
            <Flex justifyContent="space-between">
                <Box>
                    <Heading fontFamily='monolisa'>{title}</Heading>
                    <Text fontSize='lg' fontFamily='monolisa'>{content}</Text>
                    <Text fontFamily='monolisa'>{createdAt}</Text>
                </Box>
                <Box>
                    <Checkbox colorScheme="green" defaultChecked={isChecked} />
                </Box>
            </Flex>
            <TagComponentView tags={tags} />
        </Box>
}