import * as React from "react"
import {
  ChakraProvider,
  Box,
  theme,
  Flex,
} from "@chakra-ui/react"
import CategoryComponent from "./components/CategoryComponent"
import DarkModeComponent from "./components/DarkModeComponent"

const categorylist = [{id: "1", name: "JJ"}, ]

export const App = () => (
  <ChakraProvider theme={theme}>
    <Flex>
      <Box>Folder</Box>
      <Box>
        {categorylist.map((category) => <CategoryComponent id={category.id} name={category.name} />)}
      </Box>
      <DarkModeComponent />
    </Flex>
  </ChakraProvider>
)
