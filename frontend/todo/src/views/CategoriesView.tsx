import React from "react";
import { Box, Heading, useColorMode, useColorModeValue } from "@chakra-ui/react";
import CategoryComponent, { CategoryComponentProps } from "../components/CategoryComponent";

interface CategoriesViewProps {
    categories: CategoryComponentProps[];
}

export default function CategoriesView({categories}: CategoriesViewProps) {    
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
          {categories.map((category) => <CategoryComponent key={category.id} id={category.id} name={category.name} />)}
        </Box>
      </Box>
}