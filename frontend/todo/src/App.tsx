import * as React from "react"
import {
  ChakraProvider,
  Box,
  theme,
  Flex,
} from "@chakra-ui/react"
import CategoryComponent from "./components/CategoryComponent"
import PostComponent, { PostComponentProps } from "./components/PostComponent"
import { ColorModeSwitcher } from "./components/DarkModeComponent"

const categorylist = [{id: "1", name: "JJ"}, {id: "1", name: "JJ"},{id: "1", name: "JJ"},{id: "1", name: "JJ"},]
const posts: PostComponentProps[] = [
  {
    id: "1",
    title: "title",
    content: "content",
    isChecked: false,
    createdAt: "2021-08-15",
    tags: ["tag1", "tag2"],
  },
  {
    id: "1",
    title: "title",
    content: "content",
    isChecked: false,
    createdAt: "2021-08-15",
    tags: ["tag1", "tag2"],
  },
  {
    id: "1",
    title: "title",
    content: "content",
    isChecked: false,
    createdAt: "2021-08-15",
    tags: ["tag1", "tag2"],
  },
  {
    id: "1",
    title: "title",
    content: "content",
    isChecked: false,
    createdAt: "2021-08-15",
    tags: ["tag1", "tag2"],
  },
  {
    id: "1",
    title: "title",
    content: "content",
    isChecked: false,
    createdAt: "2021-08-15",
    tags: ["tag1", "tag2"],
  },
]

export const App = () => (
  <ChakraProvider theme={theme}>
    <Flex margin="10">
      <Box>
        <h1>Categories</h1>
        <Box>
          {categorylist.map((category) => <CategoryComponent id={category.id} name={category.name} />)}
        </Box>
      </Box>
      <Box>
        {posts.map((post) => <PostComponent id={post.id} title={post.title} content={post.content} isChecked={post.isChecked} createdAt={post.createdAt} tags={post.tags} />)}
      </Box>
      <ColorModeSwitcher />
    </Flex>
  </ChakraProvider>
)
