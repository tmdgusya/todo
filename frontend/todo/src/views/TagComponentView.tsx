import React from "react";
import TagComponent, { TagComponentProps } from "../components/TagComponent";
import { Flex } from "@chakra-ui/react";

interface TagComponentViewProps {
    tags: TagComponentProps[];
}

export default function TagComponentView({tags}: TagComponentViewProps) {
    return <Flex>
        {tags.map((tag) => <TagComponent key={tag.id} id={tag.id} name={tag.name} color={tag.color} />)}
    </Flex>
}