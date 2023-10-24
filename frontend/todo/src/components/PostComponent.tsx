import { Box } from "@chakra-ui/react";
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
    return <Box id={id} width="200px" height="200px" borderWidth="1px" borderRadius="md" marginTop="2">
            <h1>{title}</h1>
            <p>{content}</p>
            <p>{isChecked}</p>
            <p>{createdAt}</p>
        </Box>
}