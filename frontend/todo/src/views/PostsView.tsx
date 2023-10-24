import { Box, useColorModeValue } from "@chakra-ui/react";
import PostComponent, { PostComponentProps } from "../components/PostComponent";

interface PostsViewProps {
    posts: PostComponentProps[];
}

export default function PostsView({posts}: PostsViewProps) {
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