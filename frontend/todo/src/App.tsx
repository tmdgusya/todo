import {
  ChakraProvider,
  theme,
  Flex,
} from "@chakra-ui/react"
import React from "react"
import { ColorModeSwitcher } from "./components/DarkModeComponent"
import CategoriesView from "./views/CategoriesView"
import PostsView from "./views/PostsView"

export default function App() {
  return <>
  <ChakraProvider theme={theme}>
    <Flex margin="10">
      <CategoriesView />
      <PostsView />
      <ColorModeSwitcher />
    </Flex>
  </ChakraProvider>
  </>
}
