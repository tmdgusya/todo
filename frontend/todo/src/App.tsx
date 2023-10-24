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
      <CategoriesView categories={[{id: "1", name: "TODO"}, {id: "2", name: "WORK"}]} />
      <PostsView posts={[
        {
          id: "1",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: [{"id": "1", "name": "tag1", "color": 'red'}, {"id": "2", "name": "tag2", "color": 'green'}],
        },
        {
          id: "2",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: [],
        },
        {
          id: "3",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: [],
        },
        {
          id: "4",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: [{id: "1", name: "tag1", "color": 'red'}, {id: "2", name: "tag2", "color": 'red'}],
        },
        {
          id: "5",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: [],
        },
]} />
      <ColorModeSwitcher />
    </Flex>
  </ChakraProvider>
  </>
}
