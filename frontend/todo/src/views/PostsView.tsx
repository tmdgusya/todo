import { Box, useColorModeValue } from "@chakra-ui/react";
import PostComponent, { PostComponentProps } from "../components/PostComponent";

export default function PostsView() {
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
          id: "2",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: ["tag1", "tag2"],
        },
        {
          id: "3",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: ["tag1", "tag2"],
        },
        {
          id: "4",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: ["tag1", "tag2"],
        },
        {
          id: "5",
          title: "title",
          content: "content",
          isChecked: false,
          createdAt: "2021-08-15",
          tags: ["tag1", "tag2"],
        },
      ]
    return <Box 
    border='1px' 
    borderColor={useColorModeValue('black', 'white')} 
    rounded='md' 
    padding='3'
    >
        {posts.map((post) => <PostComponent 
            key={post.id} 
            id={post.id} 
            title={post.title} 
            content={post.content} 
            isChecked={post.isChecked} 
            createdAt={post.createdAt} 
            tags={post.tags} 
        />)}
  </Box>
}