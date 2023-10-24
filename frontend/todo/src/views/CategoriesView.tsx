import React from "react";
import { Box, useColorMode, useColorModeValue } from "@chakra-ui/react";
import CategoryComponent from "../components/CategoryComponent";

export default function CategoriesView() {
    const categorylist = [{id: "1", name: "JJ"}, {id: "1", name: "JJ"},{id: "1", name: "JJ"},{id: "1", name: "JJ"}]
    
    return <Box 
        marginRight={3}
        padding={3}
        color={useColorModeValue("black", "white")} 
        width="300px" 
        height="80vh"
        border="1px"
        rounded="md"
    >
        <h1>Categories</h1>
        <Box>
          {categorylist.map((category) => <CategoryComponent id={category.id} name={category.name} />)}
        </Box>
      </Box>
}