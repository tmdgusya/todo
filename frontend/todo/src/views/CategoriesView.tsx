import React from "react";
import { Box, Heading, useColorMode, useColorModeValue } from "@chakra-ui/react";
import CategoryComponent from "../components/CategoryComponent";

export default function CategoriesView() {
    const categorylist = [{id: "1", name: "TAG"}, {id: "1", name: "WORK"},{id: "1", name: "DOING"},{id: "1", name: "HOBBY"}]
    
    return <Box 
        marginRight={3}
        padding={3}
        color={useColorModeValue("black", "white")} 
        width="300px" 
        height="80vh"
        border="1px"
        rounded="md"
    >
        <Heading size='lg' fontFamily='monolisa'>Categories</Heading>
        <Box marginTop='5'>
          {categorylist.map((category) => <CategoryComponent id={category.id} name={category.name} />)}
        </Box>
      </Box>
}